package api

import (
	accommodation "common/proto/accommodation_service"
	pb "common/proto/reservation_service"
	user "common/proto/user_service"
	"context"
	"errors"
	"github.com/golang/protobuf/ptypes"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"reservation_service/domain/service"
	"time"
)

type ReservationHandler struct {
	pb.UnimplementedReservationServiceServer
	reservationService *service.ReservationService
	userClient         user.UserServiceClient
}

func NewReservationHandler(reservationService *service.ReservationService, userClient user.UserServiceClient) *ReservationHandler {
	return &ReservationHandler{
		reservationService: reservationService,
		userClient:         userClient,
	}
}

func (handler *ReservationHandler) Get(ctx context.Context, request *pb.GetRequest) (*pb.GetResponse, error) {
	id := request.Id
	objectId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}
	reservation, err := handler.reservationService.Get(objectId)
	if err == mongo.ErrNoDocuments {
		return nil, status.Error(codes.NotFound, "Unable to find reservation: id = "+request.Id)
	}
	response := MapReservationToGetResponse(reservation)
	return response, nil
}

func (handler *ReservationHandler) Create(ctx context.Context, request *pb.CreateRequest) (*pb.CreateResponse, error) {
	reservation, err := MapCreateRequestToReservation(request)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, "Unable to convert String to DateTime")
	}
	reservation, err = handler.reservationService.Insert(reservation)
	if err != nil {
		return nil, status.Error(codes.AlreadyExists, "Unable to insert reservation into database")
	}

	//TODO OVDE SE POZIVA SAGA IVANEEEEEEEEEEEEEEEEEEEEEEEEEEEEEE

	response := MapReservationToCreateResponse(reservation)
	return response, nil
}

func (handler *ReservationHandler) Update(ctx context.Context, request *pb.UpdateRequest) (*pb.UpdateResponse, error) {
	reservation, err := MapUpdateRequestToReservation(request)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, "Unable to convert String to DateTime")
	}
	result, err := handler.reservationService.Update(reservation)
	if err != nil {
		return nil, status.Error(codes.Unknown, "Unable to update reservation: id = "+request.Id)
	}
	if result.MatchedCount == 0 {
		return nil, status.Error(codes.NotFound, "Unable to find reservation: id = "+request.Id)
	}
	response := MapReservationToUpdateResponse(reservation)
	return response, nil
}

func (handler *ReservationHandler) Delete(ctx context.Context, request *pb.DeleteRequest) (*pb.DeleteResponse, error) {
	result, err := handler.reservationService.Delete(request.Id)
	if err != nil {
		return nil, status.Error(codes.Unknown, "Unable to delete reservation: id = "+request.Id)
	}
	if result.DeletedCount == 0 {
		return nil, status.Error(codes.NotFound, "Unable to find reservation: id = "+request.Id)
	}
	response := MapReservationToDeleteResponse()
	return response, nil
}

func (handler *ReservationHandler) GetAllByUserId(ctx context.Context, request *pb.GetAllByUserIdRequest) (*pb.GetAllByAccommodationIdResponse, error) {
	past, future, err := handler.reservationService.GetAllByUserId(request.UserId)
	if err != nil {
		return nil, status.Error(codes.NotFound, "Unable to find reservations for user: id = "+request.UserId)
	}
	response := MapToGetAllByUserIdResponse(past, future)
	return response, nil
}

func (handler *ReservationHandler) Request(ctx context.Context, request *pb.RequestRequest) (*pb.RequestResponse, error) {
	reservationRequest, err := MapRequestRequestToReservation(request)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, "Unable to convert String to DateTime")
	}
	reservationRequest, err = handler.reservationService.Insert(reservationRequest)
	if err != nil {
		return nil, status.Error(codes.AlreadyExists, "Unable to insert request into database")
	}
	response := MapReservationToRequestResponse(reservationRequest)
	return response, nil
}

func (handler *ReservationHandler) CheckIfGuestHasReservations(ctx context.Context, request *pb.CheckReservationRequest) (*pb.CheckReservationResponse, error) {
	hasReservations, err := handler.reservationService.GetActiveByUserId(request.GetId())
	if err != nil {
		return nil, err
	}
	if hasReservations {
		return nil, status.Error(codes.Canceled, "User has active reservations")
	}
	return &pb.CheckReservationResponse{Message: "Success"}, nil
}

func (handler *ReservationHandler) CheckIfHostHasReservations(ctx context.Context, request *pb.CheckReservationRequest) (*pb.CheckReservationResponse, error) {
	accommodations, err := handler.reservationService.AccommodationClient.GetAllByHostId(ctx, &accommodation.GetAllByHostIdRequest{HostId: request.GetId()})
	var ids []string
	for _, a := range accommodations.GetAccommodations() {
		ids = append(ids, a.GetId())
	}
	if len(ids) == 0 {
		return &pb.CheckReservationResponse{Message: "Success"}, nil
	}
	hasReservations, err := handler.reservationService.GetActiveForAccommodations(ids)
	if err != nil {
		return nil, err
	}
	if hasReservations {
		return nil, status.Error(codes.Canceled, "User has active reservations")
	}
	handler.reservationService.AccommodationClient.DeleteAllForHost(ctx, &accommodation.GetByIdRequest{Id: request.GetId()})
	return &pb.CheckReservationResponse{Message: "Success"}, nil
}

