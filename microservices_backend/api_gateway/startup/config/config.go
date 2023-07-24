package config

import "os"

type Config struct {
	Port          string
	CatalogueHost string
	CataloguePort string
	OrderingHost  string
	OrderingPort  string
	ShippingHost  string
	ShippingPort  string
	InventoryHost string
	InventoryPort string
}

func NewConfig() *Config {
	return &Config{
		Port:          os.Getenv("GATEWAY_PORT"),
		CatalogueHost: os.Getenv("CATALOGUE_SERVICE_HOST"),
		CataloguePort: os.Getenv("CATALOGUE_SERVICE_PORT"),
		OrderingHost:  os.Getenv("ORDERING_SERVICE_HOST"),
		OrderingPort:  os.Getenv("ORDERING_SERVICE_PORT"),
		ShippingHost:  os.Getenv("SHIPPING_SERVICE_HOST"),
		ShippingPort:  os.Getenv("SHIPPING_SERVICE_PORT"),
		InventoryHost: os.Getenv("INVENTORY_SERVICE_HOST"),
		InventoryPort: os.Getenv("INVENTORY_SERVICE_PORT"),
	}
}
