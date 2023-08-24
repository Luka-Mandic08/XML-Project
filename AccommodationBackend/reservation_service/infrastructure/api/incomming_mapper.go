package api

import (
	reservation "common/proto/reservation_service"
	"reservation_service/domain/model"
)

// INCOMMING MAPPING
func MapCreateRequestToReservation(request *reservation.CreateRequest) (*model.Reservation, error) {
	result := model.Reservation{
		AccommodationId: request.AccommodationId,
		Start:           request.Start.AsTime(),
		End:             request.End.AsTime(),
		UserId:          request.UserId,
		NumberOfGuests:  request.NumberOfGuests,
	}
	return &result, nil
}

func MapUpdateRequestToReservation(request *reservation.UpdateRequest) (*model.Reservation, error) {
	result := model.Reservation{
		AccommodationId: request.AccommodationId,
		Start:           request.Start.AsTime(),
		End:             request.End.AsTime(),
		UserId:          request.UserId,
		NumberOfGuests:  request.NumberOfGuests,
		Status:          request.Status,
		Price:           request.Price,
	}
	return &result, nil
}

func MapRequestRequestToReservation(request *reservation.RequestRequest) (*model.Reservation, error) {
	result := model.Reservation{
		AccommodationId: request.AccommodationId,
		Start:           request.Start.AsTime(),
		End:             request.End.AsTime(),
		UserId:          request.UserId,
		NumberOfGuests:  request.NumberOfGuests,
	}
	return &result, nil
}
