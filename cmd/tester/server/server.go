package server

import (
	"context"
	"github.com/lugobots/lugo4go/v2/lugo"
	"go.uber.org/zap"
	"sync"
	"time"
)

func NewServer(logger *zap.SugaredLogger) *Broadcaster {
	return &Broadcaster{
		conns:   make([]lugo.Broadcast_OnEventServer, 0),
		connMux: sync.Mutex{},
		logger:  logger,
	}
}

type Broadcaster struct {
	conns      []lugo.Broadcast_OnEventServer
	connMux    sync.Mutex
	logger     *zap.SugaredLogger
	EventQueue []*lugo.GameEvent
	Setup      *lugo.GameSetup
}

func (b *Broadcaster) GetGameSetup(ctx context.Context, request *lugo.WatcherRequest) (*lugo.GameSetup, error) {
	return b.Setup, nil
}

func (b *Broadcaster) OnEvent(request *lugo.WatcherRequest, server lugo.Broadcast_OnEventServer) error {
	b.connMux.Lock()
	b.conns = append(b.conns, server)
	b.connMux.Unlock()
	b.logger.Infof("a new client")
	time.Sleep(5 * time.Second)

	b.logger.Infof("starting stream")
	for i, event := range b.EventQueue {
		b.logger.Infof("sending event %d", i)
		if err := server.Send(event); err != nil {
			b.logger.Errorf("error sending event %d: %s", i, err)
		}
		time.Sleep(50 * time.Millisecond)
	}

	<-server.Context().Done()

	return nil
}

//func (b *Broadcaster) SendEvent(e *lugo.GameEvent) {
//	for _, client := range b.conns {
//		if err := client.Send(e); err != nil {
//			b.logger.With(err).Error("did not sent message to a client")
//		}
//	}
//}
