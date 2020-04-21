package app

import (
	"github.com/gin-gonic/gin"
	"html/template"
	"net/http"
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

func Newhandler() *gin.Engine {
	r := gin.Default()
	r.SetHTMLTemplate(html)
	v1 := r.Group("/")
	{
		v1.Static("/assets", "./assets")
		v1.GET("/", func(c *gin.Context) {
			//if pusher := c.Writer.Pusher(); pusher != nil {
			//	// use pusher.Push() to do server push
			//	if err := pusher.Push("/assets/app.js", nil); err != nil {
			//		log.Printf("Failed to push: %v", err)
			//	}
			//	go func() {
			//		time.Sleep(10 * time.Second)
			//		if err := pusher.Push("/assets/app2.js", nil); err != nil {
			//			log.Printf("Failed to push: %v", err)
			//		}
			//
			//	}()
			//}
			c.HTML(200, "https", gin.H{
				"status": "success",
			})
		})
	}

	return r
}

type Server struct {
}

func (s *Server) ServeHTTP(http.ResponseWriter, *http.Request) {

}
