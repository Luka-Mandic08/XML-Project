package config

import "os"

type Config struct {
	Host     string
	Port     string
	AuthHost string
	AuthPort string
	UserHost string
	UserPort string
}

func NewConfig() *Config {
	return &Config{
		Host:     os.Getenv("GATEWAY_HOST"),
		Port:     os.Getenv("GATEWAY_PORT"),
		AuthHost: os.Getenv("AUTH_SERVICE_HOST"),
		AuthPort: os.Getenv("AUTH_SERVICE_PORT"),
		UserHost: os.Getenv("USER_SERVICE_HOST"),
		UserPort: os.Getenv("USER_SERVICE_PORT"),
	}
}
