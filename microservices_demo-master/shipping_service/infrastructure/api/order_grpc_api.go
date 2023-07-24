package api

import (
	"context"
	pb "github.com/tamararankovic/microservices_demo/common/proto/shipping_service"
	"github.com/tamararankovic/microservices_demo/shipping_service/application"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type OrderHandler struct {
	pb.UnimplementedShippingServiceServer
	service *application.OrderService
}

func NewOrderHandler(service *application.OrderService) *OrderHandler {
	return &OrderHandler{
		service: service,
	}
}

func (handler *OrderHandler) Get(ctx context.Context, request *pb.GetRequest) (*pb.GetResponse, error) {
	id := request.Id
	objectId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}
	Order, err := handler.service.Get(objectId)
	if err != nil {
		return nil, err
	}
	OrderPb := mapOrder(Order)
	response := &pb.GetResponse{
		Order: OrderPb,
	}
	return response, nil
}

func (handler *OrderHandler) GetAll(ctx context.Context, request *pb.GetAllRequest) (*pb.GetAllResponse, error) {
	Orders, err := handler.service.GetAll()
	if err != nil {
		return nil, err
	}
	response := &pb.GetAllResponse{
		Orders: []*pb.Order{},
	}
	for _, Order := range Orders {
		current := mapOrder(Order)
		response.Orders = append(response.Orders, current)
	}
	return response, nil
}
