package samples

import (
	"fmt"
	"github.com/golang/protobuf/proto"
	"github.com/lugobots/lugo4go/v2/pkg/field"
	"github.com/lugobots/lugo4go/v2/proto"
)

type Sample struct {
	Events []*proto.GameEvent
	Setup  *proto.GameSetup
}

func getLastSampleSnap(sample Sample) *proto.GameSnapshot {
	return sample.Events[len(sample.Events)-1].GameSnapshot
}
func getInitSnap() *proto.GameSnapshot {
	return &proto.GameSnapshot{
		State: proto.GameSnapshot_WAITING,
		Turn:  0,
		HomeTeam: &proto.Team{
			Players: []*proto.Player{},
			Name:    "Team C (snapshot)",
			Score:   0,
			Side:    proto.Team_HOME,
		},
		AwayTeam: &proto.Team{
			Players: []*proto.Player{},
			Name:    "Team D (snapshot)",
			Score:   0,
			Side:    proto.Team_AWAY,
		},
		Ball: &proto.Ball{
			Position: &proto.Point{},
		},
	}
}

func CopySnap(snap *proto.GameSnapshot) *proto.GameSnapshot {
	j, err := proto.Marshal(snap)
	if err != nil {
		panic(fmt.Sprintf("error marshalling snapshot: %s", err))
	}

	m := &proto.GameSnapshot{}
	err = proto.UnmarshalMerge(j, m)
	if err != nil {
		panic(fmt.Sprintf("error marshalling snapshot: %s", err))
	}
	return m
}

func makeInitialPosition(playerNumber uint32, side proto.Team_Side) *proto.Point {
	p := proto.Point{
		X: field.FieldWidth / 4,
		Y: int32(playerNumber) * field.PlayerSize * 2,
	}

	if side == proto.Team_AWAY {
		p.X = field.FieldWidth - p.X
	}
	return &p
}

func newStateChangeEvent(snap *proto.GameSnapshot, previous proto.GameSnapshot_State) *proto.GameEvent {
	return &proto.GameEvent{
		GameSnapshot: snap,
		Event: &proto.GameEvent_StateChange{
			StateChange: &proto.EventStateChange{
				PreviousState: previous,
				NewState:      snap.State,
			},
		},
	}
}

func newGoal(snap *proto.GameSnapshot, side proto.Team_Side) *proto.GameEvent {
	return &proto.GameEvent{
		GameSnapshot: snap,
		Event: &proto.GameEvent_Goal{
			Goal: &proto.EventGoal{
				Side: side,
			},
		},
	}
}
