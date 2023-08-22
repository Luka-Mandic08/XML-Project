package api

import (
	reservation "common/proto/reservation_service"
	"reservation_service/domain/model"
)

// OUTGOING MAPPING
func MapReservationToGetResponse(u *model.Reservation) *reservation.GetResponse {
	result := reservation.GetResponse{
		Id:              u.Id.Hex(),
		AccommodationId: u.AccommodationId,
		Start:           u.Start,
		End:             u.End,
		UserId:          u.UserId,
		NumberOfGuests:  u.NumberOfGuests,
		Status:          u.Status,
		Price:           u.Price,
	}
	return &result
}

func MapReservationToCreateResponse(u *model.Reservation) *reservation.CreateResponse {
	result := reservation.CreateResponse{
		Id: u.Id.Hex(),
	}
	return &result
}

func MapReservationToUpdateResponse(u *model.Reservation) *reservation.UpdateResponse {
	result := reservation.UpdateResponse{
		Id:              u.Id.Hex(),
		AccommodationId: u.AccommodationId,
		Start:           u.Start,
		End:             u.End,
		UserId:          u.UserId,
		NumberOfGuests:  u.NumberOfGuests,
		Status:          u.Status,
		Price:           u.Price,
	}

	return &result
}

func MapReservationToDeleteResponse() *reservation.DeleteResponse {
	result := reservation.DeleteResponse{
		Message: "Reservation successfully deleted!",
	}
	return &result
}

func MapReservationToRequestResponse(u *model.Reservation) *reservation.RequestResponse {
	result := reservation.RequestResponse{
		Id: u.Id.Hex(),
	}
	return &result
}

func MapReservationsToGetAllByUserIdResponse(u []*model.Reservation) *reservation.GetAllByUserIdResponse {
	reservations := []*reservation.Reservation{}

	for _, currentReservation := range u {
		reservation := reservation.Reservation{
			Id:              currentReservation.Id.Hex(),
			AccommodationId: currentReservation.AccommodationId,
			Start:           currentReservation.Start,
			End:             currentReservation.End,
			UserId:          currentReservation.UserId,
			NumberOfGuests:  currentReservation.NumberOfGuests,
			Status:          currentReservation.Status,
			Price:           currentReservation.Price,
		}
		reservations = append(reservations, &reservation)

	}
	result := reservation.GetAllByUserIdResponse{
		Reservation: reservations,
	}
	return &result
}

func MapReservationToApproveResponse(u *model.Reservation) *reservation.ApproveResponse {
	result := reservation.ApproveResponse{
		Id: u.Id.Hex(),
	}
	return &result
}

func MapReservationToDenyResponse(u *model.Reservation) *reservation.DenyResponse {
	result := reservation.DenyResponse{
		Id: u.Id.Hex(),
	}
	return &result
}

func MapReservationToCancelResponse(u *model.Reservation) *reservation.CancelResponse {
	result := reservation.CancelResponse{
		Id: u.Id.Hex(),
	}
	return &result
}
