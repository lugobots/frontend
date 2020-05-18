package samples

import (
	"github.com/lugobots/lugo4go/v2/lugo"
	"github.com/lugobots/lugo4go/v2/pkg/field"
)

func newPlayerEvent(snap *lugo.GameSnapshot, player *lugo.Player) *lugo.GameEvent {
	if player.TeamSide == lugo.Team_HOME {
		snap.HomeTeam.Players = append(snap.HomeTeam.Players, player)
	} else {
		snap.AwayTeam.Players = append(snap.AwayTeam.Players, player)
	}

	return &lugo.GameEvent{
		GameSnapshot: snap,
		Event: &lugo.GameEvent_NewPlayer{
			NewPlayer: &lugo.EventNewPlayer{
				Player: player,
			},
		},
	}
}

func SamplePlayersConnect() Sample {
	sample := SampleServerIsUp()

	lastSnap := getLastSampleSnap(sample)

	posHome := field.HomeTeamGoal().Center
	lookingEast := lugo.NewZeroedVelocity(lugo.East())
	sample.Events = append(sample.Events, newPlayerEvent(lastSnap, &lugo.Player{
		Number:   1,
		TeamSide: lugo.Team_HOME,
		Position: &posHome,
		Velocity: &lookingEast,
	}))

	for i := uint32(2); i <= field.MaxPlayers; i++ {
		lastSnap = CopySnap(lastSnap)
		lookingEast.Speed = float64(i)
		newPlayer := &lugo.Player{
			Number:   i,
			TeamSide: lugo.Team_HOME,
			Position: makeInitialPosition(i, lugo.Team_HOME),
			Velocity: &lookingEast,
		}
		sample.Events = append(sample.Events, newPlayerEvent(lastSnap, newPlayer))
	}

	posAway := field.AwayTeamGoal().Center
	lookingWest := lugo.NewZeroedVelocity(lugo.West())
	sample.Events = append(sample.Events, newPlayerEvent(lastSnap, &lugo.Player{
		Number:   1,
		TeamSide: lugo.Team_AWAY,
		Position: &posAway,
		Velocity: &lookingWest,
	}))
	for i := uint32(2); i <= field.MaxPlayers; i++ {
		lastSnap = CopySnap(lastSnap)
		newPlayer := &lugo.Player{
			Number:   i,
			TeamSide: lugo.Team_AWAY,
			Position: makeInitialPosition(i, lugo.Team_AWAY),
			Velocity: &lookingWest,
		}
		sample.Events = append(sample.Events, newPlayerEvent(lastSnap, newPlayer))
	}

	return sample
}
