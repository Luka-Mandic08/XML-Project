package main

import (
	"auth_service/startup"
)

func main() {
	config := startup.NewConfig()
	server := startup.NewServer(config)
	server.Start()
}
