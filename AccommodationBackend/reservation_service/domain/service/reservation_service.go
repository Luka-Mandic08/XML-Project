package service

import (
	"go.mongodb.org/mongo-driver/mongo"
	"reservation_service/domain/model"
	"reservation_service/domain/repository"

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

func (service *ReservationService) Cancel(id primitive.ObjectID) (*model.Reservation, error) {
	reservation, err := service.store.Get(id)
	if err != nil {
		return nil, err
	}

	reservation.Status = "Cancelled"
	_, err = service.store.Update(reservation)
	if err != nil {
		return nil, err
	}

	return reservation, nil
}

func (service *ReservationService) GetAllByUserId(id primitive.ObjectID) ([]*model.Reservation, error) {
	return service.store.GetAllByUserId(id)
}

func (service *ReservationService) Approve(id primitive.ObjectID, price float32) (*model.Reservation, error) {
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
