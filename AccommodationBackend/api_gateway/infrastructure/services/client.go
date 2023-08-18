package services

import (
	accommodation "common/proto/accommodation_service"
	auth "common/proto/auth_service"
	rating "common/proto/rating_service"
	reservation "common/proto/reservation_service"
	user "common/proto/user_service"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
)

func NewAuthClient(address string) auth.AuthServiceClient {
	conn, err := getConnection(address)
	if err != nil {
		log.Fatalf("Failed to start gRPC connection to Catalogue service: %v", err)
	}
	return auth.NewAuthServiceClient(conn)
}

func NewUserClient(address string) user.UserServiceClient {
	conn, err := getConnection(address)
	if err != nil {
		log.Fatalf("Failed to start gRPC connection to Catalogue service: %v", err)
	}
	return user.NewUserServiceClient(conn)
}

func NewAccommodationClient(address string) accommodation.AccommodationServiceClient {
	conn, err := getConnection(address)
	if err != nil {
		log.Fatalf("Failed to start gRPC connection to Catalogue service: %v", err)
	}
	return accommodation.NewAccommodationServiceClient(conn)
}

func NewReservationClient(address string) reservation.ReservationServiceClient {
	conn, err := getConnection(address)
	if err != nil {
		log.Fatalf("Failed to start gRPC connection to Catalogue service: %v", err)
	}
	return reservation.NewReservationServiceClient(conn)
}

func NewRatingClient(address string) rating.RatingServiceClient {
	conn, err := getConnection(address)
	if err != nil {
		log.Fatalf("Failed to start gRPC connection to Catalogue service: %v", err)
	}
	return rating.NewRatingServiceClient(conn)
}

func getConnection(address string) (*grpc.ClientConn, error) {
	return grpc.Dial(address, grpc.WithTransportCredentials(insecure.NewCredentials()))
}
