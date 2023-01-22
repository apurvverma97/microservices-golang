package main

import (
	"fmt"
	"log"
	"net/http"
)

const webPort = "9091"

type Config struct {}

func main() {
	app := Config{}

	log.Printf("Broker service is started on port %s\n", webPort)

	// define http server
	srv := &http.Server{
		Addr: fmt.Sprintf(":%s", webPort),
		Handler: app.routes(),
	}

	// start the server
	err := srv.ListenAndServe()

	if(err != nil ){
		log.Panic(err)
	}

}