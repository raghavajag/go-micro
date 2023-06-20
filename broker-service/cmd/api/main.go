package main

import (
	"fmt"
	"log"
	"net/http"
)

type Config struct{}

const webPort = "8080"

func main() {
	log.Printf("Starting broker service on port %s", webPort)
	srv := &http.Server{
		Addr: fmt.Sprintf(":%s", webPort),
	}
	err := srv.ListenAndServe()
	if err != nil {
		log.Panic(err)
	}
}
