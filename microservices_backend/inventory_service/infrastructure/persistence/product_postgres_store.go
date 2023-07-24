package persistence

import (
	"errors"
	"github.com/tamararankovic/microservices_demo/inventory_service/domain"
	"gorm.io/gorm"
)

type ProductPostgresStore struct {
	db *gorm.DB
}

func NewProductPostgresStore(db *gorm.DB) (domain.ProductStore, error) {
	err := db.AutoMigrate(&domain.Product{})
	if err != nil {
		return nil, err
	}
	return &ProductPostgresStore{
		db: db,
	}, nil
}

func (store *ProductPostgresStore) Insert(product *domain.Product) error {
	result := store.db.Create(product)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (store *ProductPostgresStore) UpdateQuantity(product *domain.Product, difference int64) error {
	return updateQuantity(store.db, product, difference)
}

func (store *ProductPostgresStore) UpdateQuantityForAll(products map[*domain.Product]int64) error {
	return store.db.Transaction(func(tx *gorm.DB) error {
		for product, difference := range products {
			err := updateQuantity(tx, product, difference)
			if err != nil {
				return err
			}
		}
		return nil
	})
}

func updateQuantity(tx *gorm.DB, product *domain.Product, difference int64) error {
	tx = tx.Model(&domain.Product{}).
		Where("product_id = ? AND color_code = ? AND quantity + ? >= 0", product.ProductId, product.ColorCode, difference).
		Update("quantity", gorm.Expr("quantity + ?", difference))
	if tx.Error != nil {
		return tx.Error
	}
	if tx.RowsAffected != 1 {
		return errors.New("update error")
	}
	return nil
}

func (store *ProductPostgresStore) GetAll() (*[]domain.Product, error) {
	var products []domain.Product
	result := store.db.Find(&products)
	if result.Error != nil {
		return nil, result.Error
	}
	return &products, nil
}

func (store *ProductPostgresStore) DeleteAll() {
	store.db.Session(&gorm.Session{AllowGlobalUpdate: true}).
		Delete(&domain.Product{})
}
