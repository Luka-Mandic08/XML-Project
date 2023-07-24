package startup

import (
	"fmt"
	inventory "github.com/tamararankovic/microservices_demo/common/proto/inventory_service"
	saga "github.com/tamararankovic/microservices_demo/common/saga/messaging"
	"github.com/tamararankovic/microservices_demo/common/saga/messaging/nats"
	"github.com/tamararankovic/microservices_demo/inventory_service/application"
	"github.com/tamararankovic/microservices_demo/inventory_service/domain"
	"github.com/tamararankovic/microservices_demo/inventory_service/infrastructure/api"
	"github.com/tamararankovic/microservices_demo/inventory_service/infrastructure/persistence"
	"github.com/tamararankovic/microservices_demo/inventory_service/startup/config"
	"google.golang.org/grpc"
	"gorm.io/gorm"
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
	QueueGroup = "inventory_service"
)

func (server *Server) Start() {
	postgresClient := server.initPostgresClient()
	productStore := server.initProductStore(postgresClient)

	productService := server.initProductService(productStore)

	commandSubscriber := server.initSubscriber(server.config.CreateOrderCommandSubject, QueueGroup)
	replyPublisher := server.initPublisher(server.config.CreateOrderReplySubject)
	server.initCreateOrderHandler(productService, replyPublisher, commandSubscriber)

	productHandler := server.initProductHandler(productService)

	server.startGrpcServer(productHandler)
}

func (server *Server) initPostgresClient() *gorm.DB {
	client, err := persistence.GetClient(
		server.config.InventoryDBHost, server.config.InventoryDBUser,
		server.config.InventoryDBPass, server.config.InventoryDBName,
		server.config.InventoryDBPort)
	if err != nil {
		log.Fatal(err)
	}
	return client
}

func (server *Server) initProductStore(client *gorm.DB) domain.ProductStore {
	store, err := persistence.NewProductPostgresStore(client)
	if err != nil {
		log.Fatal(err)
	}
	store.DeleteAll()
	for _, Product := range products {
		err := store.Insert(Product)
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

func (server *Server) initProductService(store domain.ProductStore) *application.ProductService {
	return application.NewProductService(store)
}

func (server *Server) initCreateOrderHandler(service *application.ProductService, publisher saga.Publisher, subscriber saga.Subscriber) {
	_, err := api.NewCreateOrderCommandHandler(service, publisher, subscriber)
	if err != nil {
		log.Fatal(err)
	}
}

func (server *Server) initProductHandler(service *application.ProductService) *api.ProductHandler {
	return api.NewProductHandler(service)
}

func (server *Server) startGrpcServer(productHandler *api.ProductHandler) {
	listener, err := net.Listen("tcp", fmt.Sprintf(":%s", server.config.Port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	grpcServer := grpc.NewServer()
	inventory.RegisterInventoryServiceServer(grpcServer, productHandler)
	if err := grpcServer.Serve(listener); err != nil {
		log.Fatalf("failed to serve: %s", err)
	}
}
