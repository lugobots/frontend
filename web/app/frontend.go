package app

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/golang/protobuf/ptypes/empty"
	"html/template"
	"io"
	"log"
	"net/http"
	"path"
	"time"
)

func NewHandler(whereAmI, gameID string, srv EventsBroker) *gin.Engine {
	r := gin.Default()

	f := path.Join(whereAmI, "/static/dist/index.html")
	t, err := template.New("a").ParseFiles(f)
	if err != nil {
		panic(err)
	}

	r.SetHTMLTemplate(t)

	//r.Static("/js", path.Join(whereAmI, "/static/dist/js"))
	r.StaticFile("/loader.gif", path.Join(whereAmI, "/static/loader.gif"))
	r.StaticFile("/favicon.png", path.Join(whereAmI, "/static/favicon.png"))
	r.GET("/js/bundle.js", func(context *gin.Context) {
		//time.Sleep(5 * time.Second)
		context.File(path.Join(whereAmI, "/static/dist/js/bundle.js"))
	})

	r.Static("/images", path.Join(whereAmI, "/static/dist/images"))
	r.Static("/sounds", path.Join(whereAmI, "/static/dist/sounds"))
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

	remote := r.Group("/remote")
	{
		remote.POST("/pause-resume", func(context *gin.Context) {
			resp, err := srv.GetRemote().PauseOrResume(context, &empty.Empty{})
			if err != nil {
				context.JSON(http.StatusInternalServerError, err)
				return
			}
			context.JSON(http.StatusOK, resp)
		})
		remote.POST("/next-turn", func(context *gin.Context) {
			resp, err := srv.GetRemote().NextTurn(context, &empty.Empty{})
			if err != nil {
				context.JSON(http.StatusInternalServerError, err)
				return
			}
			context.JSON(http.StatusOK, resp)
		})
		remote.POST("/next-order", func(context *gin.Context) {
			resp, err := srv.GetRemote().NextTurn(context, &empty.Empty{})
			if err != nil {
				context.JSON(http.StatusInternalServerError, err)
				return
			}
			context.JSON(http.StatusOK, resp)
		})
	}

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
				return false
			case m, ok := <-streamChan:
				if !ok {
					return false
				}
				log.Printf("Sending type %s", m.Type)
				c.SSEvent(m.Type, m.Update)
			case <-time.After(500 * time.Millisecond):
				c.SSEvent("ping", "ping")
			}
			return true
		})
		log.Printf("finished streaming to %s", uuid)
	}
}
