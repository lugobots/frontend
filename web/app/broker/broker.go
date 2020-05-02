package broker

import (
	"bitbucket.org/makeitplay/lugo-frontend/web/app"
	"context"
	"fmt"
	"github.com/golang/protobuf/ptypes/empty"
	"github.com/lugobots/lugo4go/v2/lugo"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"io"
	"sync"
	"time"
)

var maxIgnoredMessaged = 20
var maxGRPCReconnect = 3
var grpcReconnectInterval = time.Second

const ErrMaxConnectionAttemptsReached = app.Error("did not connect to the game server")

func NewBinder(gameConfig app.Configuration, logger *zap.SugaredLogger) *Binder {
	return &Binder{
		consumers:    map[string]chan app.FrontEndUpdate{},
		consumersMux: sync.Mutex{},
		gameConfig:   gameConfig,
		configMux:    sync.RWMutex{},
		Logger:       logger,
	}
}

type Binder struct {
	consumers    map[string]chan app.FrontEndUpdate
	consumersMux sync.Mutex
	gameConfig   app.Configuration
	configMux    sync.RWMutex
	producerConn *grpc.ClientConn
	producer     lugo.BroadcastClient
	stopRequest  bool
	Logger       *zap.SugaredLogger
	lastUpdate   app.FrontEndUpdate
}

func (b *Binder) StreamEventsTo(uuid string) chan app.FrontEndUpdate {

	//clientChan := make(chan app.FrontEndUpdate)
	//
	//sn := &lugo.GameSnapshot{
	//	State: lugo.GameSnapshot_WAITING,
	//	Turn:  12,
	//	HomeTeam: &lugo.Team{
	//		Players: []*lugo.Player{{
	//			Number: 1,
	//			Position: &lugo.Point{
	//				X: 100,
	//				Y: 100,
	//			},
	//			Velocity:     nil,
	//			TeamSide:     0,
	//			InitPosition: nil,
	//		},
	//		},
	//		Name:  "Eu",
	//		Score: 0,
	//		Side:  lugo.Team_HOME,
	//	},
	//	AwayTeam: &lugo.Team{
	//		Players: []*lugo.Player{{
	//			Number: 1,
	//			Position: &lugo.Point{
	//				X: 100,
	//				Y: 100,
	//			},
	//			Velocity:     nil,
	//			TeamSide:     0,
	//			InitPosition: nil,
	//		},
	//		},
	//		Name:  "Eu",
	//		Score: 0,
	//		Side:  lugo.Team_AWAY,
	//	},
	//	Ball:      &lugo.Ball{},
	//	ShotClock: nil,
	//}
	//
	//go func() {
	//	for {
	//		time.Sleep(1 * time.Second)
	//		sn.Turn = uint32(time.Now().Second())
	//		clientChan <- app.FrontEndUpdate{
	//			Type: app.EventStateChange,
	//			Update: app.UpdateData{
	//				GameEvent: &lugo.GameEvent{
	//					GameSnapshot: sn,
	//					Event: &lugo.GameEvent_StateChange{StateChange: &lugo.EventStateChange{
	//						PreviousState: lugo.GameSnapshot_LISTENING,
	//						NewState:      lugo.GameSnapshot_PLAYING,
	//					}},
	//				},
	//				TimeRemaining: time.Now().Format("i:s"),
	//			},
	//		}
	//	}
	//}()
	//return clientChan

	b.consumersMux.Lock()
	defer b.consumersMux.Unlock()
	b.consumers[uuid] = make(chan app.FrontEndUpdate, maxIgnoredMessaged)
	b.consumers[uuid] <- b.lastUpdate
	return b.consumers[uuid]
}

func (b *Binder) GetGameConfig() app.Configuration {
	b.configMux.RLock()
	defer b.configMux.RUnlock()
	return b.gameConfig
}

func (b *Binder) connect() error {
	opts := []grpc.DialOption{grpc.WithBlock()}
	ctx, _ := context.WithTimeout(context.Background(), 3*time.Second)

	if b.gameConfig.Broadcast.Insecure {
		opts = append(opts, grpc.WithInsecure())
	}
	var err error
	b.producerConn, err = grpc.DialContext(ctx, b.gameConfig.Broadcast.Address, opts...)
	if err != nil {
		return err
	}

	b.producer = lugo.NewBroadcastClient(b.producerConn)
	return err
}

func (b *Binder) ListenAndBroadcast() error {
	tries := 0
	var finalErr error
	for tries < maxGRPCReconnect && !b.stopRequest {
		if err := b.connect(); err != nil {
			b.Logger.Warnf("failure on connecting to the game server: %s", err)
			time.Sleep(grpcReconnectInterval)
			tries++
		} else {
			tries = 0
			err := b.broadcast()
			if err == app.ErrGameOver {
				finalErr = err
				break
			}
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
	receiver, err := b.producer.OnEvent(ctx, &empty.Empty{})
	if err != nil {
		return err
	}
	b.Logger.Warn("starting broadcasting")
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
		remaining := time.Duration(b.gameConfig.GameDuration-event.GameSnapshot.Turn) * b.gameConfig.ListeningDuration
		update := app.FrontEndUpdate{
			Type: eventType,
			Update: app.UpdateData{
				GameEvent:     event,
				TimeRemaining: fmt.Sprintf("%02d:%02d", int(remaining.Minutes()), int(remaining.Seconds())),
			},
		}
		b.lastUpdate = update
		b.configMux.RLock()
		for uuid, consumer := range b.consumers {
			select {
			case consumer <- update:
			default:
				b.Logger.Warnf("consumer %s reached the max ignored messaged. Closing channel", uuid)
				b.dropConsumer(uuid)
			}
		}
		b.configMux.RUnlock()
		b.Logger.Infof("event sent: %s", eventType)
		if eventType == app.EventGameOver {
			// in this case we stop the connection before the server drop the broker
			return app.ErrGameOver
		}
	}
}

func (b *Binder) dropConsumer(uuid string) {
	b.consumersMux.Lock()
	defer b.consumersMux.Unlock()
	close(b.consumers[uuid])
	delete(b.consumers, uuid)
}

func (b *Binder) dropAllConsumers() {
	b.consumersMux.Lock()
	defer b.consumersMux.Unlock()
	for uuid, consumer := range b.consumers {
		close(consumer)
		delete(b.consumers, uuid)
	}
}

func eventTypeTranslator(event interface{}) (string, error) {
	switch event.(type) {
	case *lugo.GameEvent_NewPlayer:
		return app.EventNewPlayer, nil
	case *lugo.GameEvent_LostPlayer:
		return app.EventLostPlayer, nil
	case *lugo.GameEvent_StateChange:
		return app.EventStateChange, nil
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
