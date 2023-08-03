package repository

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"user_service/domain/model"
)

type UserStore interface {
	Get(id primitive.ObjectID) (*model.User, error)
	GetAll() ([]*model.User, error)
	Insert(user *model.User) (*model.User, error)
	Delete(id primitive.ObjectID) (*mongo.DeleteResult, error)
	Update(user *model.User) (*mongo.UpdateResult, error)
}
