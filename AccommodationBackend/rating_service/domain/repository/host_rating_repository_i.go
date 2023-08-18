package repository

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"rating_service/domain/model"
)

type HostRatingStore interface {
	Get(id primitive.ObjectID) (*model.HostRating, error)
	GetAllForHost(hostId string) ([]*model.HostRating, error)
	GetAverageScoreForHost(hostId string) (float32, error)
	Insert(hostRating *model.HostRating) (*model.HostRating, error)
	Delete(id primitive.ObjectID) (*mongo.DeleteResult, error)
	Update(hostRating *model.HostRating) (*mongo.UpdateResult, error)
}
