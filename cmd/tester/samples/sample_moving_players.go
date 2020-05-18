package samples

import (
	"github.com/lugobots/lugo4go/v2/lugo"
	"github.com/lugobots/lugo4go/v2/pkg/field"
)

func SampleMovePlayers() Sample {
	sample := SamplePlayersConnect()
	lastSnap := getLastSampleSnap(sample)

	playerTestIndex := 4
	for lastSnap.HomeTeam.Players[playerTestIndex].Position.X < field.FieldWidth {
		lastSnap = CopySnap(lastSnap)

		lastSnap.HomeTeam.Players[playerTestIndex].Position.X += field.PlayerMaxSpeed
		if lastSnap.HomeTeam.Players[playerTestIndex].Position.X > field.FieldWidth {
			lastSnap.HomeTeam.Players[playerTestIndex].Position.X = field.FieldWidth
		}
		sample.Events = append(sample.Events, newStateChangeEvent(lastSnap, lugo.GameSnapshot_WAITING))
	}

	for lastSnap.HomeTeam.Players[playerTestIndex].Position.Y < field.FieldHeight {
		lastSnap = CopySnap(lastSnap)

		lastSnap.HomeTeam.Players[playerTestIndex].Position.Y += field.PlayerMaxSpeed
		if lastSnap.HomeTeam.Players[playerTestIndex].Position.Y > field.FieldHeight {
			lastSnap.HomeTeam.Players[playerTestIndex].Position.Y = field.FieldHeight
		}
		sample.Events = append(sample.Events, newStateChangeEvent(lastSnap, lugo.GameSnapshot_WAITING))
	}

	for lastSnap.HomeTeam.Players[playerTestIndex].Position.X > 0 {
		lastSnap = CopySnap(lastSnap)

		lastSnap.HomeTeam.Players[playerTestIndex].Position.X -= field.PlayerMaxSpeed
		if lastSnap.HomeTeam.Players[playerTestIndex].Position.X < 0 {
			lastSnap.HomeTeam.Players[playerTestIndex].Position.X = 0
		}
		sample.Events = append(sample.Events, newStateChangeEvent(lastSnap, lugo.GameSnapshot_WAITING))
	}

	for lastSnap.HomeTeam.Players[playerTestIndex].Position.Y > 0 {
		lastSnap = CopySnap(lastSnap)

		lastSnap.HomeTeam.Players[playerTestIndex].Position.Y -= field.PlayerMaxSpeed
		if lastSnap.HomeTeam.Players[playerTestIndex].Position.Y < 0 {
			lastSnap.HomeTeam.Players[playerTestIndex].Position.Y = 0
		}
		sample.Events = append(sample.Events, newStateChangeEvent(lastSnap, lugo.GameSnapshot_WAITING))
	}

	for lastSnap.HomeTeam.Players[playerTestIndex].Position.X < field.FieldWidth {
		lastSnap = CopySnap(lastSnap)

		lastSnap.HomeTeam.Players[playerTestIndex].Position.X += field.PlayerMaxSpeed
		if lastSnap.HomeTeam.Players[playerTestIndex].Position.X > field.FieldWidth {
			lastSnap.HomeTeam.Players[playerTestIndex].Position.X = field.FieldWidth
		}
		sample.Events = append(sample.Events, newStateChangeEvent(lastSnap, lugo.GameSnapshot_WAITING))
	}

	sample.Events = append(sample.Events, newStateChangeEvent(lastSnap, lugo.GameSnapshot_WAITING))
	return sample
}
