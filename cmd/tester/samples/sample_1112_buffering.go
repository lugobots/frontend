package samples

import (
	"github.com/lugobots/lugo4go/v3/proto"
)

func SampleBuffering() Sample {
	sample := SampleScoreTime()
	lastSnap := CopySnap(getLastSampleSnap(sample))

	lastSnap.State = proto.GameSnapshot_OVER
	sample.Events = append(sample.Events, &proto.GameEvent{
		GameSnapshot: lastSnap,
		Event: &proto.GameEvent_GameOver{
			GameOver: &proto.EventGameOver{},
		},
	})
	sample.Setup.GameDuration = lastSnap.Turn
	sample.Setup.ListeningDuration = 80
	return sample
}
