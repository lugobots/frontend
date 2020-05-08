package samples

import (
	"github.com/lugobots/lugo4go/v2/lugo"
)

func SampleScoreTime() Sample {
	sample := SamplePlayersConnect()
	lastSnap := getLastSampleSnap(sample)

	//time pass

	lastSnap.State = lugo.GameSnapshot_WAITING
	for i := 0; i < 100; i++ {
		lastSnap = copySnap(lastSnap)
		lastSnap.Turn++
		sample.Events = append(sample.Events, newStateChangeEvent(lastSnap, lugo.GameSnapshot_WAITING))
	}

	lastSnap.AwayTeam.Score += 1
	lastSnap.State = lugo.GameSnapshot_GET_READY
	sample.Events = append(sample.Events, newStateChangeEvent(lastSnap, lugo.GameSnapshot_WAITING))

	lastSnap.State = lugo.GameSnapshot_WAITING
	for i := 0; i < 100; i++ {
		lastSnap = copySnap(lastSnap)
		lastSnap.Turn++
		sample.Events = append(sample.Events, newStateChangeEvent(lastSnap, lugo.GameSnapshot_WAITING))
	}

	lastSnap = copySnap(lastSnap)
	lastSnap.AwayTeam.Score += 1
	lastSnap.State = lugo.GameSnapshot_GET_READY
	sample.Events = append(sample.Events, newStateChangeEvent(lastSnap, lugo.GameSnapshot_WAITING))

	return sample
}
