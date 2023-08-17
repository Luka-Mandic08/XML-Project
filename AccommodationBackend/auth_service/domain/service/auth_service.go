package service

import (
	"auth_service/domain/model"
	"auth_service/domain/repository"
	"go.mongodb.org/mongo-driver/mongo"

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

func (service *AuthService) GetByUserId(userId string) (*model.Account, error) {
	return service.store.GetByUserId(userId)
}

func (service *AuthService) Insert(account *model.Account) (*model.Account, error) {
	return service.store.Insert(account)
}

func (service *AuthService) Update(account *model.Account) (*mongo.UpdateResult, *model.Account, error) {
	return service.store.Update(account)
}

func (service *AuthService) Delete(id string) (*mongo.DeleteResult, error) {
	return service.store.Delete(id)
}
