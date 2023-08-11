package startup

import "os"

type Config struct {
	Port              string
	ReservationDBHost string
	ReservationDBPort string
}

func NewConfig() *Config {
	return &Config{
		Port:              os.Getenv("RESERVATION_SERVICE_PORT"),
		ReservationDBHost: os.Getenv("RESERVATION_DB_HOST"),
		ReservationDBPort: os.Getenv("RESERVATION_DB_PORT"),
	}
}
