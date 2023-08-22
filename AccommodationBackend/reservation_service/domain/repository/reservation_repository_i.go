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
	GetAllByUserId(id primitive.ObjectID) ([]*model.Reservation, error)
	GetActiveByUserId(id string) ([]*model.Reservation, error)
	GetActiveForAccommodations(ids []string) ([]*model.Reservation, error)
	GetPastByUserId(guestId, accommodationId string) ([]*model.Reservation, error)
	GetPastForAccommodations(guestId string, ids []string) ([]*model.Reservation, error)
	GetAllIntercepting(reservation *model.Reservation) ([]*model.Reservation, error)
}
