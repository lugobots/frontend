package samples

import (
	"github.com/lugobots/lugo4go/v2/lugo"
	"github.com/lugobots/lugo4go/v2/pkg/field"
)

func SampleMovePlayers() []*lugo.GameEvent {
	events := SamplePlayersConnect()
	lastSnap := events[len(events)-1].GameSnapshot

	playerTestIndex := 4
	for lastSnap.HomeTeam.Players[playerTestIndex].Position.X < field.FieldWidth {
		lastSnap = copySnap(lastSnap)

		lastSnap.HomeTeam.Players[playerTestIndex].Position.X += field.PlayerMaxSpeed
		if lastSnap.HomeTeam.Players[playerTestIndex].Position.X > field.FieldWidth {
			lastSnap.HomeTeam.Players[playerTestIndex].Position.X = field.FieldWidth
		}
		events = append(events, newStateChangeEvent(lastSnap, lugo.GameSnapshot_WAITING))
	}

	for lastSnap.HomeTeam.Players[playerTestIndex].Position.Y < field.FieldHeight {
		lastSnap = copySnap(lastSnap)

		lastSnap.HomeTeam.Players[playerTestIndex].Position.Y += field.PlayerMaxSpeed
		if lastSnap.HomeTeam.Players[playerTestIndex].Position.Y > field.FieldHeight {
			lastSnap.HomeTeam.Players[playerTestIndex].Position.Y = field.FieldHeight
		}
		events = append(events, newStateChangeEvent(lastSnap, lugo.GameSnapshot_WAITING))
	}

	for lastSnap.HomeTeam.Players[playerTestIndex].Position.X > 0 {
		lastSnap = copySnap(lastSnap)

		lastSnap.HomeTeam.Players[playerTestIndex].Position.X -= field.PlayerMaxSpeed
		if lastSnap.HomeTeam.Players[playerTestIndex].Position.X < 0 {
			lastSnap.HomeTeam.Players[playerTestIndex].Position.X = 0
		}
		events = append(events, newStateChangeEvent(lastSnap, lugo.GameSnapshot_WAITING))
	}

	for lastSnap.HomeTeam.Players[playerTestIndex].Position.Y > 0 {
		lastSnap = copySnap(lastSnap)

		lastSnap.HomeTeam.Players[playerTestIndex].Position.Y -= field.PlayerMaxSpeed
		if lastSnap.HomeTeam.Players[playerTestIndex].Position.Y < 0 {
			lastSnap.HomeTeam.Players[playerTestIndex].Position.Y = 0
		}
		events = append(events, newStateChangeEvent(lastSnap, lugo.GameSnapshot_WAITING))
	}

	for lastSnap.HomeTeam.Players[playerTestIndex].Position.X < field.FieldWidth {
		lastSnap = copySnap(lastSnap)

		lastSnap.HomeTeam.Players[playerTestIndex].Position.X += field.PlayerMaxSpeed
		if lastSnap.HomeTeam.Players[playerTestIndex].Position.X > field.FieldWidth {
			lastSnap.HomeTeam.Players[playerTestIndex].Position.X = field.FieldWidth
		}
		events = append(events, newStateChangeEvent(lastSnap, lugo.GameSnapshot_WAITING))
	}

	events = append(events, newStateChangeEvent(lastSnap, lugo.GameSnapshot_WAITING))
	return events
}
