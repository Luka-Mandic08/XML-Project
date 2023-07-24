package domain

type ProductStore interface {
	Insert(product *Product) error
	UpdateQuantity(product *Product, difference int64) error
	UpdateQuantityForAll(products map[*Product]int64) error
	GetAll() (*[]Product, error)
	DeleteAll()
}
