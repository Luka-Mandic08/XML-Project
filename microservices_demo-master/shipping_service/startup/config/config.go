package config

import "os"

type Config struct {
	Port                      string
	ShippingDBHost            string
	ShippingDBPort            string
	NatsHost                  string
	NatsPort                  string
	NatsUser                  string
	NatsPass                  string
	CreateOrderCommandSubject string
	CreateOrderReplySubject   string
}

func NewConfig() *Config {
	return &Config{
		Port:                      os.Getenv("SHIPPING_SERVICE_PORT"),
		ShippingDBHost:            os.Getenv("SHIPPING_DB_HOST"),
		ShippingDBPort:            os.Getenv("SHIPPING_DB_PORT"),
		NatsHost:                  os.Getenv("NATS_HOST"),
		NatsPort:                  os.Getenv("NATS_PORT"),
		NatsUser:                  os.Getenv("NATS_USER"),
		NatsPass:                  os.Getenv("NATS_PASS"),
		CreateOrderCommandSubject: os.Getenv("CREATE_ORDER_COMMAND_SUBJECT"),
		CreateOrderReplySubject:   os.Getenv("CREATE_ORDER_REPLY_SUBJECT"),
	}
}
