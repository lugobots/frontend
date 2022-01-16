package broker

import (
	"bitbucket.org/makeitplay/lugo-frontend/web/app"
	"context"
	"encoding/json"
	"fmt"
	"github.com/golang/protobuf/jsonpb"
	"github.com/lugobots/lugo4go/v2/proto"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"io"
	"math"
	"sync"
	"time"
)

var maxIgnoredMessaged = 20
var maxGRPCReconnect = 10
var grpcReconnectInterval = time.Second

const ErrMaxConnectionAttemptsReached = app.Error("did not connect to the game server")

var defaultColor = &proto.TeamColor{}
var defaultTeamColors = &proto.TeamSettings{
	Colors: &proto.TeamColors{
		Primary:   defaultColor,
		Secondary: defaultColor,
	},
}

const MessagesRateMeasureTimeWindow = 5

func NewBinder(config app.Config, logger *zap.SugaredLogger, buffer BufferHandler) *Binder {
	return &Binder{
		consumers:    map[string]chan app.FrontEndUpdate{},
		consumersMux: sync.RWMutex{},
		config:       config,
		Logger:       logger,
		gameSetup: &proto.GameSetup{
			HomeTeam: defaultTeamColors,
			AwayTeam: defaultTeamColors,
		},
		buffer: buffer,
	}
}

const MaxUpdateBuffer = 1200 // n / 20 = time in sec

type Binder struct {
	consumers    map[string]chan app.FrontEndUpdate
	consumersMux sync.RWMutex
	config       app.Config
	gameSetup    *proto.GameSetup
	producerConn *grpc.ClientConn
	producer     proto.BroadcastClient
	remoteConn   proto.RemoteClient
	stopRequest  bool
	Logger       *zap.SugaredLogger
	lastUpdate   app.FrontEndUpdate
	buffer       BufferHandler
}

func (b *Binder) GetRemote() proto.RemoteClient {
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

	b.producer = proto.NewBroadcastClient(b.producerConn)
	b.gameSetup, err = b.producer.GetGameSetup(ctx, &proto.WatcherRequest{
		Uuid: "frontend",
	})
	if err != nil {
		return err
	}
	if !b.gameSetup.DevMode {
		bufferLoad := b.buffer.Start(func(data BufferedEvent) {
			b.lastUpdate = data.Update
			b.sendToAll(data.Update)
			if data.Update.Type == app.EventGoal {
				time.Sleep(5 * time.Second)
			} else if data.Update.Snapshot.State == proto.GameSnapshot_LISTENING {
				time.Sleep(50 * time.Millisecond)
			}
		}, b.gameSetup.GameDuration)
		go b.watchBufferNotifications(bufferLoad)
	}

	b.remoteConn = proto.NewRemoteClient(b.producerConn)

	return err
}

