package repository

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"reservation_service/domain/model"
)

type ReservationStore interface {
	Get(id primitive.ObjectID) (*model.Reservation, error)
	GetAll() ([]*model.Reservation, error)
	Insert(reservation *model.Reservation) (*model.Reservation, error)
	Delete(id primitive.ObjectID) (*mongo.DeleteResult, error)
	Update(reservation *model.Reservation) (*mongo.UpdateResult, error)
}
