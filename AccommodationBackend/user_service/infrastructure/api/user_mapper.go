package api

import (
	user "common/proto/user_service"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"user_service/domain/model"
)

func MapUserToGetResponse(u *model.User) *user.GetResponse {
	address := user.Address{
		Street:  u.Address.Street,
		City:    u.Address.City,
		Country: u.Address.Country,
	}
	response := user.GetResponse{
		Id:            u.Id.Hex(),
		Name:          u.Name,
		Surname:       u.Surname,
		Email:         u.Email,
		Address:       &address,
		IsOutstanding: false,
	}
	return &response
}

func MapCreateRequestToUser(request *user.CreateRequest) *model.User {
	address := model.Address{
		Street:  request.Address.Street,
		City:    request.Address.City,
		Country: request.Address.Country,
	}
	user := model.User{
		Name:    request.Name,
		Surname: request.Surname,
		Email:   request.Email,
		Address: address,
	}
	return &user
}

func MapUpdateRequestToUser(request *user.UpdateRequest) *model.User {
	address := model.Address{
		Street:  request.Address.Street,
		City:    request.Address.City,
		Country: request.Address.Country,
	}
	id, _ := primitive.ObjectIDFromHex(request.Id)
	user := model.User{
		Id:      id,
		Name:    request.Name,
		Surname: request.Surname,
		Email:   request.Email,
		Address: address,
	}
	return &user
}

func MapUserToGetForReservationResponse(u *model.User) *user.GetForReservationResponse {
	response := user.GetForReservationResponse{
		Name:    u.Name,
		Surname: u.Surname,
		Email:   u.Email,
	}
	return &response
}
