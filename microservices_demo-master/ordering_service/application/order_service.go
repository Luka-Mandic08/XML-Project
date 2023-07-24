package application

import (
	"github.com/tamararankovic/microservices_demo/ordering_service/domain"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type OrderService struct {
	store        domain.OrderStore
	orchestrator *CreateOrderOrchestrator
}

func NewOrderService(store domain.OrderStore, orchestrator *CreateOrderOrchestrator) *OrderService {
	return &OrderService{
		store:        store,
		orchestrator: orchestrator,
	}
}

func (service *OrderService) Get(id primitive.ObjectID) (*domain.Order, error) {
	return service.store.Get(id)
}

func (service *OrderService) GetAll() ([]*domain.Order, error) {
	return service.store.GetAll()
}

func (service *OrderService) Create(order *domain.Order, address string) error {
	order.Status = domain.PendingApproval
	order.CreatedAt = time.Now()
	err := service.store.Insert(order)
	if err != nil {
		return err
	}
	err = service.orchestrator.Start(order, address)
	if err != nil {
		order.Status = domain.Cancelled
		_ = service.store.UpdateStatus(order)
		return err
	}
	return nil
}

func (service *OrderService) Approve(order *domain.Order) error {
	order.Status = domain.Approved
	return service.store.UpdateStatus(order)
}

func (service *OrderService) Cancel(order *domain.Order) error {
	order.Status = domain.Cancelled
	return service.store.UpdateStatus(order)
}
