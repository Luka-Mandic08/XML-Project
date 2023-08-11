package api

import (
	reservation "common/proto/reservation_service"
	"fmt"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"reservation_service/domain/model"
	"time"
)

func MapReservationToGetResponse(u *model.Reservation) *reservation.GetResponse {
	result := reservation.GetResponse{
		Id:            u.Id.Hex(),
		Accommodation: u.AccommodationId,
		Start:         u.Start.String(),
		End:           u.End.String(),
		User:          u.UserId,
	}
	return &result
}

func MapCreateRequestToReservation(request *reservation.CreateRequest) (*model.Reservation, error) {
	id, _ := primitive.ObjectIDFromHex(request.Id)
	startDateString := request.Start
	endDateString := request.End

	layout := "2006-01-02T15:04:05"

	startTime, err := time.Parse(layout, startDateString)
	if err != nil {
		fmt.Println("Error parsing time:", err)
		return nil, err
	}
	endTime, err := time.Parse(layout, endDateString)
	if err != nil {
		fmt.Println("Error parsing time:", err)
		return nil, err
	}

	result := model.Reservation{
		Id:              id,
		AccommodationId: request.Accommodation,
		Start:           startTime,
		End:             endTime,
		UserId:          request.User,
	}
	return &result, nil
}

func MapUpdateRequestToReservation(request *reservation.UpdateRequest) (*model.Reservation, error) {
	id, _ := primitive.ObjectIDFromHex(request.Id)
	startDateString := request.Start
	endDateString := request.End

	layout := "2006-01-02T15:04:05"

	startTime, err := time.Parse(layout, startDateString)
	if err != nil {
		fmt.Println("Error parsing time:", err)
		return nil, err
	}
	endTime, err := time.Parse(layout, endDateString)
	if err != nil {
		fmt.Println("Error parsing time:", err)
		return nil, err
	}

	result := model.Reservation{
		Id:              id,
		AccommodationId: request.Accommodation,
		Start:           startTime,
		End:             endTime,
		UserId:          request.User,
	}
	return &result, nil
}
