package app

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"html/template"
	"io"
	"log"
	"net/http"
	"path"
	"time"

	"github.com/lugobots/lugo4go/v2/lugo"
)

type UpdateData struct {
	GameEvent     *lugo.GameEvent `json:"game_event"`
	TimeRemaining string          `json:"time_remaining"`
}
type FrontEndUpdate struct {
	Type   string     `json:"type"`
	Update UpdateData `json:"data"`
}

//type Service struct {
//	Clients []string
//}
//
//func (s *Service) StreamEventsTo(uuid string) chan FrontEndUpdate {
//
//	clientChan := make(chan FrontEndUpdate)
//
//	sn := &lugo.GameSnapshot{
//		State: lugo.GameSnapshot_WAITING,
//		Turn:  12,
//		HomeTeam: &lugo.Team{
//			Players: []*lugo.Player{{
//				Number: 1,
//				Position: &lugo.Point{
//					X: 100,
//					Y: 100,
//				},
//				Velocity:     nil,
//				TeamSide:     0,
//				InitPosition: nil,
//			},
//			},
//			Name:  "Eu",
//			Score: 0,
//			Side:  lugo.Team_HOME,
//		},
//		AwayTeam: &lugo.Team{
//			Players: []*lugo.Player{{
//				Number: 1,
//				Position: &lugo.Point{
//					X: 100,
//					Y: 100,
//				},
//				Velocity:     nil,
//				TeamSide:     0,
//				InitPosition: nil,
//			},
//			},
//			Name:  "Eu",
//			Score: 0,
//			Side:  lugo.Team_AWAY,
//		},
//		Ball:      &lugo.Ball{},
//		ShotClock: nil,
//	}
//
//	go func() {
//		for {
//			time.Sleep(1 * time.Second)
//			sn.Turn = uint32(time.Now().Second())
//			clientChan <- FrontEndUpdate{
//				Type: EventStateChange,
//				Update: UpdateData{
//					GameEvent: &lugo.GameEvent{
//						GameSnapshot: sn,
//						Event: &lugo.GameEvent_StateChange{StateChange: &lugo.EventStateChange{
//							PreviousState: lugo.GameSnapshot_LISTENING,
//							NewState:      lugo.GameSnapshot_PLAYING,
//						}},
//					},
//					TimeRemaining: time.Now().Format("i:s"),
//				},
//			}
//		}
//	}()
//	return clientChan
//}
//
//func (s *Service) GetGameConfig() Configuration {
//	return Configuration{
//		DevMode:   false,
//		StartMode: "",
//		HomeTeam: TeamConfiguration{
//			Name:   "My team",
//			Avatar: "external/profile-team-home.jpg",
//			Colors: TeamColors{
//				PrimaryColor: Color{
//					R: 255,
//					G: 255,
//				},
//				SecondaryColor: Color{
//					R: 100,
//					G: 100,
//				},
//			},
//		},
//		AwayTeam: TeamConfiguration{
//			Name:   "Other team",
//			Avatar: "external/profile-team-away.jpg",
//			Colors: TeamColors{
//				PrimaryColor: Color{
//					R: 100,
//					G: 255,
//				},
//				SecondaryColor: Color{
//					R: 100,
//					G: 200,
//					B: 50,
//				},
//			},
//		},
//	}
//}

func NewHandler(whereAmI, gameID string, srv EventsBroker) *gin.Engine {
	r := gin.Default()

	f := path.Join(whereAmI, "/static/dist/index.html")
	t, err := template.New("a").ParseFiles(f)
	if err != nil {
		panic(err)
	}

	r.SetHTMLTemplate(t)

	//r.Static("/js", path.Join(whereAmI, "/static/dist/js"))
	r.GET("/js/bundle.js", func(context *gin.Context) {
		time.Sleep(5 * time.Second)
		context.File(path.Join(whereAmI, "/static/dist/js/bundle.js"))
	})

	r.Static("/images", path.Join(whereAmI, "/static/dist/images"))
	r.Static("/external", path.Join(whereAmI, "/static/external"))

	//temp
	r.Static("/velho", path.Join(whereAmI, "/static/"))

	uniquer := Uniquer{}

	r.GET("/", func(c *gin.Context) {

		uuid, err := uniquer.New()
		if err != nil {
			c.String(http.StatusServiceUnavailable, "wow! Looks like we reached the limits of the connections. I am proud of it. Uh hoo!")
			return
		}

		uuid = fmt.Sprintf("%s_%d", uuid, time.Now().Nanosecond()%10000)

		c.HTML(http.StatusOK, "index.html", gin.H{
			"uuid":   uuid,
			"gameID": gameID,
		})
	})

	r.GET("/setup/:gameID/:uuid/", makeSetupHandler(srv))
	r.GET("/game-state/:gameID/:uuid/", makeGameStateHandler(srv))
	return r
}

func makeSetupHandler(srv EventsBroker) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(http.StatusOK, srv.GetGameConfig())
	}
}

func makeGameStateHandler(srv EventsBroker) gin.HandlerFunc {
	return func(c *gin.Context) {
		clientGone := c.Writer.CloseNotify()
		uuid := c.Param("uuid")
		streamChan := srv.StreamEventsTo(uuid)

		log.Printf("streaming to %s", uuid)
		c.Stream(func(w io.Writer) bool {
			select {
			case <-clientGone:
				log.Println("Closed")
				return false
			case m, ok := <-streamChan:
				if !ok {
					log.Println("channel closed")
					return false
				}
				log.Printf("Sending type %s: %s", m.Type, m.Update)
				c.SSEvent(m.Type, m.Update)
			case <-time.After(500 * time.Millisecond):
				c.SSEvent("ping", "ping")
			}
			return true
		})
		log.Printf("finished opening stream to %s", uuid)
	}
}
