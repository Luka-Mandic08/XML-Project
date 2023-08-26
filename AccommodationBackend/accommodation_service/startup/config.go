package startup

import "os"

type Config struct {
	Port                            string
	AccommodationDBHost             string
	AccommodationDBPort             string
	NatsHost                        string
	NatsPort                        string
	NatsUser                        string
	NatsPass                        string
	CreateReservationCommandSubject string
	CreateReservationReplySubject   string
	RatingHost                      string
	RatingPort                      string
	ReservationHost                 string
	ReservationPort                 string
}

func NewConfig() *Config {
	return &Config{
		Port:                            os.Getenv("ACCOMMODATION_SERVICE_PORT"),
		AccommodationDBHost:             os.Getenv("ACCOMMODATION_DB_HOST"),
		AccommodationDBPort:             os.Getenv("ACCOMMODATION_DB_PORT"),
		NatsHost:                        os.Getenv("NATS_HOST"),
		NatsPort:                        os.Getenv("NATS_PORT"),
		NatsUser:                        os.Getenv("NATS_USER"),
		NatsPass:                        os.Getenv("NATS_PASS"),
		CreateReservationCommandSubject: os.Getenv("CREATE_RESERVATION_COMMAND_SUBJECT"),
		CreateReservationReplySubject:   os.Getenv("CREATE_RESERVATION_REPLY_SUBJECT"),
		RatingHost:                      os.Getenv("RATING_SERVICE_HOST"),
		RatingPort:                      os.Getenv("RATING_SERVICE_PORT"),
		ReservationHost:                 os.Getenv("RESERVATION_SERVICE_HOST"),
		ReservationPort:                 os.Getenv("RESERVATION_SERVICE_PORT"),
	}
}
