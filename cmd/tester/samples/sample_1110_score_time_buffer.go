package samples

import (
	"github.com/lugobots/lugo4go/v3/proto"
)

func SampleScoreTime() Sample {
	sample := SamplePlayersConnect()
	lastSnap := getLastSampleSnap(sample)

	//time pass

	lastSnap.State = proto.GameSnapshot_LISTENING
	lastSnap.ShotClock = &proto.ShotClock{
		TeamSide:       proto.Team_AWAY,
		RemainingTurns: 300,
	}
	for i := 0; i < 200; i++ {
		lastSnap = CopySnap(lastSnap)
		lastSnap.Turn++
		lastSnap.ShotClock.RemainingTurns--
		sample.Events = append(sample.Events, newStateChangeEvent(lastSnap, proto.GameSnapshot_PLAYING))
	}

	lastSnap = CopySnap(lastSnap)
	lastSnap.AwayTeam.Score += 1
	lastSnap.State = proto.GameSnapshot_GET_READY
	lastSnap.ShotClock.TeamSide = proto.Team_HOME
	lastSnap.ShotClock.RemainingTurns = 300
	sample.Events = append(sample.Events, newGoal(lastSnap, proto.Team_AWAY))
	sample.Events = append(sample.Events, newStateChangeEvent(lastSnap, proto.GameSnapshot_PLAYING))

	lastSnap = CopySnap(lastSnap)
	lastSnap.State = proto.GameSnapshot_LISTENING
	for i := 0; i < 200; i++ {
		lastSnap = CopySnap(lastSnap)
		lastSnap.Turn++
		lastSnap.ShotClock.RemainingTurns--
		sample.Events = append(sample.Events, newStateChangeEvent(lastSnap, proto.GameSnapshot_PLAYING))
	}

	lastSnap = CopySnap(lastSnap)
	lastSnap.HomeTeam.Score += 1
	lastSnap.State = proto.GameSnapshot_GET_READY
	sample.Events = append(sample.Events, newGoal(lastSnap, proto.Team_HOME))
	sample.Events = append(sample.Events, newStateChangeEvent(lastSnap, proto.GameSnapshot_PLAYING))

	return sample
}
