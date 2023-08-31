package startup

import (
	notification "common/proto/notification_service"
	"fmt"
	"log"
	"net"
	"notification_service/domain/repository"
	"notification_service/domain/service"

	"notification_service/infrastructure/api"
	"notification_service/infrastructure/persistence"

	"go.mongodb.org/mongo-driver/mongo"
	"google.golang.org/grpc"
)

type Server struct {
	config *Config
}

func NewServer(config *Config) *Server {
	return &Server{
		config: config,
	}
}

func (server *Server) Start() {
	mongoClient := server.initMongoClient()
	notificationStore := server.initNotificationStore(mongoClient)
	selectedNotificationTypesStore := server.initSelectedNotificationTypesStore(mongoClient)
	notificationService := server.initNotificationService(notificationStore, selectedNotificationTypesStore)
	notificationHandler := server.initNotificationHandler(notificationService)
	server.startGrpcServer(notificationHandler)
}

func (server *Server) initMongoClient() *mongo.Client {
	client, err := persistence.GetClient(server.config.NotificationDBHost, server.config.NotificationDBPort)
	if err != nil {
		log.Fatal(err)
	}
	return client
}

func (server *Server) initNotificationStore(client *mongo.Client) repository.NotificationStore {
	return repository.NewNotificationMongoDBStore(client)
}

func (server *Server) initSelectedNotificationTypesStore(client *mongo.Client) repository.SelectedNotificationTypesStore {
	return repository.NewSelectedNotificationTypesMongoDBStore(client)
}

func (server *Server) initNotificationService(notificationStore repository.NotificationStore, selectedNotificationTypesStore repository.SelectedNotificationTypesStore) *service.NotificationService {
	return service.NewNotificationService(notificationStore, selectedNotificationTypesStore)
}

func (server *Server) initNotificationHandler(service *service.NotificationService) *api.NotificationHandler {
	return api.NewNotificationHandler(service)
}

func (server *Server) startGrpcServer(notificationHandler *api.NotificationHandler) {
	listener, err := net.Listen("tcp", fmt.Sprintf(":%s", server.config.Port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	grpcServer := grpc.NewServer()
	notification.RegisterNotificationServiceServer(grpcServer, notificationHandler)
	if err := grpcServer.Serve(listener); err != nil {
		log.Fatalf("failed to serve: %s", err)
	}
}
