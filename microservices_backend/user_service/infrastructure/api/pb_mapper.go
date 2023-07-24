package api

import (
	"user_service/domain"

	pb "github.com/Luka-Mandic08/XML-Project/microservices_backend/common"
)

func mapUser(user *domain.User) *pb.User {
	userPb := &pb.User{
		Id:      user.Id.Hex(),
		Name:    user.Name,
		Surname: user.Surname,
	}
	return userPb
}
