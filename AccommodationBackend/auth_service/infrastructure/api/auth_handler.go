package api

import (
	"auth_service/domain/service"
	"context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

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
	account, error := handler.service.GetByUsername(request.GetCredentials().Username)
	if error != nil {
		return nil, error
	}
	if account != nil {
		if account.Password == request.Credentials.Password {
			return LoginMapper(account), nil
		}
	}
	return nil, status.Error(codes.Unauthenticated, "Wrong username or password")

}

func (handler *AuthHandler) Register(ctx context.Context, request *pb.RegisterRequest) (*pb.RegisterResponse, error) {
	existingAccount, _ := handler.service.GetByUsername(request.GetDto().Username)
	if existingAccount != nil {
		return nil, status.Error(codes.AlreadyExists, "An account with this username already exists")
	}
	account := RegisterMapper(request)
	account, error := handler.service.Insert(account)
	if error != nil {
		return nil, error
	}
	return &pb.RegisterResponse{Id: account.Id.String()}, nil
}
