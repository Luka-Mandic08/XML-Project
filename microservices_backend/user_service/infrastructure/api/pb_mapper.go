package api

import (
	pb "common/proto/user_service"
	"user_service/domain"
)

func mapUser(user *domain.User) *pb.User {
	userPb := &pb.User{
		Id:      user.Id.Hex(),
		Name:    user.Name,
		Surname: user.Surname,
	}
	return userPb
}
