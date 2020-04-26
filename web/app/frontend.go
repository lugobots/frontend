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
)

var html = template.Must(template.New("https").Parse(`
<html>
<head>
  <title>Https Test</title>
  <script src="/assets/app.js"></script>
</head>
<body>
  <h1 style="color:red;">Welcome, Ginner!</h1>
</body>
</html>
`))

type FrontEndUpdate struct {
	Type string
	Data string
}

type Service struct {
	Clients []string
}

func (s *Service) StreamEventsTo(uuid string) chan FrontEndUpdate {

	clientChan := make(chan FrontEndUpdate)

	go func() {
		for {
			time.Sleep(1 * time.Second)
			clientChan <- FrontEndUpdate{
				Type: "ping",
				Data: "legal",
			}
		}
	}()
	return clientChan
}

func (s *Service) GetGameConfig() Configuration {
	return Configuration{
		DevMode:   false,
		StartMode: "",
		HomeTeam: TeamConfiguration{
			Name:   "My team",
			Avatar: "external/profile-team-home.jpg",
			Colors: map[string]Color{
				"a": {
					R: 255,
					G: 255,
				},
				"b": {
					R: 100,
					G: 100,
				},
			},
		},
		AwayTeam: TeamConfiguration{
			Name:   "Other team",
			Avatar: "external/profile-team-away.jpg",
			Colors: map[string]Color{
				"a": {
					R: 100,
					B: 255,
				},
				"b": {
					R: 150,
					B: 200,
				},
			},
		},
	}
}

func Newhandler(whereAmI, gameID string, srv *Service) *gin.Engine {
	r := gin.Default()

	f := path.Join(whereAmI, "/static/dist/index.html")
	t, err := template.New("a").ParseFiles(f)
	if err != nil {
		panic(err)
	}

	r.SetHTMLTemplate(t)
	r.Static("/assets", "./assets")
	r.Static("/js", path.Join(whereAmI, "/static/dist/js"))
	r.Static("/images", path.Join(whereAmI, "/static/dist/images"))
	r.Static("/external", path.Join(whereAmI, "/static/external"))

	uquiner := Uinquer{}

	r.GET("/", func(c *gin.Context) {

		uuid, err := uquiner.New()
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

	r.GET("/setup/:gameID/:uuid/", makeSetupHander(srv))
	r.GET("/game-state/:gameID/:uuid/", makeGameStateHander(srv))
	return r
}

func makeSetupHander(srv *Service) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(http.StatusOK, srv.GetGameConfig())
	}
}

func makeGameStateHander(srv *Service) gin.HandlerFunc {
	return func(c *gin.Context) {
		clientGone := c.Writer.CloseNotify()
		uuid := c.Param("uuid")
		c.Stream(func(w io.Writer) bool {
			select {
			case <-clientGone:
				log.Println("Closed")
				return false
			case m := <-srv.StreamEventsTo(uuid):
				log.Printf("Sending type %s: %s", m.Type, m.Data)
				c.SSEvent(m.Type, m.Data)
			}
			return true
		})
	}
}
