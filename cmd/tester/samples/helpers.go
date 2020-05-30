package samples

import (
	"fmt"
	"github.com/golang/protobuf/proto"
	"github.com/lugobots/lugo4go/v2/lugo"
	"github.com/lugobots/lugo4go/v2/pkg/field"
)

type Sample struct {
	Events []*lugo.GameEvent
	Setup  *lugo.GameSetup
}

func getLastSampleSnap(sample Sample) *lugo.GameSnapshot {
	return sample.Events[len(sample.Events)-1].GameSnapshot
}
func getInitSnap() *lugo.GameSnapshot {
	return &lugo.GameSnapshot{
		State: lugo.GameSnapshot_WAITING,
		Turn:  0,
		HomeTeam: &lugo.Team{
			Players: []*lugo.Player{},
			Name:    "Team C (snapshot)",
			Score:   0,
			Side:    lugo.Team_HOME,
		},
		AwayTeam: &lugo.Team{
			Players: []*lugo.Player{},
			Name:    "Team D (snapshot)",
			Score:   0,
			Side:    lugo.Team_AWAY,
		},
		Ball: &lugo.Ball{
			Position: &lugo.Point{},
		},
	}
}

func CopySnap(snap *lugo.GameSnapshot) *lugo.GameSnapshot {
	j, err := proto.Marshal(snap)
	if err != nil {
		panic(fmt.Sprintf("error marshalling snapshot: %s", err))
	}

	m := &lugo.GameSnapshot{}
	err = proto.UnmarshalMerge(j, m)
	if err != nil {
		panic(fmt.Sprintf("error marshalling snapshot: %s", err))
	}
	return m
}

func makeInitialPosition(playerNumber uint32, side lugo.Team_Side) *lugo.Point {
	p := lugo.Point{
		X: field.FieldWidth / 4,
		Y: int32(playerNumber) * field.PlayerSize * 2,
	}

	if side == lugo.Team_AWAY {
		p.X = field.FieldWidth - p.X
	}
	return &p
}

func newStateChangeEvent(snap *lugo.GameSnapshot, previous lugo.GameSnapshot_State) *lugo.GameEvent {
	return &lugo.GameEvent{
		GameSnapshot: snap,
		Event: &lugo.GameEvent_StateChange{
			StateChange: &lugo.EventStateChange{
				PreviousState: previous,
				NewState:      snap.State,
			},
		},
	}
}

func newGoal(snap *lugo.GameSnapshot, side lugo.Team_Side) *lugo.GameEvent {
	return &lugo.GameEvent{
		GameSnapshot: snap,
		Event: &lugo.GameEvent_Goal{
			Goal: &lugo.EventGoal{
				Side: side,
			},
		},
	}
}
