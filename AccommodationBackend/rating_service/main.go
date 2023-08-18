package main

import (
	"rating_service/startup"
)

func main() {
	config := startup.NewConfig()
	server := startup.NewServer(config)
	server.Start()
}
