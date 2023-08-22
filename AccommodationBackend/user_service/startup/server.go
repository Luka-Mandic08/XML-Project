package startup

import (
	rating "common/proto/rating_service"
	reservation "common/proto/reservation_service"
	saga "common/saga/messaging"
	"common/saga/messaging/nats"
	"fmt"
	"log"
	"net"
	"user_service/domain/repository"
	"user_service/domain/service"

	user "common/proto/user_service"
	"user_service/infrastructure/api"
	"user_service/infrastructure/persistence"

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
	QUEUE_GROUP = "user_service"
)

func (server *Server) Start() {
	mongoClient := server.initMongoClient()
	userStore := server.initUserStore(mongoClient)

	userService := server.initUserService(userStore)

	commandSubscriber := server.initSubscriber(server.config.CreateReservationCommandSubject, QUEUE_GROUP)
	replyPublisher := server.initPublisher(server.config.CreateReservationReplySubject)
	server.initCreateReservationHandler(userService, replyPublisher, commandSubscriber)
	reservationClient := persistence.NewReservationClient(server.config.ReservationHost, server.config.ReservationPort)
	ratingClient := persistence.NewRatingClient(server.config.RatingHost, server.config.RatingPort)
	userHandler := server.initUserHandler(userService, reservationClient, ratingClient)

	server.startGrpcServer(userHandler)
}

func (server *Server) initMongoClient() *mongo.Client {
	client, err := persistence.GetClient(server.config.UserDBHost, server.config.UserDBPort)
	if err != nil {
		log.Fatal(err)
	}
	return client
}

func (server *Server) initUserStore(client *mongo.Client) repository.UserStore {
	return repository.NewUserMongoDBStore(client)
}

func (server *Server) initUserService(store repository.UserStore) *service.UserService {
	return service.NewUserService(store)
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
func (server *Server) initCreateReservationHandler(service *service.UserService, publisher saga.Publisher, subscriber saga.Subscriber) {
	_, err := api.NewCreateReservationCommandHandler(service, publisher, subscriber)
	if err != nil {
		log.Fatal(err)
	}
}

func (server *Server) initUserHandler(service *service.UserService, reservationClient reservation.ReservationServiceClient, ratingClient rating.RatingServiceClient) *api.UserHandler {
	return api.NewUserHandler(service, reservationClient, ratingClient)
}

func (server *Server) startGrpcServer(userHandler *api.UserHandler) {
	listener, err := net.Listen("tcp", fmt.Sprintf(":%s", server.config.Port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	grpcServer := grpc.NewServer()
	user.RegisterUserServiceServer(grpcServer, userHandler)
	if err := grpcServer.Serve(listener); err != nil {
		log.Fatalf("failed to serve: %s", err)
	}
}
