package service

import (
	"auth_service/domain/model"
	"auth_service/domain/repository"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type AuthService struct {
	store repository.AuthStore
}

func NewAuthService(store repository.AuthStore) *AuthService {
	return &AuthService{
		store: store,
	}
}

func (service *AuthService) Get(id primitive.ObjectID) (*model.Account, error) {
	return service.store.GetById(id)
}

func (service *AuthService) GetByUsername(username string) (*model.Account, error) {
	return service.store.GetByUsername(username)
}

func (service *AuthService) Insert(account *model.Account) (*model.Account, error) {
	return service.Insert(account)
}
