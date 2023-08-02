package service

import (
	"go.mongodb.org/mongo-driver/mongo"
	"user_service/domain/model"
	"user_service/domain/repository"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type UserService struct {
	store repository.UserStore
}

func NewUserService(store repository.UserStore) *UserService {
	return &UserService{
		store: store,
	}
}

func (service *UserService) Get(id primitive.ObjectID) (*model.User, error) {
	return service.store.Get(id)
}

func (service *UserService) GetAll() ([]*model.User, error) {
	return service.store.GetAll()
}

func (service *UserService) Insert(user *model.User) (*model.User, error) {
	return service.store.Insert(user)
}

func (service *UserService) Update(user *model.User) (*mongo.UpdateResult, error) {
	return service.store.Update(user)
}

func (service *UserService) Delete(id string) (*mongo.DeleteResult, error) {
	return service.store.Delete(id)
}
