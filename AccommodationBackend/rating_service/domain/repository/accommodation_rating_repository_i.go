package repository

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"rating_service/domain/model"
)

type AccommodationRatingStore interface {
	Get(id primitive.ObjectID) (*model.AccommodationRating, error)
	GetAllForAccommodation(accommodationId string) ([]*model.AccommodationRating, error)
	GetAverageScoreForAccommodation(accommodationId string) (float32, error)
	Insert(accommodationRating *model.AccommodationRating) (*model.AccommodationRating, error)
	Delete(id primitive.ObjectID) (*mongo.DeleteResult, error)
	Update(accommodationRating *model.AccommodationRating) (*mongo.UpdateResult, error)
}
