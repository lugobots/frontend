package samples

import (
	"github.com/lugobots/lugo4go/v2/pkg/field"
	"github.com/lugobots/lugo4go/v2/proto"
)

func SampleMoveBall() Sample {
	sample := SampleServerIsUp()
	lastSnap := getLastSampleSnap(sample)

	lastSnap.Ball.Position = &proto.Point{}

	for lastSnap.Ball.Position.X < field.FieldWidth {
		lastSnap = CopySnap(lastSnap)

		lastSnap.Ball.Position.X += field.BallMaxSpeed
		if lastSnap.Ball.Position.X > field.FieldWidth {
			lastSnap.Ball.Position.X = field.FieldWidth
		}
		sample.Events = append(sample.Events, newStateChangeEvent(lastSnap, proto.GameSnapshot_WAITING))
	}

	for lastSnap.Ball.Position.Y < field.FieldHeight {
		lastSnap = CopySnap(lastSnap)

		lastSnap.Ball.Position.Y += field.BallMaxSpeed
		if lastSnap.Ball.Position.Y > field.FieldHeight {
			lastSnap.Ball.Position.Y = field.FieldHeight
		}
		sample.Events = append(sample.Events, newStateChangeEvent(lastSnap, proto.GameSnapshot_WAITING))
	}

	for lastSnap.Ball.Position.X > 0 {
		lastSnap = CopySnap(lastSnap)

		lastSnap.Ball.Position.X -= field.BallMaxSpeed
		if lastSnap.Ball.Position.X < 0 {
			lastSnap.Ball.Position.X = 0
		}
		sample.Events = append(sample.Events, newStateChangeEvent(lastSnap, proto.GameSnapshot_WAITING))
	}

	for lastSnap.Ball.Position.Y > 0 {
		lastSnap = CopySnap(lastSnap)

		lastSnap.Ball.Position.Y -= field.BallMaxSpeed
		if lastSnap.Ball.Position.Y < 0 {
			lastSnap.Ball.Position.Y = 0
		}
		sample.Events = append(sample.Events, newStateChangeEvent(lastSnap, proto.GameSnapshot_WAITING))
	}

	vec, _ := proto.NewVector(*lastSnap.Ball.Position, field.FieldCenter())
	vec.SetLength(field.BallMaxSpeed)
	for lastSnap.Ball.Position.DistanceTo(field.FieldCenter()) >= field.BallSize/2 {
		lastSnap = CopySnap(lastSnap)

		lastSnap.Ball.Position.X += int32(vec.X)
		lastSnap.Ball.Position.Y += int32(vec.Y)
		sample.Events = append(sample.Events, newStateChangeEvent(lastSnap, proto.GameSnapshot_WAITING))
	}

	*lastSnap.Ball.Position = field.FieldCenter()
	sample.Events = append(sample.Events, newStateChangeEvent(lastSnap, proto.GameSnapshot_WAITING))

	return sample
}
