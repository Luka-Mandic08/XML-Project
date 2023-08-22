package service

import (
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"reservation_service/domain/model"
	"reservation_service/domain/repository"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type ReservationService struct {
	store        repository.ReservationStore
	orchestrator *CreateReservationOrchestrator
}

func NewReservationService(store repository.ReservationStore, orchestrator *CreateReservationOrchestrator) *ReservationService {
	return &ReservationService{
		store:        store,
		orchestrator: orchestrator,
	}
}

func (service *ReservationService) Get(id primitive.ObjectID) (*model.Reservation, error) {
	return service.store.Get(id)
}

func (service *ReservationService) GetAll() ([]*model.Reservation, error) {
	return service.store.GetAll()
}

func (service *ReservationService) Insert(reservation *model.Reservation) (*model.Reservation, error) {
	reservation.Status = "Pending"
	_, err := service.store.Insert(reservation)
	if err != nil {
		return nil, err
	}

	//OVDE SE POZIVA SAGA
	err = service.orchestrator.Start(reservation)

	return reservation, nil
}

func (service *ReservationService) Update(reservation *model.Reservation) (*mongo.UpdateResult, error) {
	return service.store.Update(reservation)
}

func (service *ReservationService) Delete(id string) (*mongo.DeleteResult, error) {
	uuid, _ := primitive.ObjectIDFromHex(id)
	return service.store.Delete(uuid)
}

func (service *ReservationService) AutoCancel(id primitive.ObjectID, price float32) (*model.Reservation, error) {
	reservation, err := service.store.Get(id)
	if err != nil {
		return nil, err
	}

	reservation.Status = "Cancelled"
	reservation.Price = price
	_, err = service.store.Update(reservation)
	if err != nil {
		return nil, err
	}

	return reservation, nil
}

func (service *ReservationService) GetAllByUserId(id primitive.ObjectID) ([]*model.Reservation, error) {
	return service.store.GetAllByUserId(id)
}

func (service *ReservationService) AutoApprove(id primitive.ObjectID, price float32) (*model.Reservation, error) {
	reservation, err := service.store.Get(id)
	if err != nil {
		return nil, err
	}

	reservation.Status = "Approved"
	reservation.Price = price
	_, err = service.store.Update(reservation)
	if err != nil {
		return nil, err
	}

	return reservation, nil
}

func (service *ReservationService) GetActiveByUserId(id string) (bool, error) {
	reservations, err := service.store.GetActiveByUserId(id)
	if err != nil {
		return true, err
	}
	if len(reservations) == 0 {
		return false, nil
	}
	return true, nil
}

func (service *ReservationService) GetActiveForAccommodations(ids []string) (bool, error) {
	reservations, err := service.store.GetActiveForAccommodations(ids)
	if err != nil {
		return true, err
	}
	if len(reservations) == 0 {
		return false, nil
	}
	return true, nil
}

func (service *ReservationService) GetPastByUserId(guestId, accommodationId string) (bool, error) {
	reservations, err := service.store.GetPastByUserId(guestId, accommodationId)
	if err != nil {
		return true, err
	}
	if len(reservations) == 0 {
		return false, nil
	}
	return true, nil
}

func (service *ReservationService) GetPastForAccommodations(guestId string, ids []string) (bool, error) {
	reservations, err := service.store.GetPastForAccommodations(guestId, ids)
	if err != nil {
		return true, err
	}
	if len(reservations) == 0 {
		return false, nil
	}
	return true, nil
}

func (service *ReservationService) AutoPending(id primitive.ObjectID, price float32) (*model.Reservation, error) {
	reservation, err := service.store.Get(id)
	if err != nil {
		return nil, err
	}

	reservation.Status = "Pending"
	reservation.Price = price
	_, err = service.store.Update(reservation)
	if err != nil {
		return nil, err
	}

	return reservation, nil
}

func (service *ReservationService) Approve(id primitive.ObjectID) (*model.Reservation, error) {
	reservation, err := service.store.Get(id)
	if err != nil {
		return nil, err
	}

	reservation.Status = "Approved"
	_, err = service.store.Update(reservation)
	if err != nil {
		return nil, err
	}

	return reservation, nil
}

func (service *ReservationService) GetAllIntercepting(reservation *model.Reservation) ([]*model.Reservation, error) {
	reservations, err := service.store.GetAllIntercepting(reservation)
	if err != nil {
		return nil, err
	}

	result := []*model.Reservation{}

	layout := "2006-01-02T15:04:05"
	reservationFrom, err := time.Parse(layout, reservation.Start)
	if err != nil {
		fmt.Println("Error parsing time:", err)
		return nil, err
	}
	reservationTo, err := time.Parse(layout, reservation.End)
	if err != nil {
		fmt.Println("Error parsing time:", err)
		return nil, err
	}
	for _, currentReservation := range reservations {
		dateFrom, err := time.Parse(layout, currentReservation.Start)
		if err != nil {
			fmt.Println("Error parsing time:", err)
			return nil, err
		}
		dateTo, err := time.Parse(layout, currentReservation.End)
		if err != nil {
			fmt.Println("Error parsing time:", err)
			return nil, err
		}
		if (dateFrom.Before(reservationTo) && dateFrom.After(reservationFrom)) || (dateTo.Before(reservationTo) && dateTo.After(reservationFrom)) {
			result = append(result, currentReservation)
		} else if dateFrom.Equal(reservationFrom) || dateTo.Equal(reservationTo) {
			result = append(result, currentReservation)
		} else if dateFrom.Before(reservationFrom) && dateTo.After(reservationTo) {
			result = append(result, currentReservation)
		}
	}

	return result, nil
}

func (service *ReservationService) Deny(id primitive.ObjectID) (*model.Reservation, error) {
	reservation, err := service.store.Get(id)
	if err != nil {
		return nil, err
	}

	reservation.Status = "Denied"
	_, err = service.store.Update(reservation)
	if err != nil {
		return nil, err
	}

	return reservation, nil
}

func (service *ReservationService) Cancel(id primitive.ObjectID) (*model.Reservation, error) {
	reservation, err := service.store.Get(id)
	if err != nil {
		return nil, err
	}

	reservation.Status = "Canceled"
	_, err = service.store.Update(reservation)
	if err != nil {
		return nil, err
	}

	return reservation, nil
}
