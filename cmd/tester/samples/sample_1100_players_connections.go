package samples

import (
	"github.com/lugobots/lugo4go/v3/field"
	"github.com/lugobots/lugo4go/v3/proto"
	"github.com/lugobots/lugo4go/v3/specs"
)

func newPlayerEvent(snap *proto.GameSnapshot, player *proto.Player) *proto.GameEvent {
	if player.TeamSide == proto.Team_HOME {
		snap.HomeTeam.Players = append(snap.HomeTeam.Players, player)
	} else {
		snap.AwayTeam.Players = append(snap.AwayTeam.Players, player)
	}

	return &proto.GameEvent{
		GameSnapshot: snap,
		Event: &proto.GameEvent_NewPlayer{
			NewPlayer: &proto.EventNewPlayer{
				Player: player,
			},
		},
	}
}

func SamplePlayersConnect() Sample {
	sample := SampleServerIsUp()

	lastSnap := getLastSampleSnap(sample)

	posHome := field.HomeSideTeamGoal().Center
	lookingEast := proto.NewZeroedVelocity(proto.East())
	sample.Events = append(sample.Events, newPlayerEvent(lastSnap, &proto.Player{
		Number:   1,
		TeamSide: proto.Team_HOME,
		Position: &posHome,
		Velocity: &lookingEast,
	}))

	for i := uint32(2); i <= specs.MaxPlayers; i++ {
		lastSnap = CopySnap(lastSnap)
		lookingEast.Speed = float64(i)
		newPlayer := &proto.Player{
			Number:   i,
			TeamSide: proto.Team_HOME,
			Position: makeInitialPosition(i, proto.Team_HOME),
			Velocity: &lookingEast,
		}
		sample.Events = append(sample.Events, newPlayerEvent(lastSnap, newPlayer))
	}

	posAway := field.AwaySideTeamGoal().Center
	lookingWest := proto.NewZeroedVelocity(proto.West())
	sample.Events = append(sample.Events, newPlayerEvent(lastSnap, &proto.Player{
		Number:   1,
		TeamSide: proto.Team_AWAY,
		Position: &posAway,
		Velocity: &lookingWest,
	}))
	for i := uint32(2); i <= specs.MaxPlayers; i++ {
		lastSnap = CopySnap(lastSnap)
		newPlayer := &proto.Player{
			Number:   i,
			TeamSide: proto.Team_AWAY,
			Position: makeInitialPosition(i, proto.Team_AWAY),
			Velocity: &lookingWest,
		}
		sample.Events = append(sample.Events, newPlayerEvent(lastSnap, newPlayer))
	}

	return sample
}
