package config

import "os"

type Config struct {
	Address string
}

func GetConfig() Config {
	return Config{
		Address: os.Getenv("USER_SERVICE_ADDRESS"),
	}
}
