package api

import (
	pb "github.com/tamararankovic/microservices_demo/common/proto/ordering_service"
	"github.com/tamararankovic/microservices_demo/ordering_service/domain"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func mapOrder(order *domain.Order) *pb.Order {
	orderPb := &pb.Order{
		Id:        order.Id.Hex(),
		Status:    mapStatus(order.Status),
		CreatedAt: timestamppb.New(order.CreatedAt),
		Items:     make([]*pb.OrderItem, 0),
	}
	for _, item := range order.Items {
		itemPb := &pb.OrderItem{
			Product: &pb.Product{
				Id: item.Product.Id,
				Color: &pb.Color{
					Code: item.Product.Color.Code,
				},
			},
			Quantity: uint32(item.Quantity),
		}
		orderPb.Items = append(orderPb.Items, itemPb)
	}
	return orderPb
}

func mapNewOrder(orderPb *pb.NewOrder) *domain.Order {
	order := &domain.Order{
		Items: make([]domain.OrderItem, 0),
	}
	for _, itemPb := range orderPb.Items {
		item := domain.OrderItem{
			Product: domain.Product{
				Id: itemPb.Product.Id,
				Color: domain.Color{
					Code: itemPb.Product.Color.Code,
				},
			},
			Quantity: uint16(itemPb.Quantity),
		}
		order.Items = append(order.Items, item)
	}
	return order
}

func mapStatus(status domain.OrderStatus) pb.Order_OrderStatus {
	switch status {
	case domain.PendingApproval:
		return pb.Order_PendingApproval
	case domain.Approved:
		return pb.Order_Approved
	}
	return pb.Order_Cancelled
}
