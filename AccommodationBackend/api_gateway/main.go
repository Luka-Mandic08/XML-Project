package main

import (
	"api_gateway/startup"
	"log"
	"net/http"
)

func main() {
	config := startup.NewConfig()
	server := startup.NewServer(config)
	if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		log.Fatalf("listen: %s\n", err)
	}
}
