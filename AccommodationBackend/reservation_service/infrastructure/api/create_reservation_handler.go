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
	case events.CheckUserAvailable:
		println("Command: events.CheckUserAvailable")
		isAvailable, err := handler.reservationService.IsUserAvailable(command.Reservation.UserId, command.Reservation.Start, command.Reservation.End)
		if err != nil {
			reply.Type = events.UserNotAvailable
			println("Reply: events.UserNotAvailable")
		}
		if isAvailable {
			reply.Type = events.UserAvailable
			println("Reply: events.UserAvailable")
			break
		}
		reply.Type = events.UserNotAvailable
		println("Reply: events.UserNotAvailable")
	case events.CancelReservation:
		println("Command: events.CancelReservation")
		_, err := handler.reservationService.AutoCancel(id, command.Reservation.Price)
		if err != nil {
			return
		}
		reply.Type = events.ReservationCancelled
		println("Reply: events.ReservationCancelled")
	case events.ApproveReservation:
		println("Command: events.ApproveReservation")
		_, err := handler.reservationService.AutoApprove(id, command.Reservation.Price)
		if err != nil {
			return
		}
		reply.Type = events.ReservationApproved
		println("Reply: events.ReservationApproved")
	case events.PendingReservation:
		println("Command: events.PendingReservation")
		_, err := handler.reservationService.AutoPending(id, command.Reservation.Price)
		if err != nil {
			return
		}
		reply.Type = events.ReservationPending
		println("Reply: events.ReservationPending")
	case events.DeleteReservation:
		println("Command: events.DeleteReservation")
		_, err = handler.reservationService.Delete(id.Hex())
		if err != nil {
			return
		}
		reply.Type = events.ReservationDeleted
		println("Reply: events.ReservationDeleted")
	default:
		reply.Type = create_reservation.UnknownReply
	}

	if reply.Type != create_reservation.UnknownReply {
		_ = handler.replyPublisher.Publish(reply)
	}
}
