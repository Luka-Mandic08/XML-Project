package api

import (
	events "github.com/tamararankovic/microservices_demo/common/saga/create_order"
	saga "github.com/tamararankovic/microservices_demo/common/saga/messaging"
	"github.com/tamararankovic/microservices_demo/inventory_service/application"
)

type CreateOrderCommandHandler struct {
	productService    *application.ProductService
	replyPublisher    saga.Publisher
	commandSubscriber saga.Subscriber
}

func NewCreateOrderCommandHandler(productService *application.ProductService, publisher saga.Publisher, subscriber saga.Subscriber) (*CreateOrderCommandHandler, error) {
	o := &CreateOrderCommandHandler{
		productService:    productService,
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
	reply := events.CreateOrderReply{Order: command.Order}

	switch command.Type {
	case events.UpdateInventory:
		products := mapUpdateProducts(command)
		err := handler.productService.UpdateQuantityForAll(products)
		if err != nil {
			reply.Type = events.InventoryNotUpdated
			break
		}
		reply.Type = events.InventoryUpdated
	case events.RollbackInventory:
		products := mapRollbackProducts(command)
		err := handler.productService.UpdateQuantityForAll(products)
		if err != nil {
			return
		}
		reply.Type = events.InventoryRolledBack
	default:
		reply.Type = events.UnknownReply
	}

	if reply.Type != events.UnknownReply {
		_ = handler.replyPublisher.Publish(reply)
	}
}
