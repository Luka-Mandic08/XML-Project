package service

import (
	accommodation "common/proto/accommodation_service"
	notification "common/proto/notification_service"
	rating "common/proto/rating_service"
	reservation "common/proto/reservation_service"
	user "common/proto/user_service"
	"context"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"reservation_service/domain/model"
	"reservation_service/domain/repository"
	"time"
)

type ReservationService struct {
	store                repository.ReservationStore
	outstandingHostStore repository.OutstandingHostMongoDBStore
	orchestrator         *CreateReservationOrchestrator
	AccommodationClient  accommodation.AccommodationServiceClient
	ratingClient         rating.RatingServiceClient
	notificationClient   notification.NotificationServiceClient
	userClient           user.UserServiceClient
}

func NewReservationService(store repository.ReservationStore, outstandingHostStore repository.OutstandingHostMongoDBStore, orchestrator *CreateReservationOrchestrator, accommodationClient accommodation.AccommodationServiceClient, ratingClient rating.RatingServiceClient, notificationClient notification.NotificationServiceClient, userClient user.UserServiceClient) *ReservationService {
	return &ReservationService{
		store:                store,
		orchestrator:         orchestrator,
		outstandingHostStore: outstandingHostStore,
		AccommodationClient:  accommodationClient,
		ratingClient:         ratingClient,
		notificationClient:   notificationClient,
		userClient:           userClient,
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

	reservation.Status = "Canceled"
	reservation.Price = price
	_, err = service.store.Update(reservation)
	if err != nil {
		return nil, err
	}

	return reservation, nil
}

func (service *ReservationService) GetAllByUserId(id string) ([]*model.Reservation, []*model.Reservation, error) {
	past, err := service.store.GetAllPastByUserId(id)
	if err != nil {
		return nil, nil, err
	}
	future, err := service.store.GetAllFutureByUserId(id)
	if err != nil {
		return nil, nil, err
	}
	return past, future, nil
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
	service.UpdateOutstandingHostStatus(reservation)
	service.CreateNotification(reservation, "ReservationCreated")
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

	service.CreateNotification(reservation, "ReservationPending")
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
	service.CreateNotification(reservation, "ReservationApproved")
	return reservation, nil
}

func (service *ReservationService) GetAllIntercepting(reservation *model.Reservation) ([]*model.Reservation, error) {
	return service.store.GetAllOverlapping(reservation.AccommodationId, []string{"Pending"}, reservation.Start, reservation.End)
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
	service.CreateNotification(reservation, "ReservationDenied")
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
	service.CreateNotification(reservation, "ReservationCanceled")
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
			_, err = service.notificationClient.InsertNotification(context.TODO(), &notification.CreateNotification{
				NotificationText: "You have lost the status of outstanding host.",
				UserId:           hostId,
				Type:             "OutstandingHostStatus",
			})
			if err != nil {
				println("Unsuccessful notification creation: ", err.Error())
			}
		}
	}
	if status {
		response, err := service.outstandingHostStore.Insert(&model.OutstandingHost{Id: id})
		if err != nil {
			return err
		}
		if response {
			_, err = service.notificationClient.InsertNotification(context.TODO(), &notification.CreateNotification{
				NotificationText: "You have gained the status of outstanding host.",
				UserId:           hostId,
				Type:             "OutstandingHostStatus",
			})
			if err != nil {
				println("Unsuccessful notification creation: ", err.Error())
			}
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
	return service.store.GetAllOverlapping(request.GetAccommodationId(), []string{"Approved", "Pending"}, request.GetFrom().AsTime(), request.GetTo().AsTime())
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

func (service *ReservationService) GetAllByAccommodationId(id string) ([]*model.Reservation, []*model.Reservation, error) {
	past, err := service.store.GetAllPastByAccommodationId(id)
	if err != nil {
		return nil, nil, err
	}
	future, err := service.store.GetAllFutureByAccommodationId(id)
	if err != nil {
		return nil, nil, err
	}
	return past, future, nil
}

func (service *ReservationService) GetNumberOfCancelledReservationsForUser(id string) int32 {
	cancelledReservations, _ := service.store.GetAllCanceledByUserId(id)
	return int32(len(cancelledReservations))
}

func (service *ReservationService) IsUserAvailable(userId string, start time.Time, end time.Time) (bool, error) {
	reservationNum, err := service.store.IsUserAvailable(userId, start, end)
	if err != nil {
		return false, err
	}
	if reservationNum > 1 {
		return false, err
	}
	return true, nil
}

func (service *ReservationService) CreateNotification(reservation *model.Reservation, notificationType string) {
	println("Notification type: ", notificationType)
	accommodationInfo, _ := service.AccommodationClient.GetById(context.TODO(), &accommodation.GetByIdRequest{Id: reservation.AccommodationId})
	println(accommodationInfo.GetAccommodation())
	if notificationType == "ReservationCanceled" {
		guest, _ := service.userClient.Get(context.TODO(), &user.GetRequest{Id: reservation.UserId})
		println(guest.GetName())
		_, err := service.notificationClient.InsertNotification(context.TODO(), &notification.CreateNotification{
			NotificationText: "Guest " + guest.GetName() + " " + guest.GetSurname() + " has canceled a reservation for " + string(reservation.NumberOfGuests) + " people at: " + accommodationInfo.GetAccommodation().GetName(),
			UserId:           accommodationInfo.Accommodation.GetHostId(),
			Type:             notificationType,
		})
		if err != nil {
			println("Imam neki error placem: ", err.Error())
			return
		}
	}
	if notificationType == "ReservationDenied" {
		service.notificationClient.InsertNotification(context.TODO(), &notification.CreateNotification{
			NotificationText: "Your reservation at " + accommodationInfo.GetAccommodation().GetName() + " has been denied.",
			UserId:           reservation.UserId,
			Type:             "ReservationApprovedOrDenied",
		})
	}
	if notificationType == "ReservationApproved" {
		service.notificationClient.InsertNotification(context.TODO(), &notification.CreateNotification{
			NotificationText: "Your reservation at " + accommodationInfo.GetAccommodation().GetName() + " has been approved.",
			UserId:           reservation.UserId,
			Type:             "ReservationApprovedOrDenied",
		})
	}
	if notificationType == "ReservationPending" {
		guest, _ := service.userClient.Get(context.TODO(), &user.GetRequest{Id: reservation.UserId})
		service.notificationClient.InsertNotification(context.TODO(), &notification.CreateNotification{
			NotificationText: "Guest " + guest.GetName() + " " + guest.GetSurname() + " has created a pending reservation for " + string(reservation.NumberOfGuests) + " people at: " + accommodationInfo.GetAccommodation().GetName(),
			UserId:           accommodationInfo.Accommodation.GetHostId(),
			Type:             "ReservationCreated",
		})
	}
	if notificationType == "ReservationCreated" {
		guest, _ := service.userClient.Get(context.TODO(), &user.GetRequest{Id: reservation.UserId})
		service.notificationClient.InsertNotification(context.TODO(), &notification.CreateNotification{
			NotificationText: "Guest " + guest.GetName() + " " + guest.GetSurname() + " has created an auto approved reservation for " + string(reservation.NumberOfGuests) + " people at: " + accommodationInfo.GetAccommodation().GetName(),
			UserId:           accommodationInfo.Accommodation.GetHostId(),
			Type:             "ReservationCreated",
		})
	}
}
