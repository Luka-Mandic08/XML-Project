package api

import (
	reservation "common/proto/reservation_service"
	"fmt"
	"reservation_service/domain/model"
	"time"
)

// INCOMMING MAPPING
func MapCreateRequestToReservation(request *reservation.CreateRequest) (*model.Reservation, error) {
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
		AccommodationId: request.AccommodationId,
		Start:           startTime.Format("2006-01-02T15:04:05.000000000"),
		End:             endTime.Format("2006-01-02T15:04:05.000000000"),
		UserId:          request.UserId,
		NumberOfGuests:  request.NumberOfGuests,
	}
	return &result, nil
}

func MapUpdateRequestToReservation(request *reservation.UpdateRequest) (*model.Reservation, error) {
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
		AccommodationId: request.AccommodationId,
		Start:           startTime.Format("2006-01-02T15:04:05.000000000"),
		End:             endTime.Format("2006-01-02T15:04:05.000000000"),
		UserId:          request.UserId,
		NumberOfGuests:  request.NumberOfGuests,
		Status:          request.Status,
		Price:           request.Price,
	}
	return &result, nil
}

func MapRequestRequestToReservation(request *reservation.RequestRequest) (*model.Reservation, error) {
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
		AccommodationId: request.AccommodationId,
		Start:           startTime.Format("2006-01-02T15:04:05.000000000"),
		End:             endTime.Format("2006-01-02T15:04:05.000000000"),
		UserId:          request.UserId,
		NumberOfGuests:  request.NumberOfGuests,
	}
	return &result, nil
}
