package propagator

import (
	"context"
	"github.com/lugobots/frontend/web/app"
	"github.com/lugobots/lugo4go/v2/proto"
	"github.com/pkg/errors"
)

// Propagator is a Broadcast gRPC server that propagates the same events that are coming from the GameServer. However,
// the events are sent on the Frontend pace, that may be not the same as the backend due to different configurations,
// e.g. listening time, frontend animations, etc.
type Propagator struct {
	broker app.EventsBroker
}

func NewPropagator(broker app.EventsBroker) *Propagator {
	return &Propagator{
		broker: broker,
	}
}

func (p *Propagator) OnEvent(request *proto.WatcherRequest, server proto.Broadcast_OnEventServer) error {
	stream := p.broker.StreamEventsTo(request.Uuid)
	for {
		select {
		case <-server.Context().Done():
			return nil
		case e, ok := <-stream:
			if !ok {
				return errors.New("client disconnected")
			}
			err := server.Send(e.GameEvent)
			if err != nil {
				return errors.Wrap(err, "fail to propagate the game event")
			}
		}
	}
}

func (p *Propagator) GetGameSetup(ctx context.Context, request *proto.WatcherRequest) (*proto.GameSetup, error) {
	panic("not implemented in alpha version")
}

func (p *Propagator) StartGame(ctx context.Context, request *proto.StartRequest) (*proto.GameSetup, error) {
	return &proto.GameSetup{}, p.broker.StartGame(request.WatcherUuid)
}
