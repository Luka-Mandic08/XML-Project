package startup

import (
	reservation "common/proto/reservation_service"
	"fmt"
	"log"
	"net"
	"reservation_service/domain/repository"
	"reservation_service/domain/service"
	"reservation_service/infrastructure/api"
	"reservation_service/infrastructure/persistence"

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
	reservationStore := server.initReservationStore(mongoClient)

	reservationService := server.initReservationService(reservationStore)

	reservationHandler := server.initReservationHandler(reservationService)

	server.startGrpcServer(reservationHandler)
}

func (server *Server) initMongoClient() *mongo.Client {
	client, err := persistence.GetClient(server.config.ReservationDBHost, server.config.ReservationDBPort)
	if err != nil {
		log.Fatal(err)
	}
	return client
}

func (server *Server) initReservationStore(client *mongo.Client) repository.ReservationStore {
	return repository.NewReservationMongoDBStore(client)
}

func (server *Server) initReservationService(store repository.ReservationStore) *service.ReservationService {
	return service.NewReservationService(store)
}

func (server *Server) initReservationHandler(service *service.ReservationService) *api.ReservationHandler {
	return api.NewReservationHandler(service)
}

func (server *Server) startGrpcServer(reservationHandler *api.ReservationHandler) {
	listener, err := net.Listen("tcp", fmt.Sprintf(":%s", server.config.Port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	grpcServer := grpc.NewServer()
	reservation.RegisterReservationServiceServer(grpcServer, reservationHandler)
	if err := grpcServer.Serve(listener); err != nil {
		log.Fatalf("failed to serve: %s", err)
	}
}
