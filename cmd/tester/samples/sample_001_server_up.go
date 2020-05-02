package samples

import (
	"github.com/lugobots/lugo4go/v2/lugo"
)

func SampleServerIsUp() []*lugo.GameEvent {
	var events []*lugo.GameEvent
	lastSnap := getInitSnap()
	events = append(events, newStateChangeEvent(lastSnap, lugo.GameSnapshot_WAITING))
	return events
}
