package startup

import (
	"fmt"
	shipping "github.com/tamararankovic/microservices_demo/common/proto/shipping_service"
	saga "github.com/tamararankovic/microservices_demo/common/saga/messaging"
	"github.com/tamararankovic/microservices_demo/common/saga/messaging/nats"
	"github.com/tamararankovic/microservices_demo/shipping_service/application"
	"github.com/tamararankovic/microservices_demo/shipping_service/domain"
	"github.com/tamararankovic/microservices_demo/shipping_service/infrastructure/api"
	"github.com/tamararankovic/microservices_demo/shipping_service/infrastructure/persistence"
	"github.com/tamararankovic/microservices_demo/shipping_service/startup/config"
	"go.mongodb.org/mongo-driver/mongo"
	"google.golang.org/grpc"
	"log"
	"net"
)

type Server struct {
	config *config.Config
}

func NewServer(config *config.Config) *Server {
	return &Server{
		config: config,
	}
}

const (
	QueueGroup = "shipping_service"
)

func (server *Server) Start() {
	mongoClient := server.initMongoClient()
	orderStore := server.initOrderStore(mongoClient)

	orderService := server.initOrderService(orderStore)

	commandSubscriber := server.initSubscriber(server.config.CreateOrderCommandSubject, QueueGroup)
	replyPublisher := server.initPublisher(server.config.CreateOrderReplySubject)
	server.initCreateOrderHandler(orderService, replyPublisher, commandSubscriber)

	orderHandler := server.initOrderHandler(orderService)

	server.startGrpcServer(orderHandler)
}

func (server *Server) initMongoClient() *mongo.Client {
	client, err := persistence.GetClient(server.config.ShippingDBHost, server.config.ShippingDBPort)
	if err != nil {
		log.Fatal(err)
	}
	return client
}

func (server *Server) initOrderStore(client *mongo.Client) domain.OrderStore {
	store := persistence.NewOrderMongoDBStore(client)
	store.DeleteAll()
	for _, Order := range orders {
		err := store.Insert(Order)
		if err != nil {
			log.Fatal(err)
		}
	}
	return store
}

func (server *Server) initPublisher(subject string) saga.Publisher {
	publisher, err := nats.NewNATSPublisher(
		server.config.NatsHost, server.config.NatsPort,
		server.config.NatsUser, server.config.NatsPass, subject)
	if err != nil {
		log.Fatal(err)
	}
	return publisher
}

func (server *Server) initSubscriber(subject, queueGroup string) saga.Subscriber {
	subscriber, err := nats.NewNATSSubscriber(
		server.config.NatsHost, server.config.NatsPort,
		server.config.NatsUser, server.config.NatsPass, subject, queueGroup)
	if err != nil {
		log.Fatal(err)
	}
	return subscriber
}

func (server *Server) initOrderService(store domain.OrderStore) *application.OrderService {
	return application.NewOrderService(store)
}

func (server *Server) initCreateOrderHandler(service *application.OrderService, publisher saga.Publisher, subscriber saga.Subscriber) {
	_, err := api.NewCreateOrderCommandHandler(service, publisher, subscriber)
	if err != nil {
		log.Fatal(err)
	}
}

func (server *Server) initOrderHandler(service *application.OrderService) *api.OrderHandler {
	return api.NewOrderHandler(service)
}

func (server *Server) startGrpcServer(orderHandler *api.OrderHandler) {
	listener, err := net.Listen("tcp", fmt.Sprintf(":%s", server.config.Port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	grpcServer := grpc.NewServer()
	shipping.RegisterShippingServiceServer(grpcServer, orderHandler)
	if err := grpcServer.Serve(listener); err != nil {
		log.Fatalf("failed to serve: %s", err)
	}
}
