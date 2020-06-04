package server

import (
	"context"
	"fmt"
	"github.com/golang/protobuf/ptypes/empty"
	"github.com/lugobots/lugo4go/v2/lugo"
	"github.com/lugobots/lugo4go/v2/pkg/field"
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

	breakpoint chan bool
	lastSnap   *lugo.GameSnapshot

	// gambiarra pra nao precisar fazer lista de conecoes. Nao vai funcionar se tiver mais de um front end connectado!
	shortcutHole chan *lugo.GameEvent
}

func (b *Broadcaster) PauseOrResume(ctx context.Context, empty *empty.Empty) (*lugo.CommandResponse, error) {
	if b.breakpoint == nil {
		b.breakpoint = make(chan bool)
		b.shortcutHole <- &lugo.GameEvent{
			GameSnapshot: b.lastSnap,
			Event: &lugo.GameEvent_Breakpoint{
				Breakpoint: &lugo.EventDebugBreakpoint{Breakpoint: lugo.EventDebugBreakpoint_ORDERS},
			},
		}
	} else {
		b.shortcutHole <- &lugo.GameEvent{
			GameSnapshot: b.lastSnap,
			Event: &lugo.GameEvent_DebugReleased{
				DebugReleased: &lugo.EventDebugReleased{},
			},
		}
		close(b.breakpoint)
		b.breakpoint = nil
	}
	return &lugo.CommandResponse{
		Code:         lugo.CommandResponse_SUCCESS,
		GameSnapshot: b.lastSnap,
		Details:      ":-)",
	}, nil
}

func (b *Broadcaster) NextTurn(ctx context.Context, empty *empty.Empty) (*lugo.CommandResponse, error) {
	return b.sendBreakpoint(lugo.EventDebugBreakpoint_TURN)
}

func (b *Broadcaster) NextOrder(ctx context.Context, empty *empty.Empty) (*lugo.CommandResponse, error) {
	return b.sendBreakpoint(lugo.EventDebugBreakpoint_ORDERS)
}

func (b *Broadcaster) sendBreakpoint(breakType lugo.EventDebugBreakpoint_Breakpoint) (*lugo.CommandResponse, error) {
	if b.breakpoint == nil {
		return &lugo.CommandResponse{
			Code:         lugo.CommandResponse_OTHER,
			GameSnapshot: b.lastSnap,
			Details:      ":-)",
		}, nil
	}
	b.shortcutHole <- &lugo.GameEvent{
		GameSnapshot: b.lastSnap,
		Event: &lugo.GameEvent_Breakpoint{
			Breakpoint: &lugo.EventDebugBreakpoint{Breakpoint: breakType},
		},
	}
	close(b.breakpoint)
	b.breakpoint = make(chan bool)
	return &lugo.CommandResponse{
		Code:         lugo.CommandResponse_SUCCESS,
		GameSnapshot: b.lastSnap,
		Details:      ":-)",
	}, nil
}

func (b *Broadcaster) SetBallProperties(ctx context.Context, properties *lugo.BallProperties) (*lugo.CommandResponse, error) {
	panic("implement me")
}

func (b *Broadcaster) SetPlayerProperties(ctx context.Context, properties *lugo.PlayerProperties) (*lugo.CommandResponse, error) {
	p := field.GetPlayer(b.lastSnap, properties.Side, properties.Number)
	if p == nil {
		return nil, fmt.Errorf("player not found: %s-%d", properties.Side, properties.Number)
	}
	p.Position = properties.Position

	b.logger.Infof("player %s-%d moved to %v", properties.Side, properties.Number, p.Position)
	return &lugo.CommandResponse{
		Code:         lugo.CommandResponse_SUCCESS,
		GameSnapshot: b.lastSnap,
		Details:      "player moved",
	}, nil
}

func (b *Broadcaster) SetGameProperties(ctx context.Context, properties *lugo.GameProperties) (*lugo.CommandResponse, error) {
	panic("implement me")
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
	go func() {
		b.shortcutHole = make(chan *lugo.GameEvent)
		for {
			v, ok := <-b.shortcutHole
			if !ok {
				return
			}
			if err := server.Send(v); err != nil {
				b.logger.Errorf("error sending event through the hole: %s", err)
			}
		}
	}()
	for i, event := range b.EventQueue {
		if b.breakpoint != nil {
			<-b.breakpoint
		}
		b.logger.Infof("sending event %d", i)
		if err := server.Send(event); err != nil {
			b.logger.Errorf("error sending event %d: %s", i, err)
		}
		b.lastSnap = event.GameSnapshot
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
