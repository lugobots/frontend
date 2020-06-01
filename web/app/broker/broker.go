package broker

import (
	"bitbucket.org/makeitplay/lugo-frontend/web/app"
	"context"
	"encoding/json"
	"fmt"
	"github.com/golang/protobuf/jsonpb"
	"github.com/lugobots/lugo4go/v2/lugo"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"io"
	"sync"
	"time"
)

var maxIgnoredMessaged = 20
var maxGRPCReconnect = 10
var grpcReconnectInterval = time.Second

const ErrMaxConnectionAttemptsReached = app.Error("did not connect to the game server")

var defaultColor = &lugo.TeamColor{}
var defaultTeamColors = &lugo.TeamSettings{
	Colors: &lugo.TeamColors{
		Primary:   defaultColor,
		Secondary: defaultColor,
	},
}

func NewBinder(config app.Config, logger *zap.SugaredLogger) *Binder {
	return &Binder{
		consumers:    map[string]chan app.FrontEndUpdate{},
		consumersMux: sync.RWMutex{},
		config:       config,
		Logger:       logger,
		gameSetup: &lugo.GameSetup{
			HomeTeam: defaultTeamColors,
			AwayTeam: defaultTeamColors,
		},
	}
}

type Binder struct {
	consumers    map[string]chan app.FrontEndUpdate
	consumersMux sync.RWMutex
	config       app.Config
	gameSetup    *lugo.GameSetup
	producerConn *grpc.ClientConn
	producer     lugo.BroadcastClient
	remoteConn   lugo.RemoteClient
	stopRequest  bool
	Logger       *zap.SugaredLogger
	lastUpdate   app.FrontEndUpdate
}

func (b *Binder) GetRemote() lugo.RemoteClient {
	return b.remoteConn
}

func (b *Binder) StreamEventsTo(uuid string) chan app.FrontEndUpdate {
	b.consumersMux.Lock()
	defer b.consumersMux.Unlock()
	b.consumers[uuid] = make(chan app.FrontEndUpdate, maxIgnoredMessaged)
	// it won't block because the channel is still empty and its cap is larger than 1
	b.consumers[uuid] <- b.lastUpdate
	return b.consumers[uuid]
}

func (b *Binder) GetGameConfig(uuid string) (app.FrontEndSet, error) {
	state := app.ConnStateUp
	if b.producerConn == nil {
		state = app.ConnStateDown
	}

	go func() {
		time.Sleep(1 * time.Second)
		b.consumersMux.RLock()
		defer b.consumersMux.RUnlock()
		b.consumers[uuid] <- b.lastUpdate
		b.Logger.Warn("sending last update")
	}()
	marshal := jsonpb.Marshaler{
		OrigName:     true,
		EmitDefaults: true,
	}
	raw, err := marshal.MarshalToString(b.gameSetup)
	if err != nil {
		return app.FrontEndSet{}, fmt.Errorf("error marshalling event message: %w", err)
	}
	return app.FrontEndSet{
		GameSetup:       json.RawMessage(raw),
		ConnectionState: state,
	}, nil
}

func (b *Binder) connect() error {
	opts := []grpc.DialOption{grpc.WithBlock()}
	ctx, _ := context.WithTimeout(context.Background(), 3*time.Second)

	if b.config.GRPCInsecure {
		opts = append(opts, grpc.WithInsecure())
	}
	var err error
	b.producerConn, err = grpc.DialContext(ctx, b.config.GRPCAddress, opts...)
	if err != nil {
		return err
	}

	b.producer = lugo.NewBroadcastClient(b.producerConn)
	b.gameSetup, err = b.producer.GetGameSetup(ctx, &lugo.WatcherRequest{
		Uuid: "frontend",
	})

	b.remoteConn = lugo.NewRemoteClient(b.producerConn)

	if err != nil {
		return err
	}
	return err
}

func (b *Binder) ListenAndBroadcast() error {
	tries := 0
	var finalErr error
	for tries < maxGRPCReconnect && !b.stopRequest {
		if err := b.connect(); err != nil {
			b.broadcastConnectionLost()
			b.Logger.Warnf("failure on connecting to the game server: %s", err)
			time.Sleep(grpcReconnectInterval)
			tries++
		} else {
			b.broadcastConnectionRees()
			tries = 0
			err := b.broadcast()
			if err == app.ErrGameOver {
				finalErr = err
				break
			}
			b.broadcastConnectionLost()
			b.Logger.Warnf("broadcast interrupted: %s", err)
		}
	}
	if b.stopRequest {
		finalErr = app.ErrStopRequested
	}
	if tries >= maxGRPCReconnect {
		finalErr = ErrMaxConnectionAttemptsReached
	}
	if err := b.Stop(); err != nil {
		return fmt.Errorf("error stopping: %w (initial error: %s)", err, finalErr)
	}

	return finalErr
}

