package startup

import "github.com/tamararankovic/microservices_demo/inventory_service/domain"

var products = []*domain.Product{
	{
		ProductId: "623b0cc3a34d25d8567f9f82",
		ColorCode: "R",
		Quantity:  10,
	},
	{
		ProductId: "623b0cc3a34d25d8567f9f82",
		ColorCode: "B",
		Quantity:  12,
	},
	{
		ProductId: "623b0cc3a34d25d8567f9f83",
		ColorCode: "R",
		Quantity:  3,
	},
	{
		ProductId: "623b0cc3a34d25d8567f9f83",
		ColorCode: "G",
		Quantity:  7,
	},
}
