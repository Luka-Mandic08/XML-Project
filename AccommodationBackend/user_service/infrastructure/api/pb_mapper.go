package api

import (
	"user_service/domain"

	pb "common/proto/user_service"
)

func mapUser(user *domain.User) *pb.User {
	userPb := &pb.User{
		Id:      user.Id.Hex(),
		Name:    user.Name,
		Surname: user.Surname,
	}
	return userPb
}
