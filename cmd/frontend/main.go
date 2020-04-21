package main

import (
	"bitbucket.org/makeitplay/lugo-frontend/web/app"
	"log"
	"net/http"
)

func main() {

	http.Handle("/", app.Newhandler("/Users/rubenssilva/go/src/bitbucket.org/makeitplay/lugo-frontend/web/"))

	log.Fatal(http.ListenAndServeTLS(":8080", "testdata/server.pem", "testdata/server.key", nil))

}
