package repository

import (
	"auth_service/domain/model"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type AuthStore interface {
	GetById(id primitive.ObjectID) (*model.Account, error)
	Insert(user *model.Account) (*model.Account, error)
	GetByUsername(username string) (*model.Account, error)
	GetByUserId(userId string) (*model.Account, error)
	Update(*model.Account) (*mongo.UpdateResult, *model.Account, error)
	Delete(id string) (*mongo.DeleteResult, error)
	GenerateAPIKey(userId string, apiKey string) (*mongo.UpdateResult, error)
	LinkAPIKey(userId string) (string, error)
}
