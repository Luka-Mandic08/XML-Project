package startup

import (
	"auth_service/domain/repository"
	"auth_service/domain/service"
	"fmt"
	"log"
	"net"

	"auth_service/infrastructure/api"
	"auth_service/infrastructure/persistence"

	auth "common/proto/auth_service"

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
	userStore := server.initAuthStore(mongoClient)
	userService := server.initAuthService(userStore)
	userHandler := server.initAuthHandler(userService)
	server.startGrpcServer(userHandler)
}

func (server *Server) initMongoClient() *mongo.Client {
	client, err := persistence.GetClient(server.config.AuthDBHost, server.config.AuthDBPort)
	if err != nil {
		log.Fatal(err)
	}
	return client
}

func (server *Server) initAuthStore(client *mongo.Client) repository.AuthStore {
	return repository.NewAuthMongoDBStore(client)
}

func (server *Server) initAuthService(store repository.AuthStore) *service.AuthService {
	return service.NewAuthService(store)
}

func (server *Server) initAuthHandler(service *service.AuthService) *api.AuthHandler {
	return api.NewAuthHandler(service)
}

func (server *Server) startGrpcServer(authHandler *api.AuthHandler) {
	listener, err := net.Listen("tcp", fmt.Sprintf(":%s", server.config.Port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	grpcServer := grpc.NewServer()
	auth.RegisterAuthServiceServer(grpcServer, authHandler)
	if grpcServer == nil || authHandler == nil {
		fmt.Println("GRESKA")
		fmt.Println("GRESKA")
		fmt.Println("GRESKA")
		fmt.Println("GRESKA")
	}
	if err := grpcServer.Serve(listener); err != nil {
		log.Fatalf("failed to serve: %s", err)
	}
}
