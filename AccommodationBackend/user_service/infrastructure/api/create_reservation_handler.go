package api

import (
	events "common/saga/create_reservation"
	saga "common/saga/messaging"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"user_service/domain/service"
)

type CreateReservationCommandHandler struct {
	userService       *service.UserService
	replyPublisher    saga.Publisher
	commandSubscriber saga.Subscriber
}

func NewCreateReservationCommandHandler(userService *service.UserService, publisher saga.Publisher, subscriber saga.Subscriber) (*CreateReservationCommandHandler, error) {
	o := &CreateReservationCommandHandler{
		userService:       userService,
		replyPublisher:    publisher,
		commandSubscriber: subscriber,
	}
	err := o.commandSubscriber.Subscribe(o.handle)
	if err != nil {
		return nil, err
	}
	return o, nil
}

func (handler *CreateReservationCommandHandler) handle(command *events.CreateReservationCommand) {
	reply := events.CreateReservationReply{Reservation: command.Reservation}

	id, err := primitive.ObjectIDFromHex(command.Reservation.UserId)
	if err != nil {
		reply.Type = events.UserNotExist
		_ = handler.replyPublisher.Publish(reply)
	}

	switch command.Type {
	case events.CheckUserExists:
		println("Command: events.CheckUserExists")
		err := handler.userService.CheckUserExists(id)
		if err != nil {
			reply.Type = events.UserNotExist
			println("Reply: events.UserNotExist")
			println(err.Error())
			break
		}
		reply.Type = events.UserExists
		println("Reply: events.UserExists")
	default:
		reply.Type = events.UnknownReply
	}

	if reply.Type != events.UnknownReply {
		_ = handler.replyPublisher.Publish(reply)
	}
}
