package persistence

import (
	reservation "common/proto/reservation_service"
	"context"
	"fmt"
	"github.com/neo4j/neo4j-go-driver/v5/neo4j"
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

func GetDriver(username string, password string, uri string) (neo4j.Driver, error) {
	auth := neo4j.BasicAuth(username, password, "")

	driver, err := neo4j.NewDriver(uri, auth)
	if err != nil {
		return nil, fmt.Errorf("could not create neo4j driver with context: %s", err.Error())
	}

	err = driver.VerifyConnectivity()
	if err != nil {
		println("could not establish connection with neo4j driver: ", err.Error())
		return nil, fmt.Errorf("could not establish connection with neo4j driver: %s", err.Error())
	}

	println("Neo4J server address: ", driver.Target().Host)
	return driver, nil
}

func NewReservationClient(host, port string) reservation.ReservationServiceClient {
	address := fmt.Sprintf("%s:%s", host, port)
	conn, err := getConnection(address)
	if err != nil {
		log.Fatalf("Failed to start gRPC connection to Reservation service: %v", err)
	}
	return reservation.NewReservationServiceClient(conn)
}

func getConnection(address string) (*grpc.ClientConn, error) {
	return grpc.Dial(address, grpc.WithTransportCredentials(insecure.NewCredentials()))
}
