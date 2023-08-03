package startup

import (
	"accommodation_service/domain/repository"
	"accommodation_service/domain/service"
	"fmt"
	"log"
	"net"

	"accommodation_service/infrastructure/api"
	"accommodation_service/infrastructure/persistence"

	accommodation "common/proto/accommodation_service"

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
	userStore := server.initAccommodationStore(mongoClient)
	userService := server.initAccommodationService(userStore)
	userHandler := server.initAccommodationHandler(userService)
	server.startGrpcServer(userHandler)
}

func (server *Server) initMongoClient() *mongo.Client {
	client, err := persistence.GetClient(server.config.AccommodationDBHost, server.config.AccommodationDBPort)
	if err != nil {
		log.Fatal(err)
	}
	return client
}

func (server *Server) initAccommodationStore(client *mongo.Client) repository.AccommodationStore {
	return repository.NewAccommodationMongoDBStore(client)
}

func (server *Server) initAccommodationService(store repository.AccommodationStore) *service.AccommodationService {
	return service.NewAccommodationService(store)
}

func (server *Server) initAccommodationHandler(service *service.AccommodationService) *api.AccommodationHandler {
	return api.NewAccommodationHandler(service)
}

func (server *Server) startGrpcServer(accommodationHandler *api.AccommodationHandler) {
	listener, err := net.Listen("tcp", fmt.Sprintf(":%s", server.config.Port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	grpcServer := grpc.NewServer()
	accommodation.RegisterAccommodationServiceServer(grpcServer, accommodationHandler)
	if grpcServer == nil || accommodationHandler == nil {
		fmt.Println("GRESKA")
		fmt.Println("GRESKA")
		fmt.Println("GRESKA")
		fmt.Println("GRESKA")
	}
	if err := grpcServer.Serve(listener); err != nil {
		log.Fatalf("failed to serve: %s", err)
	}
}
