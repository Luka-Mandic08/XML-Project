package api

import (
	events "github.com/tamararankovic/microservices_demo/common/saga/create_order"
	saga "github.com/tamararankovic/microservices_demo/common/saga/messaging"
	"github.com/tamararankovic/microservices_demo/ordering_service/application"
	"github.com/tamararankovic/microservices_demo/ordering_service/domain"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type CreateOrderCommandHandler struct {
	orderService      *application.OrderService
	replyPublisher    saga.Publisher
	commandSubscriber saga.Subscriber
}

func NewCreateOrderCommandHandler(orderService *application.OrderService, publisher saga.Publisher, subscriber saga.Subscriber) (*CreateOrderCommandHandler, error) {
	o := &CreateOrderCommandHandler{
		orderService:      orderService,
		replyPublisher:    publisher,
		commandSubscriber: subscriber,
	}
	err := o.commandSubscriber.Subscribe(o.handle)
	if err != nil {
		return nil, err
	}
	return o, nil
}

func (handler *CreateOrderCommandHandler) handle(command *events.CreateOrderCommand) {
	id, err := primitive.ObjectIDFromHex(command.Order.Id)
	if err != nil {
		return
	}
	order := &domain.Order{Id: id}

	reply := events.CreateOrderReply{Order: command.Order}

	switch command.Type {
	case events.ApproveOrder:
		err := handler.orderService.Approve(order)
		if err != nil {
			return
		}
		reply.Type = events.OrderApproved
	case events.CancelOrder:
		err := handler.orderService.Cancel(order)
		if err != nil {
			return
		}
		reply.Type = events.OrderCancelled
	default:
		reply.Type = events.UnknownReply
	}

	if reply.Type != events.UnknownReply {
		_ = handler.replyPublisher.Publish(reply)
	}
}
