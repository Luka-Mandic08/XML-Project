package api

import (
	"context"
	"encoding/json"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/tamararankovic/microservices_demo/api_gateway/domain"
	"github.com/tamararankovic/microservices_demo/api_gateway/infrastructure/services"
	catalogue "github.com/tamararankovic/microservices_demo/common/proto/catalogue_service"
	ordering "github.com/tamararankovic/microservices_demo/common/proto/ordering_service"
	shipping "github.com/tamararankovic/microservices_demo/common/proto/shipping_service"
	"net/http"
)

type OrderingHandler struct {
	orderingClientAddress  string
	catalogueClientAddress string
	shippingClientAddress  string
}

func NewOrderingHandler(orderingClientAddress, catalogueClientAddress, shippingClientAddress string) Handler {
	return &OrderingHandler{
		orderingClientAddress:  orderingClientAddress,
		catalogueClientAddress: catalogueClientAddress,
		shippingClientAddress:  shippingClientAddress,
	}
}

func (handler *OrderingHandler) Init(mux *runtime.ServeMux) {
	err := mux.HandlePath("GET", "/order/{orderId}/details", handler.GetDetails)
	if err != nil {
		panic(err)
	}
}

func (handler *OrderingHandler) GetDetails(w http.ResponseWriter, r *http.Request, pathParams map[string]string) {
	id := pathParams["orderId"]
	if id == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	orderDetails := &domain.OrderDetails{Id: id}

	err := handler.addOrderInfo(orderDetails)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	handler.addShippingInfo(orderDetails)
	handler.addProductInfo(orderDetails)

	response, err := json.Marshal(orderDetails)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write(response)
}

func (handler *OrderingHandler) addOrderInfo(orderDetails *domain.OrderDetails) error {
	orderingClient := services.NewOrderingClient(handler.orderingClientAddress)
	orderInfo, err := orderingClient.Get(context.TODO(), &ordering.GetRequest{Id: orderDetails.Id})
	if err != nil {
		return err
	}
	orderDetails.Id = orderInfo.Order.Id
	orderDetails.CreatedAt = orderInfo.Order.CreatedAt.AsTime()
	orderDetails.Status = orderInfo.Order.Status.String()
	orderDetails.OrderItems = make([]domain.OrderItem, 0)
	for _, item := range orderInfo.Order.Items {
		itemDetails := domain.OrderItem{
			Product:  domain.Product{Id: item.Product.Id, ColorCode: item.Product.Color.Code},
			Quantity: uint16(item.Quantity),
		}
		orderDetails.OrderItems = append(orderDetails.OrderItems, itemDetails)
	}
	return nil
}

func (handler *OrderingHandler) addShippingInfo(orderDetails *domain.OrderDetails) {
	shippingClient := services.NewShippingClient(handler.shippingClientAddress)
	shippingInfo, err := shippingClient.Get(context.TODO(), &shipping.GetRequest{Id: orderDetails.Id})
	if err != nil {
		return
	}
	orderDetails.ShippingStatus = shippingInfo.Order.Status.String()
	orderDetails.ShippingAddress = shippingInfo.Order.ShippingAddress
}

func (handler *OrderingHandler) addProductInfo(orderDetails *domain.OrderDetails) {
	for i, item := range orderDetails.OrderItems {
		catalogueClient := services.NewCatalogueClient(handler.catalogueClientAddress)
		catalogueInfo, err := catalogueClient.Get(context.TODO(), &catalogue.GetRequest{Id: item.Product.Id})
		if err != nil {
			continue
		}
		item.Product.Name = catalogueInfo.Product.Name
		item.Product.ClothingBrand = catalogueInfo.Product.ClothingBrand
		item.Product.ColorName = getColorName(item.Product.ColorCode, catalogueInfo.Product.Colors)
		orderDetails.OrderItems[i] = item
	}
}

func getColorName(code string, colors []*catalogue.Color) string {
	for _, color := range colors {
		if color.Code == code {
			return color.Name
		}
	}
	return ""
}
