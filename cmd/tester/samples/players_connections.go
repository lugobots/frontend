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

func AllPlayersConnect() []*lugo.GameEvent {
	var events []*lugo.GameEvent

	lastSnap := getInitSnap()

	posHome := field.HomeTeamGoal().Center
	events = append(events, newPlayerEvent(lastSnap, &lugo.Player{
		Number:   1,
		TeamSide: lugo.Team_HOME,
		Position: &posHome,
	}))

	for i := uint32(2); i <= field.MaxPlayers; i++ {
		lastSnap = copySnap(lastSnap)

		newPlayer := &lugo.Player{
			Number:   i,
			TeamSide: lugo.Team_HOME,
			Position: makeInitialPosition(i, lugo.Team_HOME),
		}
		events = append(events, newPlayerEvent(lastSnap, newPlayer))
	}

	posAway := field.AwayTeamGoal().Center
	events = append(events, newPlayerEvent(lastSnap, &lugo.Player{
		Number:   1,
		TeamSide: lugo.Team_AWAY,
		Position: &posAway,
	}))
	for i := uint32(2); i <= field.MaxPlayers; i++ {
		lastSnap = copySnap(lastSnap)
		newPlayer := &lugo.Player{
			Number:   i,
			TeamSide: lugo.Team_AWAY,
			Position: makeInitialPosition(i, lugo.Team_AWAY),
		}
		events = append(events, newPlayerEvent(lastSnap, newPlayer))
	}

	return events
}
