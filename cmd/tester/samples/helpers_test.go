package samples

import (
	"github.com/lugobots/lugo4go/v2/proto"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCopySnap(t *testing.T) {
	initial := getInitSnap()

	initial.Ball.Position.Y = 400
	initial.AwayTeam.Players = []*proto.Player{
		{
			Number: 1,
			Velocity: &proto.Velocity{
				Direction: &proto.Vector{
					X: 12,
					Y: 120,
				},
				Speed: 200,
			},
		},
	}

	newValue := CopySnap(initial)

	newValue.Turn = 200
	newValue.HomeTeam.Score = 200
	newValue.AwayTeam.Players[0].Number = 5
	newValue.AwayTeam.Players[0].Velocity.Direction.X = 400
	newValue.AwayTeam.Players[0].Velocity.Direction.Y = 100
	newValue.Ball.Position = &proto.Point{}

	assert.Equal(t, uint32(0), initial.Turn)
	assert.Equal(t, uint32(0), initial.HomeTeam.Score)
	assert.Equal(t, int32(400), initial.Ball.Position.Y)
	assert.Equal(t, uint32(1), initial.AwayTeam.Players[0].Number)
	assert.Equal(t, float64(12), initial.AwayTeam.Players[0].Velocity.Direction.X)

}
