package startup

import (
	accommodation "common/proto/accommodation_service"
	rating "common/proto/rating_service"
	reservation "common/proto/reservation_service"
	user "common/proto/user_service"
	saga "common/saga/messaging"
	"common/saga/messaging/nats"
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

const (
	QUEUE_GROUP = "reservation_service"
)

func (server *Server) Start() {
	mongoClient := server.initMongoClient()

	reservationStore := server.initReservationStore(mongoClient)
	outstandingHostStore := server.initOutstandingHostStore(mongoClient)

	commandPublisher := server.initPublisher(server.config.CreateReservationCommandSubject)
	replySubscriber := server.initSubscriber(server.config.CreateReservationReplySubject, QUEUE_GROUP)
	createReservationOrchestrator := server.initCreateReservationOrchestrator(commandPublisher, replySubscriber)

	accommodationClient := persistence.NewAccommodationClient(server.config.AccommodationHost, server.config.AccommodationPort)
	ratingClient := persistence.NewRatingClient(server.config.RatingHost, server.config.RatingPort)

	reservationService := server.initReservationService(reservationStore, *outstandingHostStore, createReservationOrchestrator, accommodationClient, ratingClient)

	commandSubscriber := server.initSubscriber(server.config.CreateReservationCommandSubject, QUEUE_GROUP)
	replyPublisher := server.initPublisher(server.config.CreateReservationReplySubject)
	server.initCreateReservationHandler(reservationService, replyPublisher, commandSubscriber)

	userClient := persistence.NewUserClient(server.config.UserHost, server.config.UserPort)
	reservationHandler := server.initReservationHandler(reservationService, userClient)

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

func (server *Server) initOutstandingHostStore(client *mongo.Client) *repository.OutstandingHostMongoDBStore {
	return repository.NewOutstandingHostMongoDBStore(client)
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
func (server *Server) initCreateReservationOrchestrator(publisher saga.Publisher, subscriber saga.Subscriber) *service.CreateReservationOrchestrator {
	orchestrator, err := service.NewCreateReservationOrchestrator(publisher, subscriber)
	if err != nil {
		log.Fatal(err)
	}
	return orchestrator
}

func (server *Server) initReservationService(store repository.ReservationStore, outstandingHostStore repository.OutstandingHostMongoDBStore, reservationOrchestrator *service.CreateReservationOrchestrator, accommodationClient accommodation.AccommodationServiceClient, ratingClient rating.RatingServiceClient) *service.ReservationService {
	return service.NewReservationService(store, outstandingHostStore, reservationOrchestrator, accommodationClient, ratingClient)
}

func (server *Server) initCreateReservationHandler(service *service.ReservationService, publisher saga.Publisher, subscriber saga.Subscriber) {
	_, err := api.NewCreateReservationCommandHandler(service, publisher, subscriber)
	if err != nil {
		log.Fatal(err)
	}
}

func (server *Server) initReservationHandler(reservationService *service.ReservationService, userClient user.UserServiceClient) *api.ReservationHandler {
	return api.NewReservationHandler(reservationService, userClient)
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
