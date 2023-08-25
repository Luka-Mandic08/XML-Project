package persistence

import (
	accommodation "common/proto/accommodation_service"
	rating "common/proto/rating_service"
	user "common/proto/user_service"
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
)

func GetClient(host, port string) (*mongo.Client, error) {
	uri := fmt.Sprintf("mongodb://%s:%s/", host, port)
	options := options.Client().ApplyURI(uri)
	return mongo.Connect(context.TODO(), options)
}

func NewAccommodationClient(host, port string) accommodation.AccommodationServiceClient {
	address := fmt.Sprintf("%s:%s", host, port)
	conn, err := getConnection(address)
	if err != nil {
		log.Fatalf("Failed to start gRPC connection to Accommodation service: %v", err)
	}
	return accommodation.NewAccommodationServiceClient(conn)
}

func NewRatingClient(host, port string) rating.RatingServiceClient {
	address := fmt.Sprintf("%s:%s", host, port)
	conn, err := getConnection(address)
	if err != nil {
		log.Fatalf("Failed to start gRPC connection to Accommodation service: %v", err)
	}
	return rating.NewRatingServiceClient(conn)
}

func NewUserClient(host, port string) user.UserServiceClient {
	address := fmt.Sprintf("%s:%s", host, port)
	conn, err := getConnection(address)
	if err != nil {
		log.Fatalf("Failed to start gRPC connection to User service: %v", err)
	}
	return user.NewUserServiceClient(conn)
}

func getConnection(address string) (*grpc.ClientConn, error) {
	return grpc.Dial(address, grpc.WithTransportCredentials(insecure.NewCredentials()))
}
