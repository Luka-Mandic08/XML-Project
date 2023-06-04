package application

import "github.com/tamararankovic/microservices_demo/inventory_service/domain"

type ProductService struct {
	store domain.ProductStore
}

func NewProductService(store domain.ProductStore) *ProductService {
	return &ProductService{
		store: store,
	}
}

func (service *ProductService) UpdateQuantity(product *domain.Product, difference int64) error {
	return service.store.UpdateQuantity(product, difference)
}

func (service *ProductService) UpdateQuantityForAll(products map[*domain.Product]int64) error {
	return service.store.UpdateQuantityForAll(products)
}

func (service *ProductService) GetAll() (*[]domain.Product, error) {
	return service.store.GetAll()
}

func (service *ProductService) DeleteAll() {
	service.store.DeleteAll()
}
