package config

import "os"

type Config struct {
	Port                      string
	InventoryDBHost           string
	InventoryDBPort           string
	InventoryDBName           string
	InventoryDBUser           string
	InventoryDBPass           string
	NatsHost                  string
	NatsPort                  string
	NatsUser                  string
	NatsPass                  string
	CreateOrderCommandSubject string
	CreateOrderReplySubject   string
}

func NewConfig() *Config {
	return &Config{
		Port:                      os.Getenv("INVENTORY_SERVICE_PORT"),
		InventoryDBHost:           os.Getenv("INVENTORY_DB_HOST"),
		InventoryDBPort:           os.Getenv("INVENTORY_DB_PORT"),
		InventoryDBName:           os.Getenv("INVENTORY_DB_NAME"),
		InventoryDBUser:           os.Getenv("INVENTORY_DB_USER"),
		InventoryDBPass:           os.Getenv("INVENTORY_DB_PASS"),
		NatsHost:                  os.Getenv("NATS_HOST"),
		NatsPort:                  os.Getenv("NATS_PORT"),
		NatsUser:                  os.Getenv("NATS_USER"),
		NatsPass:                  os.Getenv("NATS_PASS"),
		CreateOrderCommandSubject: os.Getenv("CREATE_ORDER_COMMAND_SUBJECT"),
		CreateOrderReplySubject:   os.Getenv("CREATE_ORDER_REPLY_SUBJECT"),
	}
}
