package startup

import "os"

type Config struct {
	Port       string
	AuthDBHost string
	AuthDBPort string
}

func NewConfig() *Config {
	return &Config{
		Port:       os.Getenv("AUTH_SERVICE_PORT"),
		AuthDBHost: os.Getenv("AUTH_DB_HOST"),
		AuthDBPort: os.Getenv("AUTH_DB_PORT"),
	}
}
