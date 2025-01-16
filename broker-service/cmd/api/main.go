package main

import (
	"fmt"
	"log"
	"net/http"
)

const webPort = "80"

type Config struct {
}

func main() {
	// create app based on Config
	app := Config{}

	log.Printf("Starting broker service on port %s\n", webPort)

	// define HTTP server
	// A Server defines parameters for running an HTTP server
	srv := &http.Server{
		Addr:    fmt.Sprintf(":%s", webPort),
		Handler: app.routes(),
	}

	// start the server
	// ListenAndServe listens on srv.Addr and then calls
	// [Serve] to handle requests on incoming connections with srv.Handler
	err := srv.ListenAndServe()
	if err != nil {
		log.Panic(err)
	}
}
