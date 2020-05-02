package samples

import (
	"github.com/lugobots/lugo4go/v2/lugo"
	"github.com/lugobots/lugo4go/v2/pkg/field"
)

func AllPlayersConnect() []*lugo.GameEvent {
	var events []*lugo.GameEvent

	lastSnap := getInitSnap()

	for i := uint32(1); i <= field.MaxPlayers; i++ {
		lastSnap = copySnap(lastSnap)

		newPlayer := &lugo.Player{
			Number:   i,
			TeamSide: lugo.Team_HOME,
			Position: makeInitialPosition(i, lugo.Team_HOME),
		}
		lastSnap.HomeTeam.Players = append(lastSnap.HomeTeam.Players, newPlayer)
		events = append(events, &lugo.GameEvent{
			GameSnapshot: lastSnap,
			Event: &lugo.GameEvent_NewPlayer{
				NewPlayer: &lugo.EventNewPlayer{
					Player: newPlayer,
				},
			},
		})
	}
	for i := uint32(1); i <= field.MaxPlayers; i++ {
		lastSnap = copySnap(lastSnap)

		newPlayer := &lugo.Player{
			Number:   i,
			TeamSide: lugo.Team_AWAY,
			Position: makeInitialPosition(i, lugo.Team_AWAY),
		}
		lastSnap.AwayTeam.Players = append(lastSnap.AwayTeam.Players, newPlayer)
		events = append(events, &lugo.GameEvent{
			GameSnapshot: lastSnap,
			Event: &lugo.GameEvent_NewPlayer{
				NewPlayer: &lugo.EventNewPlayer{
					Player: newPlayer,
				},
			},
		})
	}
	return events
}
