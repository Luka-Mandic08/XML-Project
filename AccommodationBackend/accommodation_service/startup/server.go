package startup

import (
	"accommodation_service/domain/repository"
	"accommodation_service/domain/service"
	rating "common/proto/rating_service"
	saga "common/saga/messaging"
	"common/saga/messaging/nats"
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

const (
	QUEUE_GROUP = "accommodation_service"
)

func (server *Server) Start() {
	mongoClient := server.initMongoClient()
	accommodationStore := server.initAccommodationStore(mongoClient)
	availabilityStore := server.initAvailabilityStore(mongoClient)
	accommodationService := server.initAccommodationService(accommodationStore, *availabilityStore)

	commandSubscriber := server.initSubscriber(server.config.CreateReservationCommandSubject, QUEUE_GROUP)
	replyPublisher := server.initPublisher(server.config.CreateReservationReplySubject)
	server.initCreateReservationHandler(accommodationService, replyPublisher, commandSubscriber)
	ratingClient := persistence.NewRatingClient(server.config.RatingHost, server.config.RatingPort)
	accommodationHandler := server.initAccommodationHandler(accommodationService, ratingClient)
	server.startGrpcServer(accommodationHandler)
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

func (server *Server) initAvailabilityStore(client *mongo.Client) *repository.AvailabilityStore {
	return repository.NewAvailabilityStore(client)
}

func (server *Server) initAccommodationService(accommodationStore repository.AccommodationStore, availabilityStore repository.AvailabilityStore) *service.AccommodationService {
	return service.NewAccommodationService(accommodationStore, availabilityStore)
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
func (server *Server) initCreateReservationHandler(service *service.AccommodationService, publisher saga.Publisher, subscriber saga.Subscriber) {
	_, err := api.NewCreateReservationCommandHandler(service, publisher, subscriber)
	if err != nil {
		log.Fatal(err)
	}
}

func (server *Server) initAccommodationHandler(service *service.AccommodationService, ratingClient rating.RatingServiceClient) *api.AccommodationHandler {
	return api.NewAccommodationHandler(service, ratingClient)
}

func (server *Server) startGrpcServer(accommodationHandler *api.AccommodationHandler) {
	listener, err := net.Listen("tcp", fmt.Sprintf(":%s", server.config.Port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	grpcServer := grpc.NewServer()
	accommodation.RegisterAccommodationServiceServer(grpcServer, accommodationHandler)
	if err := grpcServer.Serve(listener); err != nil {
		log.Fatalf("failed to serve: %s", err)
	}
}
