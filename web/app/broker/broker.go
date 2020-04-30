package broker

import (
	"bitbucket.org/makeitplay/lugo-frontend/web/app"
	"fmt"
	"github.com/lugobots/lugo4go/v2/lugo"
	"go.uber.org/zap"
	"io"
	"sync"
	"time"
)

var maxIgnoredMessaged = 20

func NewBinder(producer lugo.Broadcast_OnEventClient, initConfig app.Configuration, logger *zap.SugaredLogger) *Binder {
	return &Binder{
		consumers:    map[string]chan app.FrontEndUpdate{},
		consumersMux: sync.Mutex{},
		gameConfig:   initConfig,
		configMux:    sync.RWMutex{},
		Producer:     producer,
		Logger:       logger,
	}
}

type Binder struct {
	consumers    map[string]chan app.FrontEndUpdate
	consumersMux sync.Mutex
	gameConfig   app.Configuration
	configMux    sync.RWMutex
	Producer     lugo.Broadcast_OnEventClient
	Logger       *zap.SugaredLogger
}

func (b *Binder) StreamEventsTo(uuid string) chan app.FrontEndUpdate {
	b.consumersMux.Lock()
	defer b.consumersMux.Unlock()
	b.consumers[uuid] = make(chan app.FrontEndUpdate, maxIgnoredMessaged)
	return b.consumers[uuid]
}

func (b *Binder) GetGameConfig() app.Configuration {
	b.configMux.RLock()
	defer b.configMux.RUnlock()
	return b.gameConfig
}

func (b *Binder) Listen() error {
	for {
		event, err := b.Producer.Recv()
		if err != nil {
			if err != io.EOF {
				return fmt.Errorf("%w: %s", app.ErrGRPCConnectionLost, err)
			}
			return app.ErrGRPCConnectionClosed
		}
		b.configMux.Lock()
		newSnap := event.GameSnapshot
		remaining := time.Duration(b.gameConfig.GameDuration-event.GameSnapshot.Turn) * b.gameConfig.ListeningDuration
		b.gameConfig.TimeRemaining = fmt.Sprintf("%02d:%02d", int(remaining.Minutes()), int(remaining.Seconds()))
		b.gameConfig.HomeTeam.Score = newSnap.HomeTeam.Score
		b.gameConfig.AwayTeam.Score = newSnap.AwayTeam.Score

		eventType, err := eventTypeTranslator(event.GetEvent())
		if err != nil {
			b.Logger.With(err).Error("ignoring game event")
			continue
		}
		update := app.FrontEndUpdate{
			Type: eventType,
			Data: event,
		}
		for uuid, consumer := range b.consumers {
			select {
			case consumer <- update:
			default:
				b.Logger.Warnf("consumer %s reached the max ignored messaged. Closing channel", uuid)
				b.dropConsumer(uuid)
			}
		}
		b.configMux.Unlock()
	}
}

func (b *Binder) dropConsumer(uuid string) {
	b.consumersMux.Lock()
	defer b.consumersMux.Unlock()
	close(b.consumers[uuid])
	delete(b.consumers, uuid)
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
