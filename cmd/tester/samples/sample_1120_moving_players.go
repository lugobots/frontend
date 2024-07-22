package samples

import (
	"github.com/lugobots/lugo4go/v3/proto"
	"github.com/lugobots/lugo4go/v3/specs"
)

func SampleMovePlayers() Sample {
	sample := SamplePlayersConnect()
	lastSnap := getLastSampleSnap(sample)

	playerTestIndex := 4
	for lastSnap.HomeTeam.Players[playerTestIndex].Position.X < specs.FieldWidth {
		lastSnap = CopySnap(lastSnap)

		lastSnap.HomeTeam.Players[playerTestIndex].Position.X += specs.PlayerMaxSpeed
		if lastSnap.HomeTeam.Players[playerTestIndex].Position.X > specs.FieldWidth {
			lastSnap.HomeTeam.Players[playerTestIndex].Position.X = specs.FieldWidth
		}
		sample.Events = append(sample.Events, newStateChangeEvent(lastSnap, proto.GameSnapshot_WAITING))
	}

	for lastSnap.HomeTeam.Players[playerTestIndex].Position.Y < specs.FieldHeight {
		lastSnap = CopySnap(lastSnap)

		lastSnap.HomeTeam.Players[playerTestIndex].Position.Y += specs.PlayerMaxSpeed
		if lastSnap.HomeTeam.Players[playerTestIndex].Position.Y > specs.FieldHeight {
			lastSnap.HomeTeam.Players[playerTestIndex].Position.Y = specs.FieldHeight
		}
		sample.Events = append(sample.Events, newStateChangeEvent(lastSnap, proto.GameSnapshot_WAITING))
	}

	for lastSnap.HomeTeam.Players[playerTestIndex].Position.X > 0 {
		lastSnap = CopySnap(lastSnap)

		lastSnap.HomeTeam.Players[playerTestIndex].Position.X -= specs.PlayerMaxSpeed
		if lastSnap.HomeTeam.Players[playerTestIndex].Position.X < 0 {
			lastSnap.HomeTeam.Players[playerTestIndex].Position.X = 0
		}
		sample.Events = append(sample.Events, newStateChangeEvent(lastSnap, proto.GameSnapshot_WAITING))
	}

	for lastSnap.HomeTeam.Players[playerTestIndex].Position.Y > 0 {
		lastSnap = CopySnap(lastSnap)

		lastSnap.HomeTeam.Players[playerTestIndex].Position.Y -= specs.PlayerMaxSpeed
		if lastSnap.HomeTeam.Players[playerTestIndex].Position.Y < 0 {
			lastSnap.HomeTeam.Players[playerTestIndex].Position.Y = 0
		}
		sample.Events = append(sample.Events, newStateChangeEvent(lastSnap, proto.GameSnapshot_WAITING))
	}

	for lastSnap.HomeTeam.Players[playerTestIndex].Position.X < specs.FieldWidth {
		lastSnap = CopySnap(lastSnap)

		lastSnap.HomeTeam.Players[playerTestIndex].Position.X += specs.PlayerMaxSpeed
		if lastSnap.HomeTeam.Players[playerTestIndex].Position.X > specs.FieldWidth {
			lastSnap.HomeTeam.Players[playerTestIndex].Position.X = specs.FieldWidth
		}
		sample.Events = append(sample.Events, newStateChangeEvent(lastSnap, proto.GameSnapshot_WAITING))
	}

	sample.Events = append(sample.Events, newStateChangeEvent(lastSnap, proto.GameSnapshot_WAITING))
	return sample
}
