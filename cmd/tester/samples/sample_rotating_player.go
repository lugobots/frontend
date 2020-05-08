package samples

import (
	"github.com/lugobots/lugo4go/v2/lugo"
)

func SampleRotatePlayers() Sample {
	sample := SamplePlayersConnect()
	lastSnap := copySnap(getLastSampleSnap(sample))

	playerTestIndex := 4

	dirs := []lugo.Vector{
		lugo.NorthEast(),
		lugo.North(),
		lugo.NorthWest(),
		lugo.West(),
		lugo.SouthWest(),
		lugo.South(),
		lugo.SouthEast(),
		lugo.East(),
	}

	for _, d := range dirs {
		x := d
		lastSnap = copySnap(lastSnap)
		lastSnap.HomeTeam.Players[playerTestIndex].Velocity.Direction = &x
		for i := 0; i < 20; i++ {
			sample.Events = append(sample.Events, newStateChangeEvent(lastSnap, lugo.GameSnapshot_WAITING))
		}
	}

	return sample
}
