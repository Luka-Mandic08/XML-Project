package main

import (
	"api_gateway/startup"
	"api_gateway/startup/config"
	"log"
	"net/http"
)

func main() {
	config := config.NewConfig()
	server := startup.NewServer(config)
	if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		log.Fatalf("listen: %s\n", err)
	}
}
