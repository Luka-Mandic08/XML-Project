package startup

import "os"

type Config struct {
	Port                            string
	UserDBHost                      string
	UserDBPort                      string
	NatsHost                        string
	NatsPort                        string
	NatsUser                        string
	NatsPass                        string
	CreateReservationCommandSubject string
	CreateReservationReplySubject   string
	ReservationHost                 string
	ReservationPort                 string
}

func NewConfig() *Config {
	return &Config{
		Port:                            os.Getenv("USER_SERVICE_PORT"),
		UserDBHost:                      os.Getenv("USER_DB_HOST"),
		UserDBPort:                      os.Getenv("USER_DB_PORT"),
		NatsHost:                        os.Getenv("NATS_HOST"),
		NatsPort:                        os.Getenv("NATS_PORT"),
		NatsUser:                        os.Getenv("NATS_USER"),
		NatsPass:                        os.Getenv("NATS_PASS"),
		CreateReservationCommandSubject: os.Getenv("CREATE_RESERVATION_COMMAND_SUBJECT"),
		CreateReservationReplySubject:   os.Getenv("CREATE_RESERVATION_REPLY_SUBJECT"),
		ReservationHost:                 os.Getenv("RESERVATION_SERVICE_HOST"),
		ReservationPort:                 os.Getenv("RESERVATION_SERVICE_PORT"),
	}
}
