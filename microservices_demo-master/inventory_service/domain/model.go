package domain

type Product struct {
	ProductId string `gorm:"index:idx_name,unique"`
	ColorCode string `gorm:"index:idx_name,unique"`
	Quantity  uint64
}