func (handler *ReservationHandler) CheckIfGuestVisitedAccommodation(ctx context.Context, request *pb.CheckPreviousReservationRequest) (*pb.CheckReservationResponse, error) {
	hasReservations, err := handler.reservationService.GetPastByUserId(request.GetGuestId(), request.GetId())
	if err != nil {
		return nil, err
	}
	if !hasReservations {
		return nil, status.Error(codes.Canceled, "User has no previous reservations")
	}
	return &pb.CheckReservationResponse{Message: "Success"}, nil
}

func (handler *ReservationHandler) CheckIfGuestVisitedHost(ctx context.Context, request *pb.CheckPreviousReservationRequest) (*pb.CheckReservationResponse, error) {
	accommodations, err := handler.reservationService.AccommodationClient.GetAllByHostId(ctx, &accommodation.GetAllByHostIdRequest{HostId: request.GetId()})
	var ids []string
	for _, a := range accommodations.GetAccommodations() {
		ids = append(ids, a.GetId())
	}
	if len(ids) == 0 {
		return nil, status.Error(codes.Canceled, "User has no previous reservations 1")
	}
	hasReservations, err := handler.reservationService.GetPastForAccommodations(request.GetGuestId(), ids)
	if err != nil {
		return nil, err
	}
	if !hasReservations {
		return nil, status.Error(codes.Canceled, "User has no previous reservations 2")
	}
	return &pb.CheckReservationResponse{Message: "Success"}, nil
}

func (handler *ReservationHandler) Approve(ctx context.Context, request *pb.ApproveRequest) (*pb.ApproveResponse, error) {
	id := request.Id
	reservationId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}
	reservation, err := handler.reservationService.Get(reservationId)
	if err != nil {
		return nil, err
	}

	if reservation.Status == "Approved" {
		return nil, errors.New("Reservation alredy Approved id: " + id)
	}
	if reservation.Status == "Denied" {
		return nil, errors.New("Reservation alredy Denied id: " + id)
	}
	if reservation.Status == "Canceled" {
		return nil, errors.New("Reservation alredy Canceled id: " + id)
	}

	// Check if reservation is in past
	if reservation.Start.Before(time.Now()) {
		return nil, errors.New("This reservation is in past!")
	}

	start, _ := ptypes.TimestampProto(reservation.Start)
	end, _ := ptypes.TimestampProto(reservation.End)
	canApprove, err := handler.reservationService.AccommodationClient.CheckCanApprove(ctx, &accommodation.CheckCanApproveRequest{
		AccommodationId: reservation.AccommodationId,
		Start:           start,
		End:             end,
	})
	if err != nil {
		return nil, err
	}

	if canApprove.CanApprove != "true" {
		return &pb.ApproveResponse{Id: id}, errors.New("Cannot Approve Reservation")
	}

	reservation, err = handler.reservationService.Approve(reservationId)
	if err != nil {
		return nil, err
	}

	interceptingReservations, err := handler.reservationService.GetAllIntercepting(reservation)
	if err != nil {
		return nil, err
	}
	for _, reservation := range interceptingReservations {
		_, err = handler.reservationService.Deny(reservation.Id)
		if err != nil {
			return nil, err
		}
	}
	//TODO Add CheckOutstandingHost
	handler.reservationService.UpdateOutstandingHostStatus(reservation)
	response := MapReservationToApproveResponse(reservation)
	return response, nil
}

func (handler *ReservationHandler) Deny(ctx context.Context, request *pb.DenyRequest) (*pb.DenyResponse, error) {
	id := request.Id
	objectId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}
	reservation, err := handler.reservationService.Get(objectId)
	if err != nil {
		return nil, err
	}

	if reservation.Status == "Approved" {
		return nil, errors.New("Reservation alredy Approved id: " + id)
	}
	if reservation.Status == "Denied" {
		return nil, errors.New("Reservation alredy Denied id: " + id)
	}
	if reservation.Status == "Canceled" {
		return nil, errors.New("Reservation alredy Canceled id: " + id)
	}

	reservation, err = handler.reservationService.Deny(objectId)
	if err == mongo.ErrNoDocuments {
		return nil, status.Error(codes.NotFound, "Unable to find reservation: id = "+request.Id)
	}
	response := MapReservationToDenyResponse(reservation)
	return response, nil
}

