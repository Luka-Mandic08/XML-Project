package persistence

import (
	rating "common/proto/rating_service"
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

func NewRatingClient(host, port string) rating.RatingServiceClient {
	address := fmt.Sprintf("%s:%s", host, port)
	conn, err := getConnection(address)
	if err != nil {
		log.Fatalf("Failed to start gRPC connection to Rating service: %v", err)
	}
	return rating.NewRatingServiceClient(conn)
}

func getConnection(address string) (*grpc.ClientConn, error) {
	return grpc.Dial(address, grpc.WithTransportCredentials(insecure.NewCredentials()))
}
