package api

import (
	events "github.com/tamararankovic/microservices_demo/common/saga/create_order"
	saga "github.com/tamararankovic/microservices_demo/common/saga/messaging"
	"github.com/tamararankovic/microservices_demo/shipping_service/application"
	"github.com/tamararankovic/microservices_demo/shipping_service/domain"
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
	order := &domain.Order{Id: id, ShippingAddress: command.Order.Address}

	reply := events.CreateOrderReply{Order: command.Order}

	switch command.Type {
	case events.ShipOrder:
		err := handler.orderService.Create(order)
		if err != nil {
			reply.Type = events.OrderShippingNotScheduled
			break
		}
		reply.Type = events.OrderShippingScheduled
	default:
		reply.Type = events.UnknownReply
	}

	if reply.Type != events.UnknownReply {
		_ = handler.replyPublisher.Publish(reply)
	}
}
