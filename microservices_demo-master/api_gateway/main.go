package main

import (
	"github.com/tamararankovic/microservices_demo/api_gateway/startup"
	"github.com/tamararankovic/microservices_demo/api_gateway/startup/config"
)

func main() {
	config := config.NewConfig()
	server := startup.NewServer(config)
	server.Start()
}