func (b *Binder) ListenAndBroadcast() error {
	tries := 0
	var finalErr error
	for !b.stopRequest && (b.config.StaysIfDisconnected || tries < maxGRPCReconnect) {
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
			b.buffer.Stop()
			b.broadcastConnectionLost()
			b.Logger.Warnf("broadcast interrupted: %s", err)
		}
	}
	if b.stopRequest {
		finalErr = app.ErrStopRequested
	}
	if !b.config.StaysIfDisconnected && tries >= maxGRPCReconnect {
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
	receiver, err := b.producer.OnEvent(ctx, &proto.WatcherRequest{
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
		var update app.FrontEndUpdate
		update, debugging, err = b.createFrame(event, debugging)
		if err != nil {
			b.Logger.Errorf("ignoring game event: %s", err)
			continue
		}
		if b.gameSetup.DevMode {
			b.lastUpdate = update
			b.sendToAll(update)
		} else {
			if err := b.buffer.QueueUp(update); err != nil {
				b.Logger.Errorf("ignoring game event: %s", err)
			}
			//if the game is ended, the grpx will be dropped soon, we must wait the buffer be consumed.
			if update.Type == app.EventGameOver {
				for range time.Tick(1 * time.Second) {
					b.Logger.Info("game ended, waiting buffer be consumed")
					// very ugly solution!
					if b.lastUpdate.Type == app.EventGameOver {
						return app.ErrGameOver
					}
				}
			}
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
	update := app.FrontEndUpdate{
		Type:            app.EventConnectionLost,
		Update:          b.lastUpdate.Update,
		ConnectionState: app.ConnStateDown,
	}
	b.lastUpdate = update
	b.sendToAll(update)
}

func (b *Binder) broadcastConnectionRees() {
	update := app.FrontEndUpdate{
		Type:            app.EventConnectionReestablished,
		Update:          b.lastUpdate.Update,
		ConnectionState: app.ConnStateUp,
	}
	b.lastUpdate = update
	b.sendToAll(update)
}

func (b *Binder) createFrame(event *proto.GameEvent, debugging bool) (app.FrontEndUpdate, bool, error) {
	eventType, err := eventTypeTranslator(event.GetEvent())
	if err != nil {
		return app.FrontEndUpdate{}, false, err
	}

	marshal := jsonpb.Marshaler{
		OrigName:     true,
		EmitDefaults: true,
	}
	raw, err := marshal.MarshalToString(event)
	if err != nil {
		return app.FrontEndUpdate{}, false, fmt.Errorf("error marshalling event message: %w", err)
	}

	frameTime := time.Duration(b.gameSetup.ListeningDuration) * time.Millisecond
	remaining := time.Duration(b.gameSetup.GameDuration-event.GameSnapshot.Turn) * frameTime
	shotRemaining := time.Duration(0)
	if event.GameSnapshot.ShotClock != nil {
		shotRemaining = time.Duration(event.GameSnapshot.ShotClock.RemainingTurns) * frameTime
	}

	if eventType == app.EventBreakpoint {
		debugging = true
	} else if eventType == app.EventDebugReleased {
		debugging = false
	}
	update := app.FrontEndUpdate{
		Type:     eventType,
		Snapshot: event.GameSnapshot,
		Update: app.UpdateData{
			GameEvent:     json.RawMessage(raw),
			TimeRemaining: fmt.Sprintf("%02d:%02d", int(remaining.Minutes()), int(remaining.Seconds())%60),
			ShotTime:      fmt.Sprintf("%02d", int(shotRemaining.Seconds())),
			DebugMode:     debugging,
		},
		ConnectionState: app.ConnStateUp,
	}
	return update, debugging, nil
}

func (b *Binder) watchBufferNotifications(bufferLoad <-chan float32) {
	for {
		select {
		case percentile, ok := <-bufferLoad:
			if !ok {
				return
			}
			b.lastUpdate.Update.Buffer = int(math.Round(float64(percentile) * 100))
			update := app.FrontEndUpdate{
				Update:          b.lastUpdate.Update,
				ConnectionState: app.ConnStateUp,
			}
			update.Type = app.EventBufferReady
			if percentile < 1 {
				update.Type = app.EventBuffering
			}
			b.lastUpdate = update
			b.sendToAll(update)
			b.Logger.Infof("Buffer load %f", percentile)
			time.Sleep(5 * time.Second)
		}
	}
}

func eventTypeTranslator(event interface{}) (app.EventType, error) {
	switch event.(type) {
	case *proto.GameEvent_NewPlayer:
		return app.EventNewPlayer, nil
	case *proto.GameEvent_LostPlayer:
		return app.EventLostPlayer, nil
	case *proto.GameEvent_StateChange:
		return app.EventStateChange, nil
	case *proto.GameEvent_Goal:
		return app.EventGoal, nil
	case *proto.GameEvent_Breakpoint:
		return app.EventBreakpoint, nil
	case *proto.GameEvent_DebugReleased:
		return app.EventDebugReleased, nil
	case *proto.GameEvent_GameOver:
		return app.EventGameOver, nil
	default:
		return "unknown", app.ErrUnknownGameEvent
	}
}
