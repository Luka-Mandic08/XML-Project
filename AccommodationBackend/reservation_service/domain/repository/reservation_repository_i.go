package repository

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"reservation_service/domain/model"
	"time"
)

type ReservationStore interface {
	Get(id primitive.ObjectID) (*model.Reservation, error)
	GetAll() ([]*model.Reservation, error)
	Insert(reservation *model.Reservation) (*model.Reservation, error)
	Delete(id primitive.ObjectID) (*mongo.DeleteResult, error)
	Update(reservation *model.Reservation) (*mongo.UpdateResult, error)
	GetAllPastByUserId(id string) ([]*model.Reservation, error)
	GetAllFutureByUserId(id string) ([]*model.Reservation, error)
	GetActiveByUserId(id string) ([]*model.Reservation, error)
	GetActiveForAccommodations(ids []string) ([]*model.Reservation, error)
	GetPastByUserId(guestId, accommodationId string) ([]*model.Reservation, error)
	GetPastForAccommodations(guestId string, ids []string) ([]*model.Reservation, error)
	GetReservationsForAccommodationsByStatus(accommodationIds []string, status string) ([]*model.Reservation, error)
	GetAllOverlapping(id string, statuses []string, from, to time.Time) ([]*model.Reservation, error)
	GetAllPastByAccommodationId(id string) ([]*model.Reservation, error)
	GetAllFutureByAccommodationId(id string) ([]*model.Reservation, error)
}