func (handler *ReservationHandler) Cancel(ctx context.Context, request *pb.CancelRequest) (*pb.CancelResponse, error) {
	id := request.Id
	reservationId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}
	reservation, err := handler.reservationService.Get(reservationId)
	if err != nil {
		return nil, err
	}

	if reservation.Status == "Denied" {
		return nil, errors.New("Reservation alredy Denied id: " + id)
	}
	if reservation.Status == "Canceled" {
		return nil, errors.New("Reservation alredy Canceled id: " + id)
	}

	if reservation.Status == "Approved" {
		start, _ := ptypes.TimestampProto(reservation.Start)
		end, _ := ptypes.TimestampProto(reservation.End)
		_, err = handler.reservationService.AccommodationClient.GetAndCancelAllAvailabilitiesToCancel(ctx, &accommodation.GetAndCancelAllAvailabilitiesToCancelRequest{
			AccommodationId: reservation.AccommodationId,
			Start:           start,
			End:             end,
		})
		if err != nil {
			return nil, err
		}
	}
	_, err = handler.reservationService.Cancel(reservationId)
	if err == mongo.ErrNoDocuments {
		return nil, status.Error(codes.NotFound, "Unable to find reservation: id = "+request.Id)
	}
	response := MapReservationToCancelResponse(reservation)
	handler.reservationService.UpdateOutstandingHostStatus(reservation)
	return response, nil
}

func (handler *ReservationHandler) UpdateOutstandingHostStatus(ctx context.Context, request *pb.UpdateOutstandingHostStatusRequest) (*pb.UpdateOutstandingHostStatusResponse, error) {
	if !request.GetShouldUpdate() {
		err := handler.reservationService.ChangeOutstandingHostStatus(false, request.GetHostId())
		if err != nil {
			return nil, err
		}
		return &pb.UpdateOutstandingHostStatusResponse{Message: "Status changed"}, nil
	}
	response, err := handler.reservationService.AccommodationClient.GetAllByHostId(ctx, &accommodation.GetAllByHostIdRequest{HostId: request.GetHostId()})
	if err != nil {
		return nil, err
	}
	var ids []string
	for _, acc := range response.GetAccommodations() {
		ids = append(ids, acc.Id)
	}
	res, err := handler.reservationService.CheckOutstandingHostStatus(ids)
	if err != nil {
		return nil, err
	}
	if res {
		err = handler.reservationService.ChangeOutstandingHostStatus(true, request.GetHostId())
		if err != nil {
			return nil, err
		}
		return &pb.UpdateOutstandingHostStatusResponse{Message: "Status changed"}, nil
	}
	return &pb.UpdateOutstandingHostStatusResponse{Message: "Status unchanged"}, nil
}

func (handler *ReservationHandler) GetOutstandingHost(ctx context.Context, request *pb.GetRequest) (*pb.GetRequest, error) {
	response, err := handler.reservationService.GetOutstandingHost(request.GetId())
	if err != nil {
		return nil, err
	}
	return &pb.GetRequest{Id: response.Id.Hex()}, nil
}

func (handler *ReservationHandler) GetAllOutstandingHosts(ctx context.Context, request *pb.GetAllOutstandingHostsRequest) (*pb.GetAllOutstandingHostsResponse, error) {
	response, err := handler.reservationService.GetAllOutstandingHosts()
	if err != nil {
		return nil, err
	}
	var ids []string
	for _, id := range response {
		ids = append(ids, id.Id.Hex())
	}
	return &pb.GetAllOutstandingHostsResponse{Ids: ids}, nil
}

func (handler *ReservationHandler) GetAllForDateRange(ctx context.Context, request *pb.GetAllForDateRangeRequest) (*pb.GetAllForDateRangeResponse, error) {
	res, err := handler.reservationService.GetAllOverlapping(*request)
	if err != nil {
		return nil, err
	}
	mapped := MapReservationsToGetAllByUserIdResponse(res)
	return &pb.GetAllForDateRangeResponse{Reservations: mapped.GetReservation()}, nil
}

func (handler *ReservationHandler) GetAllByAccommodationId(ctx context.Context, request *pb.GetRequest) (*pb.GetAllByAccommodationIdResponse, error) {
	past, future, err := handler.reservationService.GetAllByAccommodationId(request.GetId())
	if err != nil {
		return nil, status.Error(codes.Aborted, "Error while reading reservations for accommodations")
	}
	var pastUsers []*user.GetForReservationResponse
	for _, res := range past {
		u, _ := handler.userClient.GetForReservation(ctx, &user.GetRequest{Id: res.UserId})
		if u != nil {
			pastUsers = append(pastUsers, u)
		}
		if u == nil {
			pastUsers = append(pastUsers, &user.GetForReservationResponse{
				Name:    "[Deleted]",
				Surname: "[Deleted]",
				Email:   "[Deleted]",
			})
		}
	}
	var futureUsers []*user.GetForReservationResponse
	var cancellations []int32
	for _, res := range future {
		user, err := handler.userClient.GetForReservation(ctx, &user.GetRequest{Id: res.UserId})
		if err != nil {
			return nil, status.Error(codes.Aborted, "Error while fetching guest from user service")
		}
		futureUsers = append(futureUsers, user)
		cancellations = append(cancellations, handler.reservationService.GetNumberOfCancelledReservationsForUser(res.UserId))
	}
	response := MapToGetAllByAccommodationIdResponse(past, future, pastUsers, futureUsers, cancellations)
	return response, nil
}
