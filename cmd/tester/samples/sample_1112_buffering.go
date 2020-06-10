package samples

import (
	"github.com/lugobots/lugo4go/v2/lugo"
)

func SampleBuffering() Sample {
	sample := SampleScoreTime()
	lastSnap := CopySnap(getLastSampleSnap(sample))

	lastSnap.State = lugo.GameSnapshot_OVER
	sample.Events = append(sample.Events, &lugo.GameEvent{
		GameSnapshot: lastSnap,
		Event: &lugo.GameEvent_GameOver{
			GameOver: &lugo.EventGameOver{},
		},
	})
	sample.Setup.GameDuration = lastSnap.Turn
	sample.Setup.ListeningDuration = 80
	return sample
}
