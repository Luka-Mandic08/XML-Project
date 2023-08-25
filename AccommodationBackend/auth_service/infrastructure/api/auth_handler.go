package api

import (
	"auth_service/domain/service"
	"context"
	"golang.org/x/crypto/bcrypt"
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
	account, err := handler.service.GetByUsername(request.Username)
	if err != nil {
		return nil, err
	}
	if account != nil {
		err = bcrypt.CompareHashAndPassword([]byte(account.Password), []byte(request.Password))
		if err == nil {
			return LoginMapper(account), nil
		}
	}
	return nil, status.Error(codes.Unauthenticated, "Wrong username or password")

}

func (handler *AuthHandler) Register(ctx context.Context, request *pb.RegisterRequest) (*pb.RegisterResponse, error) {
	existingAccount, _ := handler.service.GetByUsername(request.Username)
	if existingAccount != nil {
		return nil, status.Error(codes.AlreadyExists, "An account with this username already exists")
	}
	account := RegisterMapper(request)
	account, err := handler.service.Insert(account)
	if err != nil {
		return nil, status.Error(codes.AlreadyExists, "Unable to insert account into database")
	}
	return &pb.RegisterResponse{Id: account.Id.String()}, nil
}

func (handler *AuthHandler) Update(ctx context.Context, request *pb.UpdateRequest) (*pb.UpdateResponse, error) {
	account := UpdateMapper(request)
	result, acc, err := handler.service.Update(account)
	if err != nil {
		return nil, status.Error(codes.Unknown, "Unable to update account")
	}
	if result.MatchedCount == 0 {
		return nil, status.Error(codes.NotFound, "Unable to find account")
	}
	return UpdateMapperToAccount(acc), nil
}

func (handler *AuthHandler) Delete(ctx context.Context, request *pb.DeleteRequest) (*pb.DeleteResponse, error) {
	result, err := handler.service.Delete(request.Id)
	if err != nil {
		return nil, status.Error(codes.Unknown, "Unable to delete account")
	}
	if result.DeletedCount == 0 {
		return nil, status.Error(codes.NotFound, "Unable to find account")
	}
	return &pb.DeleteResponse{Message: "Account successfully deleted"}, nil
}

func (handler *AuthHandler) GetByUserId(ctx context.Context, request *pb.GetByUserIdRequest) (*pb.GetByUserIdResponse, error) {
	result, err := handler.service.GetByUserId(request.UserId)
	if err != nil {
		return nil, status.Error(codes.Unknown, "Unable to get account")
	}
	return GetMapper(result), nil
}

func (handler *AuthHandler) GenerateAPIKey(ctx context.Context, request *pb.GenerateAPIKeyRequest) (*pb.GenerateAPIKeyResponse, error) {
	_, err := handler.service.GenerateAPIKey(request.UserId)
	if err != nil {
		return nil, status.Error(codes.Unknown, err.Error())
	}
	return &pb.GenerateAPIKeyResponse{Message: "APIKey generated successfully."}, nil
}

func (handler *AuthHandler) LinkAPIKey(ctx context.Context, request *pb.LinkAPIKeyRequest) (*pb.LinkAPIKeyResponse, error) {
	result, err := handler.service.LinkAPIKey(request.UserId)
	if err != nil {
		return nil, status.Error(codes.Unknown, err.Error())
	}
	return &pb.LinkAPIKeyResponse{ApiKey: result}, nil
}
