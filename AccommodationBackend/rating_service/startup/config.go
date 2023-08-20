package startup

import "os"

type Config struct {
	Port            string
	RatingDBHost    string
	RatingDBPort    string
	ReservationHost string
	ReservationPort string
}

func NewConfig() *Config {
	return &Config{
		Port:            os.Getenv("RATING_SERVICE_PORT"),
		RatingDBHost:    os.Getenv("RATING_DB_HOST"),
		RatingDBPort:    os.Getenv("RATING_DB_PORT"),
		ReservationHost: os.Getenv("RESERVATION_SERVICE_HOST"),
		ReservationPort: os.Getenv("RESERVATION_SERVICE_PORT"),
	}
}
