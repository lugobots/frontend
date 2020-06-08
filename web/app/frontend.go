package app

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/golang/protobuf/ptypes/empty"
	"github.com/lugobots/lugo4go/v2/lugo"
	"html/template"
	"io"
	"log"
	"net/http"
	"path"
	"strconv"
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
		remote.PATCH("/players/:team/:number", func(context *gin.Context) {
			side := lugo.Team_HOME
			if context.Param("team") == "away" {
				side = lugo.Team_AWAY
			}
			n, err := strconv.Atoi(context.Param("number"))
			if err != nil {
				e := fmt.Errorf("invalid player number %s: %w", context.Param("number"), err)
				log.Println(e)
				context.JSON(http.StatusInternalServerError, e)
				return
			}

			props, err := context.GetRawData()
			if err != nil {
				e := fmt.Errorf("error reading playload: %w", err)
				log.Println(e)
				context.JSON(http.StatusInternalServerError, e)
				return
			}
			playerProperties := &lugo.PlayerProperties{
				Side:     side,
				Number:   uint32(n),
				Position: &lugo.Point{},
				Velocity: nil,
			}

			if err := json.Unmarshal(props, playerProperties.Position); err != nil {
				e := fmt.Errorf("not a valid JSON: %w", err)
				log.Println(e)
				context.JSON(http.StatusInternalServerError, e)
				return
			}

			resp, err := srv.GetRemote().SetPlayerProperties(context, playerProperties)
			if err != nil {
				e := fmt.Errorf("error from game server: %w", err)
				log.Println(e)
				context.JSON(http.StatusInternalServerError, e)
				return
			}
			context.JSON(http.StatusOK, resp)
		})
	}

	return r
}

func makeSetupHandler(srv EventsBroker) gin.HandlerFunc {
	return func(c *gin.Context) {
		if config, err := srv.GetGameConfig(c.Param("uuid")); err == nil {
			c.JSON(http.StatusOK, config)
		} else {
			c.JSON(http.StatusBadGateway, map[string]interface{}{"error": err.Error()})
		}
	}
}

func makeGameStateHandler(srv EventsBroker) gin.HandlerFunc {
	return func(c *gin.Context) {
		clientGone := c.Writer.CloseNotify()
		uuid := c.Param("uuid")
		streamChan := srv.StreamEventsTo(uuid)

		//
		//Problema para resolver:
		//	O backend e o frontend estao perdendo sincronia quando a conectao eh aberta e o backend esta no mobo debug
		//
		//
		//isso acontece pq o frontend processa o SETUP depois que a conecao com o stream ja enviou o ultimo quadro
		//que diz que o modo debug is on.
		//	I o frontend manda pra o "listening" depois de fazer o setup>
		//
		//	opcao 1: mandar um novo frame assim que o cara pede o setup
		//opcao 2: fazer endpint pra solicitar ultimo evento
		//opcao 3:  no frontend mudar a logic pra o setup nao definir proximo estado do frontend

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
				c.SSEvent(string(m.Type), m.Update)
			case <-time.After(500 * time.Millisecond):
				c.SSEvent("ping", "ping")
			}
			return true
		})
		log.Printf("finished streaming to %s", uuid)
	}
}
