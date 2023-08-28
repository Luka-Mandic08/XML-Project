package service

import (
	events "common/saga/create_reservation"
	saga "common/saga/messaging"
	"reservation_service/domain/model"
)

type CreateReservationOrchestrator struct {
	commandPublisher saga.Publisher
	replySubscriber  saga.Subscriber
}

func NewCreateReservationOrchestrator(publisher saga.Publisher, subscriber saga.Subscriber) (*CreateReservationOrchestrator, error) {
	o := &CreateReservationOrchestrator{
		commandPublisher: publisher,
		replySubscriber:  subscriber,
	}
	err := o.replySubscriber.Subscribe(o.handle)
	if err != nil {
		return nil, err
	}
	return o, nil
}

func (o *CreateReservationOrchestrator) Start(reservation *model.Reservation) error {
	event := &events.CreateReservationCommand{
		Type: events.CheckAccommodationExists,
		Reservation: events.Reservation{
			Id:              reservation.Id.Hex(),
			AccommodationId: reservation.AccommodationId,
			Start:           reservation.Start,
			End:             reservation.End,
			UserId:          reservation.UserId,
			NumberOfGuests:  reservation.NumberOfGuests,
			Status:          reservation.Status,
			Price:           0,
		},
	}
	return o.commandPublisher.Publish(event)
}

func (o *CreateReservationOrchestrator) handle(reply *events.CreateReservationReply) {
	command := events.CreateReservationCommand{Reservation: reply.Reservation}
	command.Type = o.nextCommandType(reply.Type)
	if command.Type != events.UnknownCommand {
		_ = o.commandPublisher.Publish(command)
	}
}

func (o *CreateReservationOrchestrator) nextCommandType(reply events.CreateReservationReplyType) events.CreateReservationCommandType {
	switch reply {
	case events.AccommodationExists:
		return events.CheckAvailableAccommodation
	case events.AccommodationNotExist:
		return events.DeleteReservation

	case events.AccommodationAvailable:
		return events.CheckUserExists
	case events.AccommodationNotAvailable:
		return events.DeleteReservation

	case events.UserExists:
		return events.CheckAutomaticApproveReservation
	case events.UserNotExist:
		return events.RevertAvailability

	case events.AutoApproveReservation:
		return events.ApproveReservation
	case events.ManualPendingReservation:
		return events.PendingReservation

	case events.AvailabilityReverted:
		return events.DeleteReservation
	case events.AvailabilityNotReverted:
		return events.DeleteReservation

	default:
		return events.UnknownCommand
	}
}
