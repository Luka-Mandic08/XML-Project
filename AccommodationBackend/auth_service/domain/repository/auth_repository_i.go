package repository

import (
	"auth_service/domain/model"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type AuthStore interface {
	GetById(id primitive.ObjectID) (*model.Account, error)
	Insert(user *model.Account) error
	GetByUsername(username string) (*model.Account, error)
}
