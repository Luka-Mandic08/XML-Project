package main

import (
	"user_service/startup"
)

func main() {
	config := startup.NewConfig()
	server := startup.NewServer(config)
	server.Start()
}
