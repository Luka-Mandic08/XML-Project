package service

import (
	accommodation "common/proto/accommodation_service"
	rating "common/proto/rating_service"
	reservation "common/proto/reservation_service"
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"reservation_service/domain/model"
	"reservation_service/domain/repository"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type ReservationService struct {
	store                repository.ReservationStore
	outstandingHostStore repository.OutstandingHostMongoDBStore
	orchestrator         *CreateReservationOrchestrator
	AccommodationClient  accommodation.AccommodationServiceClient
	ratingClient         rating.RatingServiceClient
}

func NewReservationService(store repository.ReservationStore, outstandingHostStore repository.OutstandingHostMongoDBStore, orchestrator *CreateReservationOrchestrator, accommodationClient accommodation.AccommodationServiceClient, ratingClient rating.RatingServiceClient) *ReservationService {
	return &ReservationService{
		store:                store,
		orchestrator:         orchestrator,
		outstandingHostStore: outstandingHostStore,
		AccommodationClient:  accommodationClient,
		ratingClient:         ratingClient,
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
	//TODO Add CheckOutstandingHost Saga :(

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

func (service *ReservationService) GetAllOverlapping(request reservation.GetAllForDateRangeRequest) ([]*model.Reservation, error) {
	reservations, err := service.store.GetAllOverlapping(request.GetAccommodationId())
	if err != nil {
		return nil, err
	}
	result := []*model.Reservation{}
	layout := "2006-01-02T15:04:05"
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
		if (dateFrom.Before(request.To.AsTime()) && dateFrom.After(request.From.AsTime())) || (dateTo.Before(request.To.AsTime()) && dateTo.After(request.From.AsTime())) {
			result = append(result, currentReservation)
		} else if dateFrom.Equal(request.From.AsTime()) || dateTo.Equal(request.To.AsTime()) {
			result = append(result, currentReservation)
		} else if dateFrom.Before(request.From.AsTime()) && dateTo.After(request.To.AsTime()) {
			result = append(result, currentReservation)
		}
	}
	return result, nil
}

func (service *ReservationService) UpdateOutstandingHostStatus(reservation *model.Reservation) {
	accResponse, err := service.AccommodationClient.GetAllForHostByAccommodationId(context.TODO(), &accommodation.GetByIdRequest{Id: reservation.AccommodationId})
	if err != nil {
		return
	}
	ratingResponse, err := service.ratingClient.GetAverageScoreForHost(context.TODO(), &rating.IdRequest{Id: accResponse.GetHostId()})
	if err != nil {
		return
	}
	if ratingResponse.GetScore() <= 4.7 {
		service.ChangeOutstandingHostStatus(false, accResponse.HostId)
		return
	}
	shouldBeOutstanding, _ := service.CheckOutstandingHostStatus(accResponse.AccommodationIds)
	service.ChangeOutstandingHostStatus(shouldBeOutstanding, accResponse.HostId)
	return
}
