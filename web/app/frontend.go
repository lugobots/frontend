package app

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"html/template"
	"io"
	"log"
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

type Service struct {
}

func Newhandler(whereAmI string, srv Service) *gin.Engine {
	r := gin.Default()

	config := Configuration{
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

	r.GET("/", func(c *gin.Context) {
		//if pusher := c.Writer.Pusher(); pusher != nil {
		//	// use pusher.Push() to do server push
		//	if err := pusher.Push("/assets/app.js", nil); err != nil {
		//		log.Printf("Failed to push: %v", err)
		//	} else {
		//		log.Printf("PUUUUUSHED")
		//	}
		//}
		c.HTML(200, "index.html", gin.H{})
	})

	r.GET("/game-config", func(ctx *gin.Context) {

		ctx.JSON(200, config)
	})

	r.GET("/stream", func(c *gin.Context) {
		clientGone := c.Writer.CloseNotify()
		c.Stream(func(w io.Writer) bool {
			select {
			case <-clientGone:
				log.Println("Closed")
				return false
			case <-time.After(2 * time.Second):
				c.SSEvent("ping", "o que eh isso")
				log.Println("Send")

				if jjj, err := json.Marshal(config); err == nil {
					c.SSEvent("setup", string(jjj))
				}
				return true
			}
		})
	})
	return r
}
