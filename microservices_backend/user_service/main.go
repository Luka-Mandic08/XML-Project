package main

import (
	"user_service/startup"
	cfg "user_service/startup/config"
)

func main() {
	config := cfg.NewConfig()
	server := startup.NewServer(config)
	server.Start()
}
