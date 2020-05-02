package samples

import (
	"github.com/lugobots/lugo4go/v2/lugo"
	"github.com/lugobots/lugo4go/v2/pkg/field"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCopySnap(t *testing.T) {
	initial := getInitSnap()
	newValue := copySnap(initial)

	newValue.Turn = 200
	newValue.HomeTeam.Score = 200
	newValue.AwayTeam.Players = append(newValue.AwayTeam.Players, &lugo.Player{
		Number: 200,
	})
	newValue.Ball.Position = &lugo.Point{}
	*newValue.Ball.Position = field.AwayTeamGoal().Center

	assert.Equal(t, uint32(0), initial.Turn)
	assert.Equal(t, uint32(0), initial.HomeTeam.Score)
	assert.Len(t, initial.AwayTeam.Players, 0)
	assert.Nil(t, initial.Ball.Position)

}
