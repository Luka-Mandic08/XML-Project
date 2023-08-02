package startup

import "os"

type Config struct {
	Port       string
	UserDBHost string
	UserDBPort string
}

func NewConfig() *Config {
	return &Config{
		Port:       os.Getenv("USER_SERVICE_PORT"),
		UserDBHost: os.Getenv("USER_DB_HOST"),
		UserDBPort: os.Getenv("USER_DB_PORT"),
	}
}
