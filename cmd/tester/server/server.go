package server

import (
	"context"
	"fmt"
	"sync"
	"time"

	lugo4go "github.com/lugobots/lugo4go/v3"
	"github.com/lugobots/lugo4go/v3/proto"
	"github.com/pkg/errors"
	"go.uber.org/zap"
)

func NewServer(logger *zap.SugaredLogger) *Broadcaster {
	return &Broadcaster{
		conns:   make([]proto.Broadcast_OnEventServer, 0),
		connMux: sync.Mutex{},
		logger:  logger,
	}
}

type Broadcaster struct {
	conns      []proto.Broadcast_OnEventServer
	connMux    sync.Mutex
	logger     *zap.SugaredLogger
	EventQueue []*proto.GameEvent
	Setup      *proto.GameSetup

	breakpoint chan bool
	lastSnap   *proto.GameSnapshot

	// gambiarra pra nao precisar fazer lista de conecoes. Nao vai funcionar se tiver mais de um front end connectado!
	shortcutHole chan *proto.GameEvent
}

func (b *Broadcaster) ResumeListeningPhase(ctx context.Context, request *proto.ResumeListeningRequest) (*proto.ResumeListeningResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (b *Broadcaster) StartGame(ctx context.Context, request *proto.StartRequest) (*proto.GameSetup, error) {
	//TODO implement me
	panic("implement me")
}

func (b *Broadcaster) PauseOrResume(ctx context.Context, _ *proto.PauseResumeRequest) (*proto.CommandResponse, error) {
	if b.breakpoint == nil {
		b.breakpoint = make(chan bool)
		b.shortcutHole <- &proto.GameEvent{
			GameSnapshot: b.lastSnap,
			Event: &proto.GameEvent_Breakpoint{
				Breakpoint: &proto.EventDebugBreakpoint{Breakpoint: proto.EventDebugBreakpoint_ORDERS},
			},
		}
	} else {
		b.shortcutHole <- &proto.GameEvent{
			GameSnapshot: b.lastSnap,
			Event: &proto.GameEvent_DebugReleased{
				DebugReleased: &proto.EventDebugReleased{},
			},
		}
		close(b.breakpoint)
		b.breakpoint = nil
	}
	return &proto.CommandResponse{
		Code:         proto.CommandResponse_SUCCESS,
		GameSnapshot: b.lastSnap,
		Details:      ":-)",
	}, nil
}

func (b *Broadcaster) NextTurn(ctx context.Context, empty *proto.NextTurnRequest) (*proto.CommandResponse, error) {
	return b.sendBreakpoint(proto.EventDebugBreakpoint_TURN)
}

func (b *Broadcaster) NextOrder(ctx context.Context, empty *proto.NextOrderRequest) (*proto.CommandResponse, error) {
	return b.sendBreakpoint(proto.EventDebugBreakpoint_ORDERS)
}

func (b *Broadcaster) sendBreakpoint(breakType proto.EventDebugBreakpoint_Breakpoint) (*proto.CommandResponse, error) {
	if b.breakpoint == nil {
		return &proto.CommandResponse{
			Code:         proto.CommandResponse_OTHER,
			GameSnapshot: b.lastSnap,
			Details:      ":-)",
		}, nil
	}
	b.shortcutHole <- &proto.GameEvent{
		GameSnapshot: b.lastSnap,
		Event: &proto.GameEvent_Breakpoint{
			Breakpoint: &proto.EventDebugBreakpoint{Breakpoint: breakType},
		},
	}
	close(b.breakpoint)
	b.breakpoint = make(chan bool)
	return &proto.CommandResponse{
		Code:         proto.CommandResponse_SUCCESS,
		GameSnapshot: b.lastSnap,
		Details:      ":-)",
	}, nil
}

func (b *Broadcaster) SetBallProperties(ctx context.Context, properties *proto.BallProperties) (*proto.CommandResponse, error) {
	panic("implement me")
}

func (b *Broadcaster) SetPlayerProperties(ctx context.Context, properties *proto.PlayerProperties) (*proto.CommandResponse, error) {
	inspector, err := lugo4go.NewGameSnapshotInspector(properties.Side, int(properties.Number), b.lastSnap)
	if err == nil {
		return nil, errors.Wrap(err, "failed to create inspector when setting player property")
	}

	p := inspector.GetMe()
	if p == nil {
		return nil, fmt.Errorf("player not found: %s-%d", properties.Side, properties.Number)
	}
	p.Position = properties.Position

	//b.logger.Infof("player %s-%d moved to %v", properties.Side, properties.Number, p.Position)
	return &proto.CommandResponse{
		Code:         proto.CommandResponse_SUCCESS,
		GameSnapshot: b.lastSnap,
		Details:      "player moved",
	}, nil
}

func (b *Broadcaster) SetGameProperties(ctx context.Context, properties *proto.GameProperties) (*proto.CommandResponse, error) {
	panic("implement me")
}

func (b *Broadcaster) GetGameSetup(ctx context.Context, request *proto.WatcherRequest) (*proto.GameSetup, error) {
	return b.Setup, nil
}

func (b *Broadcaster) OnEvent(request *proto.WatcherRequest, server proto.Broadcast_OnEventServer) error {
	b.connMux.Lock()
	b.conns = append(b.conns, server)
	b.connMux.Unlock()
	b.logger.Infof("a new client")
	time.Sleep(5 * time.Second)

	b.logger.Infof("starting stream")
	go func() {
		b.shortcutHole = make(chan *proto.GameEvent)
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
		b.logger.Infof("sending event %d (%s)", i, event.GameSnapshot.State)
		if err := server.Send(event); err != nil {
			b.logger.Errorf("error sending event %d: %s", i, err)
		}
		b.lastSnap = event.GameSnapshot
		time.Sleep(time.Duration(b.Setup.ListeningDuration) * time.Millisecond)
		//if event.GetGoal() != nil {
		//	b.logger.Infof("waiting GOAL time")
		//	time.Sleep(3 * time.Second)
		//}
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
