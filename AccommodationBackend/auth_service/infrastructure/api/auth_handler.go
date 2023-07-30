package api

import (
	"auth_service/domain/service"
	"context"
	"errors"

	pb "common/proto/auth_service"
)

type AuthHandler struct {
	pb.UnimplementedAuthServiceServer
	service *service.AuthService
}

func NewAuthHandler(service *service.AuthService) *AuthHandler {
	return &AuthHandler{
		service: service,
	}
}

func (handler *AuthHandler) Login(ctx context.Context, request *pb.LoginRequest) (*pb.LoginResponse, error) {
	account, error := handler.service.GetByUsername(request.Credentials.Username)
	if error != nil {
		return nil, error
	}
	if account != nil {
		if account.Password == request.Credentials.Password {
			return &pb.LoginResponse{}, nil
		}
		err := errors.New("wrong username or password")
		return nil, err
	}
	return nil, nil
}

func (handler *AuthHandler) Register(ctx context.Context, request *pb.RegisterRequest) (*pb.RegisterResponse, error) {
	existingAccount, _ := handler.service.GetByUsername(request.Dto.Username)
	if existingAccount != nil {
		return nil, errors.New("username taken")
	}
	account := RegisterMapper(request)
	account, error := handler.service.Insert(account)
	if error != nil {
		return nil, error
	}
	return &pb.RegisterResponse{}, nil
}
