package main

import (
	"bitbucket.org/makeitplay/lugo-frontend/web/app"
	"bitbucket.org/makeitplay/lugo-frontend/web/app/broker"
	"log"
	"net/http"
)

func main() {
	srv := &broker.Binder{}

	server := app.Newhandler("/home/rubens/go/src/bitbucket.org/makeitplay/lugo-frontend/web", "local", srv)
	http.Handle("/", server)
	log.Fatal(http.ListenAndServeTLS(":8080",
		"testdata/server.pem",
		"testdata/server.key",
		nil))

}
