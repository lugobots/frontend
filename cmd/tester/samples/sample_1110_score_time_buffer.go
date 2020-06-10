package samples

import (
	"github.com/lugobots/lugo4go/v2/lugo"
)

func SampleScoreTime() Sample {
	sample := SamplePlayersConnect()
	lastSnap := getLastSampleSnap(sample)

	//time pass

	lastSnap.State = lugo.GameSnapshot_LISTENING
	lastSnap.ShotClock = &lugo.ShotClock{
		TeamSide: lugo.Team_AWAY,
		Turns:    300,
	}
	for i := 0; i < 200; i++ {
		lastSnap = CopySnap(lastSnap)
		lastSnap.Turn++
		lastSnap.ShotClock.Turns--
		sample.Events = append(sample.Events, newStateChangeEvent(lastSnap, lugo.GameSnapshot_PLAYING))
	}

	lastSnap = CopySnap(lastSnap)
	lastSnap.AwayTeam.Score += 1
	lastSnap.State = lugo.GameSnapshot_GET_READY
	lastSnap.ShotClock.TeamSide = lugo.Team_HOME
	lastSnap.ShotClock.Turns = 300
	sample.Events = append(sample.Events, newGoal(lastSnap, lugo.Team_AWAY))
	sample.Events = append(sample.Events, newStateChangeEvent(lastSnap, lugo.GameSnapshot_PLAYING))

	lastSnap = CopySnap(lastSnap)
	lastSnap.State = lugo.GameSnapshot_LISTENING
	for i := 0; i < 200; i++ {
		lastSnap = CopySnap(lastSnap)
		lastSnap.Turn++
		lastSnap.ShotClock.Turns--
		sample.Events = append(sample.Events, newStateChangeEvent(lastSnap, lugo.GameSnapshot_PLAYING))
	}

	lastSnap = CopySnap(lastSnap)
	lastSnap.HomeTeam.Score += 1
	lastSnap.State = lugo.GameSnapshot_GET_READY
	sample.Events = append(sample.Events, newGoal(lastSnap, lugo.Team_HOME))
	sample.Events = append(sample.Events, newStateChangeEvent(lastSnap, lugo.GameSnapshot_PLAYING))

	return sample
}
