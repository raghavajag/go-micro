package main

import (
	"fmt"
	"log"
	"net/http"
)

const (
	webPort  = "8080"
	rpcPort  = "5001"
	mongoURL = "mongodb://mongo:27017"
	gRpcPort = "50001"
)

type Config struct {
}

func main() {
	log.Println("Starting Logger service")

	app := Config{}
	srv := &http.Server{
		Addr:    fmt.Sprintf(":%s", webPort),
		Handler: app.routes(),
	}
	err := srv.ListenAndServe()
	if err != nil {
		log.Panic()
	}
}
