package startup

import (
	"fmt"
	"github.com/neo4j/neo4j-go-driver/v5/neo4j"
	"log"
	"net"
	"rating_service/domain/repository"
	"rating_service/domain/service"

	"rating_service/infrastructure/api"
	"rating_service/infrastructure/persistence"

	rating "common/proto/rating_service"
	reservation "common/proto/reservation_service"

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
	neo4jDriver, err := server.initNeo4jClient()
	if err != nil {
		panic("kurac")
	}
	guestAccommodationGraphStore := server.initGuestAccommodationGraphStore(neo4jDriver)
	hostRatingStore := server.initHostRatingStore(mongoClient)
	accommodationRatingStore := server.initAccommodationRatingStore(mongoClient)
	ratingService := server.initRatingService(hostRatingStore, accommodationRatingStore, guestAccommodationGraphStore)
	reservationClient := persistence.NewReservationClient(server.config.ReservationHost, server.config.ReservationPort)
	ratingHandler := server.initRatingHandler(ratingService, reservationClient)
	server.startGrpcServer(ratingHandler)
}

func (server *Server) initMongoClient() *mongo.Client {
	client, err := persistence.GetClient(server.config.RatingDBHost, server.config.RatingDBPort)
	if err != nil {
		log.Fatal(err)
	}
	return client
}

func (server *Server) initNeo4jClient() (*neo4j.DriverWithContext, error) {
	// Local instance
	auth := neo4j.BasicAuth(server.config.GraphUsername, server.config.GraphPassword, "")

	driver, err := neo4j.NewDriverWithContext(server.config.GraphUri, auth)
	if err != nil {
		return nil, err
	}

	return &driver, nil
}

func (server *Server) initGuestAccommodationGraphStore(driver *neo4j.DriverWithContext) repository.GuestAccommodationGraphStore {
	return repository.NewGuestAccommodationGraphStore(*driver)
}

func (server *Server) initHostRatingStore(client *mongo.Client) repository.HostRatingStore {
	return repository.NewHostRatingMongoDBStore(client)
}

func (server *Server) initAccommodationRatingStore(client *mongo.Client) repository.AccommodationRatingStore {
	return repository.NewAccommodationMongoDBStore(client)
}

func (server *Server) initRatingService(hostStore repository.HostRatingStore, accommodationStore repository.AccommodationRatingStore, guestAccommodationGraphStore repository.GuestAccommodationGraphStore) *service.RatingService {
	return service.NewRatingService(hostStore, accommodationStore, guestAccommodationGraphStore)
}

func (server *Server) initRatingHandler(service *service.RatingService, reservationClient reservation.ReservationServiceClient) *api.RatingHandler {
	return api.NewRatingHandler(service, reservationClient)
}

func (server *Server) startGrpcServer(ratingHandler *api.RatingHandler) {
	listener, err := net.Listen("tcp", fmt.Sprintf(":%s", server.config.Port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	grpcServer := grpc.NewServer()
	rating.RegisterRatingServiceServer(grpcServer, ratingHandler)
	if err := grpcServer.Serve(listener); err != nil {
		log.Fatalf("failed to serve: %s", err)
	}
}
