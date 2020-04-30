package server

import (
	"github.com/golang/protobuf/ptypes/empty"
	"github.com/lugobots/lugo4go/v2/lugo"
	"go.uber.org/zap"
	"sync"
)

func NewServer(logger *zap.SugaredLogger) *Broadcaster {
	return &Broadcaster{
		conns:   make([]lugo.Broadcast_OnEventServer, 0),
		connMux: sync.Mutex{},
		logger:  logger,
	}
}

type Broadcaster struct {
	conns   []lugo.Broadcast_OnEventServer
	connMux sync.Mutex
	logger  *zap.SugaredLogger
}

func (b *Broadcaster) OnEvent(_ *empty.Empty, server lugo.Broadcast_OnEventServer) error {
	b.connMux.Lock()
	b.conns = append(b.conns, server)
	b.connMux.Unlock()
	b.logger.Error("a new client")
	b.SendEvent(&lugo.GameEvent{
		GameSnapshot: &lugo.GameSnapshot{
			State: lugo.GameSnapshot_WAITING,
			Turn:  0,
			HomeTeam: &lugo.Team{
				Players: []*lugo.Player{},
				Name:    "AI",
				Score:   0,
				Side:    0,
			},
			AwayTeam: &lugo.Team{
				Players: []*lugo.Player{},
				Name:    "AI",
				Score:   0,
				Side:    1,
			},
			Ball: &lugo.Ball{},
		},
		Event: &lugo.GameEvent_NewPlayer{},
	})
	<-server.Context().Done()
	return nil
}

func (b *Broadcaster) SendEvent(e *lugo.GameEvent) {
	for _, client := range b.conns {
		if err := client.Send(e); err != nil {
			b.logger.With(err).Error("did not sent message to a client")
		}
	}
}
