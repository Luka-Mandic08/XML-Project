package create_reservation

import (
	"time"
)

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

// TODO Command
const (
	CheckAccommodationExists CreateReservationCommandType = iota
	CheckAvailableAccommodation
	ChangeAvailability
	CheckUserExists
	CheckAutomaticApproveReservation
	RevertAvailability

	PendingReservation
	ApproveReservation
	CancelReservation
	UnknownCommand
)

type CreateReservationCommand struct {
	Reservation Reservation
	Type        CreateReservationCommandType
}

type CreateReservationReplyType int8

// TODO Reply
const (
	AccommodationExists CreateReservationReplyType = iota
	AccommodationNotExist

	AccommodationAvailable
	AccommodationNotAvailable

	AvailabilityChanged
	AvailabilityNotChanged

	UserExists
	UserNotExist

	//TODO !!!
	AvailabilityReverted
	AvailabilityNotReverted

	AutoApproveReservation
	ManualPendingReservation

	ReservationPending
	ReservationApproved
	ReservationCancelled
	UnknownReply
)

type CreateReservationReply struct {
	Reservation Reservation
	Type        CreateReservationReplyType
}
