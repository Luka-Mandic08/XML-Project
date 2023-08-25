package service

import (
	"auth_service/domain/model"
	"auth_service/domain/repository"
	"crypto/rand"
	"encoding/base64"
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

func (service *AuthService) GenerateAPIKey(userId string) (*mongo.UpdateResult, error) {
	keyLength := 32

	// Generate random bytes
	randomBytes := make([]byte, keyLength)
	_, err := rand.Read(randomBytes)
	if err != nil {
		return nil, err
	}

	// Encode random bytes to base64
	apiKey := base64.URLEncoding.EncodeToString(randomBytes)
	println(apiKey)

	return service.store.GenerateAPIKey(userId, apiKey)
}

func (service *AuthService) LinkAPIKey(userId string) (string, error) {
	return service.store.LinkAPIKey(userId)
}
