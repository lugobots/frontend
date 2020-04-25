package main

import (
	"bitbucket.org/makeitplay/lugo-frontend/web/app"
	"log"
	"net/http"
)

func main() {
	srv := app.Service{}
	server := app.Newhandler("/home/rubens/go/src/bitbucket.org/makeitplay/lugo-frontend/web", srv)
	http.Handle("/", server)
	log.Fatal(http.ListenAndServeTLS(":8080",
		"testdata/server.pem",
		"testdata/server.key",
		nil))

}
