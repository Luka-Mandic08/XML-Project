package domain

import "time"

type Product struct {
	Id            string
	Name          string
	ClothingBrand string
	ColorCode     string
	ColorName     string
}

type OrderItem struct {
	Product  Product
	Quantity uint16
}

type OrderDetails struct {
	Id              string
	CreatedAt       time.Time
	Status          string
	ShippingAddress string
	ShippingStatus  string
	OrderItems      []OrderItem
}
