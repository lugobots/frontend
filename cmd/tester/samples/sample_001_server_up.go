package samples

import (
	"github.com/lugobots/lugo4go/v2/lugo"
)

func SampleServerIsUp() Sample {
	var events []*lugo.GameEvent
	lastSnap := getInitSnap()
	events = append(events, newStateChangeEvent(lastSnap, lugo.GameSnapshot_WAITING))
	return Sample{
		Events: events,
		Setup: &lugo.GameSetup{
			DevMode:           false,
			StartMode:         lugo.GameSetup_DELAY,
			ListeningMode:     lugo.GameSetup_REMOTE,
			ListeningDuration: 50,
			GameDuration:      5 * 60 * (1000 / 50),
			HomeTeam: &lugo.TeamSettings{
				Name:   "Team A (setup)",
				Side:   lugo.Team_HOME,
				Avatar: "external/profile-team-home.jpg",
				Colors: &lugo.TeamColors{
					Primary: &lugo.TeamColor{
						Red:   240,
						Green: 0,
						Blue:  0,
					},
					Secondary: &lugo.TeamColor{
						Red:   255,
						Green: 255,
						Blue:  255,
					},
				},
			},
			AwayTeam: &lugo.TeamSettings{
				Name:   "Team B (setup)",
				Side:   lugo.Team_AWAY,
				Avatar: "external/profile-team-away.jpg",
				Colors: &lugo.TeamColors{
					Primary: &lugo.TeamColor{
						Red:   0,
						Green: 200,
						Blue:  0,
					},
					Secondary: &lugo.TeamColor{
						Red:   0,
						Green: 240,
						Blue:  240,
					},
				},
			},
		},
	}
}
