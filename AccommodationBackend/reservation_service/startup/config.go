package startup

import "os"

type Config struct {
	Port                            string
	ReservationDBHost               string
	ReservationDBPort               string
	NatsHost                        string
	NatsPort                        string
	NatsUser                        string
	NatsPass                        string
	CreateReservationCommandSubject string
	CreateReservationReplySubject   string
	AccommodationHost               string
	AccommodationPort               string
}

func NewConfig() *Config {
	return &Config{
		Port:                            os.Getenv("RESERVATION_SERVICE_PORT"),
		ReservationDBHost:               os.Getenv("RESERVATION_DB_HOST"),
		ReservationDBPort:               os.Getenv("RESERVATION_DB_PORT"),
		NatsHost:                        os.Getenv("NATS_HOST"),
		NatsPort:                        os.Getenv("NATS_PORT"),
		NatsUser:                        os.Getenv("NATS_USER"),
		NatsPass:                        os.Getenv("NATS_PASS"),
		CreateReservationCommandSubject: os.Getenv("CREATE_RESERVATION_COMMAND_SUBJECT"),
		CreateReservationReplySubject:   os.Getenv("CREATE_RESERVATION_REPLY_SUBJECT"),
		AccommodationHost:               os.Getenv("ACCOMMODATION_SERVICE_HOST"),
		AccommodationPort:               os.Getenv("ACCOMMODATION_SERVICE_PORT"),
	}
}
