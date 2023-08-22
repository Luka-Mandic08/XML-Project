package service

import (
	"go.mongodb.org/mongo-driver/mongo"
	"reservation_service/domain/model"
	"reservation_service/domain/repository"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type ReservationService struct {
	store                repository.ReservationStore
	outstandingHostStore repository.OutstandingHostMongoDBStore
	orchestrator         *CreateReservationOrchestrator
}

func NewReservationService(store repository.ReservationStore, outstandingHostStore repository.OutstandingHostMongoDBStore, orchestrator *CreateReservationOrchestrator) *ReservationService {
	return &ReservationService{
		store:                store,
		orchestrator:         orchestrator,
		outstandingHostStore: outstandingHostStore,
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

func (service *ReservationService) CheckOutstandingHostStatus(accommodationIds []string) (bool, error) {
	approvedReservation, err := service.store.GetReservationsForAccommodationsByStatus(accommodationIds, "Approved")
	if err != nil {
		return false, err
	}
	if len(approvedReservation) < 5 {
		return false, nil
	}

	var totalDuration int32
	for _, r := range approvedReservation {
		totalDuration += r.CalculateDuration()
	}
	if totalDuration < 50 {
		return false, nil
	}

	canceledReservation, err := service.store.GetReservationsForAccommodationsByStatus(accommodationIds, "Canceled")
	if err != nil {
		return false, err
	}
	if float32(len(canceledReservation))/float32(len(approvedReservation)) >= 0.05 {
		return false, nil
	}

	return true, nil
}

func (service *ReservationService) ChangeOutstandingHostStatus(status bool, hostId string) error {
	id, err := primitive.ObjectIDFromHex(hostId)
	if err != nil {
		return err
	}
	if !status {
		response, _ := service.outstandingHostStore.Delete(id)
		if response.DeletedCount == 1 {
			//TODO Send notification to host
		}
	}
	if status {
		response, err := service.outstandingHostStore.Insert(&model.OutstandingHost{Id: id})
		if err != nil {
			return err
		}
		if response {
			//TODO Send notification to host
		}
	}
	return nil
}

func (service *ReservationService) GetOutstandingHost(hostId string) (*model.OutstandingHost, error) {
	id, _ := primitive.ObjectIDFromHex(hostId)
	return service.outstandingHostStore.Get(id)
}

func (service *ReservationService) GetAllOutstandingHosts() ([]*model.OutstandingHost, error) {
	return service.outstandingHostStore.GetAll()
}
