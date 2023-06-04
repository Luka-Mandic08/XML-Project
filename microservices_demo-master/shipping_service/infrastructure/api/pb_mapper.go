package api

import (
	pb "github.com/tamararankovic/microservices_demo/common/proto/shipping_service"
	"github.com/tamararankovic/microservices_demo/shipping_service/domain"
)

func mapOrder(order *domain.Order) *pb.Order {
	orderPb := &pb.Order{
		Id:              order.Id.Hex(),
		Status:          mapStatus(order.Status),
		ShippingAddress: order.ShippingAddress,
	}
	return orderPb
}

func mapStatus(status domain.OrderStatus) pb.Order_OrderStatus {
	switch status {
	case domain.Scheduled:
		return pb.Order_Scheduled
	case domain.InTransport:
		return pb.Order_InTransport
	case domain.Delivered:
		return pb.Order_Delivered
	}
	return pb.Order_Cancelled
}
