package samples

import (
	"fmt"

	"github.com/golang/protobuf/proto"
	protoLugo "github.com/lugobots/lugo4go/v3/proto"
	"github.com/lugobots/lugo4go/v3/specs"
)

type Sample struct {
	Events []*protoLugo.GameEvent
	Setup  *protoLugo.GameSetup
}

func getLastSampleSnap(sample Sample) *protoLugo.GameSnapshot {
	return sample.Events[len(sample.Events)-1].GameSnapshot
}
func getInitSnap() *protoLugo.GameSnapshot {
	return &protoLugo.GameSnapshot{
		State: protoLugo.GameSnapshot_WAITING,
		Turn:  0,
		HomeTeam: &protoLugo.Team{
			Players: []*protoLugo.Player{},
			Name:    "Team C (snapshot)",
			Score:   0,
			Side:    protoLugo.Team_HOME,
		},
		AwayTeam: &protoLugo.Team{
			Players: []*protoLugo.Player{},
			Name:    "Team D (snapshot)",
			Score:   0,
			Side:    protoLugo.Team_AWAY,
		},
		Ball: &protoLugo.Ball{
			Position: &protoLugo.Point{},
		},
	}
}

func CopySnap(snap *protoLugo.GameSnapshot) *protoLugo.GameSnapshot {
	j, err := proto.Marshal(snap)
	if err != nil {
		panic(fmt.Sprintf("error marshalling snapshot: %s", err))
	}

	m := &protoLugo.GameSnapshot{}
	err = proto.UnmarshalMerge(j, m)
	if err != nil {
		panic(fmt.Sprintf("error marshalling snapshot: %s", err))
	}
	return m
}

func makeInitialPosition(playerNumber uint32, side protoLugo.Team_Side) *protoLugo.Point {
	p := protoLugo.Point{
		X: specs.FieldWidth / 4,
		Y: int32(playerNumber) * specs.PlayerSize * 2,
	}

	if side == protoLugo.Team_AWAY {
		p.X = specs.FieldWidth - p.X
	}
	return &p
}

func newStateChangeEvent(snap *protoLugo.GameSnapshot, previous protoLugo.GameSnapshot_State) *protoLugo.GameEvent {
	return &protoLugo.GameEvent{
		GameSnapshot: snap,
		Event: &protoLugo.GameEvent_StateChange{
			StateChange: &protoLugo.EventStateChange{
				PreviousState: previous,
				NewState:      snap.State,
			},
		},
	}
}

func newGoal(snap *protoLugo.GameSnapshot, side protoLugo.Team_Side) *protoLugo.GameEvent {
	return &protoLugo.GameEvent{
		GameSnapshot: snap,
		Event: &protoLugo.GameEvent_Goal{
			Goal: &protoLugo.EventGoal{
				Side: side,
			},
		},
	}
}
