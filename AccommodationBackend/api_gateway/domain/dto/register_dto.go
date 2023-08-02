package dto

import (
	auth "common/proto/auth_service"
	user "common/proto/user_service"
)

type RegisterDto struct {
	Name    string
	Surname string
	Email   string
	Address Address

	Username string
	Password string
	Role     string
}

type Address struct {
	Street  string
	City    string
	Country string
}

func RegisterDtoToUser(dto *RegisterDto) *user.CreateRequest {
	address := user.Address{
		Street:  dto.Address.Street,
		City:    dto.Address.City,
		Country: dto.Address.Country,
	}
	user := user.CreateRequest{
		Name:    dto.Name,
		Surname: dto.Surname,
		Email:   dto.Email,
		Address: &address,
	}
	return &user
}

func RegisterDtoToAccount(dto *RegisterDto, userid string) *auth.RegisterRequest {
	account := auth.RegisterRequest{
		Username: dto.Username,
		Password: dto.Password,
		Role:     dto.Role,
		Userid:   userid,
	}
	return &account
}
