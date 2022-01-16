package samples

import (
	"github.com/lugobots/lugo4go/v2/proto"
)

func SampleRotatePlayers() Sample {
	sample := SamplePlayersConnect()
	lastSnap := CopySnap(getLastSampleSnap(sample))

	playerTestIndex := 4

	dirs := []proto.Vector{
		proto.NorthEast(),
		proto.North(),
		proto.NorthWest(),
		proto.West(),
		proto.SouthWest(),
		proto.South(),
		proto.SouthEast(),
		proto.East(),
	}

	for _, d := range dirs {
		x := d
		lastSnap = CopySnap(lastSnap)
		lastSnap.HomeTeam.Players[playerTestIndex].Velocity.Direction = &x
		for i := 0; i < 20; i++ {
			sample.Events = append(sample.Events, newStateChangeEvent(lastSnap, proto.GameSnapshot_WAITING))
		}
	}

	return sample
}
