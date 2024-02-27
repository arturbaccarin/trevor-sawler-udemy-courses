package main

import (
	"log"
	"net/http"

	"github.com/alexedwards/scs/v2"
)

type application struct {
	Session *scs.SessionManager
}

func main() {
	// setup an app config
	app := application{}

	// get a session manager
	app.Session = getSession()

	// print out a message - app start
	log.Println("Starting server on port 8080...")

	// start the server
	err := http.ListenAndServe(":8080", app.routes())
	if err != nil {
		log.Fatal(err)
	}
}
