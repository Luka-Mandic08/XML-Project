package repository

import (
	"accommodation_service/domain/model"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type AccommodationStore interface {
	GetById(id primitive.ObjectID) (*model.Accommodation, error)
	Insert(user *model.Accommodation) (*model.Accommodation, error)
	GetByAddress(username model.Address) (*model.Accommodation, error)
	Update(*model.Accommodation) (*mongo.UpdateResult, error)
	Delete(id string) (*mongo.DeleteResult, error)
	GetAllByHostId(hostId string) ([]*model.Accommodation, error)
	GetAll(page int) ([]*model.Accommodation, error)
	DeleteAllForHost(id string) (*mongo.DeleteResult, error)
	GetAllForHostByAccommodationId(id primitive.ObjectID) ([]string, error)
}
