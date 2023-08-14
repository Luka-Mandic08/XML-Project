package api

import (
	"common/saga/create_reservation"
	events "common/saga/create_reservation"
	saga "common/saga/messaging"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"reservation_service/domain/service"
)

type CreateReservationCommandHandler struct {
	reservationService *service.ReservationService
	replyPublisher     saga.Publisher
	commandSubscriber  saga.Subscriber
}

func NewCreateReservationCommandHandler(reservationService *service.ReservationService, publisher saga.Publisher, subscriber saga.Subscriber) (*CreateReservationCommandHandler, error) {
	o := &CreateReservationCommandHandler{
		reservationService: reservationService,
		replyPublisher:     publisher,
		commandSubscriber:  subscriber,
	}
	err := o.commandSubscriber.Subscribe(o.handle)
	if err != nil {
		return nil, err
	}
	return o, nil
}

func (handler *CreateReservationCommandHandler) handle(command *events.CreateReservationCommand) {
	id, err := primitive.ObjectIDFromHex(command.Reservation.Id)
	if err != nil {
		return
	}

	reply := events.CreateReservationReply{Reservation: command.Reservation}

	switch command.Type {
	case events.CancelReservation:
		println("Command: events.CancelReservation")
		_, err := handler.reservationService.Cancel(id)
		if err != nil {
			return
		}
		reply.Type = events.ReservationCancelled
		println("Reply: events.ReservationCancelled")
	case events.ApproveReservation:
		println("Command: events.ApproveReservation")
		_, err := handler.reservationService.Approve(id, command.Reservation.Price)
		if err != nil {
			return
		}
		reply.Type = events.ReservationApproved
		println("Reply: events.ReservationApproved")
	case events.PendingReservation:
		println("Command: events.PendingReservation")
		reply.Type = events.ReservationPending
		println("Reply: events.ReservationPending")
	default:
		reply.Type = create_reservation.UnknownReply
	}

	if reply.Type != create_reservation.UnknownReply {
		_ = handler.replyPublisher.Publish(reply)
	}
}
