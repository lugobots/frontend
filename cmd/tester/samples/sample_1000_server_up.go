package samples

import (
	"github.com/lugobots/lugo4go/v2/proto"
)

func SampleServerIsUp() Sample {
	var events []*proto.GameEvent
	lastSnap := getInitSnap()
	events = append(events, newStateChangeEvent(lastSnap, proto.GameSnapshot_WAITING))
	return Sample{
		Events: events,
		Setup: &proto.GameSetup{
			DevMode:           false,
			StartMode:         proto.GameSetup_DELAY,
			ListeningMode:     proto.GameSetup_REMOTE,
			ListeningDuration: 50,
			GameDuration:      5 * 60 * (1000 / 50),
			HomeTeam: &proto.TeamSettings{
				Name:   "Team A (setup)",
				Avatar: "external/profile-team-home.jpg",
				Colors: &proto.TeamColors{
					Primary: &proto.TeamColor{
						Red:   240,
						Green: 0,
						Blue:  0,
					},
					Secondary: &proto.TeamColor{
						Red:   255,
						Green: 255,
						Blue:  255,
					},
				},
			},
			AwayTeam: &proto.TeamSettings{
				Name:   "Team B (setup)",
				Avatar: "external/profile-team-away.jpg",
				Colors: &proto.TeamColors{
					Primary: &proto.TeamColor{
						Red:   0,
						Green: 200,
						Blue:  0,
					},
					Secondary: &proto.TeamColor{
						Red:   255,
						Green: 211,
						Blue:  0,
					},
				},
			},
		},
	}
}
