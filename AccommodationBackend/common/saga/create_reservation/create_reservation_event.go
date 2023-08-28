package create_reservation

import "time"

type Reservation struct {
	Id              string
	AccommodationId string
	Start           time.Time
	End             time.Time
	UserId          string
	NumberOfGuests  int32
	Status          string
	Price           float32
}

type CreateReservationCommandType int8

const (
	CheckAccommodationExists CreateReservationCommandType = iota
	CheckAvailableAccommodation
	CheckUserExists
	CheckAutomaticApproveReservation
	RevertAvailability

	PendingReservation
	ApproveReservation
	CancelReservation
	DeleteReservation
	UnknownCommand
)

type CreateReservationCommand struct {
	Reservation Reservation
	Type        CreateReservationCommandType
}

type CreateReservationReplyType int8

const (
	AccommodationExists CreateReservationReplyType = iota
	AccommodationNotExist

	AccommodationAvailable
	AccommodationNotAvailable

	UserExists
	UserNotExist

	AvailabilityReverted
	AvailabilityNotReverted

	AutoApproveReservation
	ManualPendingReservation

	ReservationPending
	ReservationApproved
	ReservationCancelled
	ReservationDeleted
	UnknownReply
)

type CreateReservationReply struct {
	Reservation Reservation
	Type        CreateReservationReplyType
}
