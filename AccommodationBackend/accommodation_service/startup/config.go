package startup

import "os"

type Config struct {
	Port                string
	AccommodationDBHost string
	AccommodationDBPort string
}

func NewConfig() *Config {
	return &Config{
		Port:                os.Getenv("ACCOMMODATION_SERVICE_PORT"),
		AccommodationDBHost: os.Getenv("ACCOMMODATION_DB_HOST"),
		AccommodationDBPort: os.Getenv("ACCOMMODATION_DB_PORT"),
	}
}