func (b *Binder) Stop() error {
	b.stopRequest = true
	if b.producerConn != nil {
		if err := b.producerConn.Close(); err != nil {
			return err
		}
	}
	b.dropAllConsumers()
	return nil
}

func (b *Binder) broadcast() error {
	ctx := context.Background()
	receiver, err := b.producer.OnEvent(ctx, &lugo.WatcherRequest{
		Uuid: "frontend",
	})
	if err != nil {
		return err
	}
	b.Logger.Warn("starting broadcasting")
	/**
	@todo the frontend server currently has no accurate information about the debugging state, so we presume it is not paused
	*/
	debugging := false
	for {
		event, err := receiver.Recv()
		if err != nil {
			if err != io.EOF {
				return fmt.Errorf("%w: %s", app.ErrGRPCConnectionLost, err)
			}
			return app.ErrGRPCConnectionClosed
		}
		eventType, err := eventTypeTranslator(event.GetEvent())
		if err != nil {
			b.Logger.With(err).Error("ignoring game event")
			continue
		}

		marshal := jsonpb.Marshaler{
			OrigName:     true,
			EmitDefaults: true,
		}
		raw, err := marshal.MarshalToString(event)
		if err != nil {
			b.Logger.With(err).Error("error marshalling event message")
			continue
		}

		frameTime := time.Duration(b.gameSetup.ListeningDuration) * time.Millisecond
		remaining := time.Duration(b.gameSetup.GameDuration-event.GameSnapshot.Turn) * frameTime
		shotRemaining := time.Duration(0)
		if event.GameSnapshot.ShotClock != nil {
			shotRemaining = time.Duration(event.GameSnapshot.ShotClock.Turns) * frameTime
		}
		if eventType == app.EventBreakpoint {
			debugging = true
		} else if eventType == app.EventDebugReleased {
			debugging = false
		}

		update := app.FrontEndUpdate{
			Type: eventType,
			Update: app.UpdateData{
				GameEvent:     json.RawMessage(raw),
				TimeRemaining: fmt.Sprintf("%02d:%02d", int(remaining.Minutes()), int(remaining.Seconds())%60),
				ShotTime:      fmt.Sprintf("%02d", int(shotRemaining.Seconds())),
				DebugMode:     debugging,
			},
			ConnectionState: app.ConnStateUp,
		}
		b.lastUpdate = update
		b.sendToAll(update)
		if eventType == app.EventGameOver {
			// in this case we stop the connection before the server drop the broker
			return app.ErrGameOver
		}
	}
}

func (b *Binder) dropConsumer(uuid string) {
	b.consumersMux.Lock()
	defer b.consumersMux.Unlock()
	// this method may be called twice with the same argument because of concurrency between lock grant attempts
	_, stillThere := b.consumers[uuid]
	if stillThere {
		close(b.consumers[uuid])
		delete(b.consumers, uuid)
	}
}

func (b *Binder) dropAllConsumers() {
	for uuid := range b.consumers {
		b.dropConsumer(uuid)
	}
}

func (b *Binder) sendToAll(update app.FrontEndUpdate) {
	b.consumersMux.RLock()
	for uuid, consumer := range b.consumers {
		select {
		case consumer <- update:
		default:
			b.Logger.Warnf("consumer %s reached the max ignored messaged. Closing channel", uuid)
			go b.dropConsumer(uuid)
		}
	}
	b.consumersMux.RUnlock()
	b.Logger.Infof("event sent: %s", update.Type)
}

func (b *Binder) broadcastConnectionLost() {
	b.sendToAll(app.FrontEndUpdate{
		Type:            app.EventConnectionLost,
		Update:          b.lastUpdate.Update,
		ConnectionState: app.ConnStateDown,
	})
}

func (b *Binder) broadcastConnectionRees() {
	b.sendToAll(app.FrontEndUpdate{
		Type:            app.EventConnectionReestablished,
		Update:          b.lastUpdate.Update,
		ConnectionState: app.ConnStateUp,
	})
}

func eventTypeTranslator(event interface{}) (string, error) {
	switch event.(type) {
	case *lugo.GameEvent_NewPlayer:
		return app.EventNewPlayer, nil
	case *lugo.GameEvent_LostPlayer:
		return app.EventLostPlayer, nil
	case *lugo.GameEvent_StateChange:
		return app.EventStateChange, nil
	case *lugo.GameEvent_Goal:
		return app.EventGoal, nil
	case *lugo.GameEvent_Breakpoint:
		return app.EventBreakpoint, nil
	case *lugo.GameEvent_DebugReleased:
		return app.EventDebugReleased, nil
	case *lugo.GameEvent_GameOver:
		return app.EventGameOver, nil
	default:
		return "unknown", app.ErrUnknownGameEvent
	}
}
