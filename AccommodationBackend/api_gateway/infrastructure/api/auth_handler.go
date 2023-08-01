package api

import (
	"api_gateway/domain"
	auth "common/proto/auth_service"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"net/http"
	"time"
)

type AuthHandler struct {
	client auth.AuthServiceClient
}

func NewAuthHandler(client auth.AuthServiceClient) *AuthHandler {

	return &AuthHandler{client: client}
}

func (handler *AuthHandler) Login(ctx *gin.Context) {
	var credentials auth.LoginRequest
	error := ctx.ShouldBindJSON(&credentials)
	if error != nil {
		ctx.JSON(http.StatusBadRequest, error.Error())
		return
	}

	response, err := handler.client.Login(ctx, &credentials)
	if err != nil {
		grpcError, ok := status.FromError(err)
		if ok {
			switch grpcError.Code() {
			case codes.Unauthenticated:
				ctx.JSON(http.StatusBadRequest, grpcError.Message())
				return
			default:
				ctx.JSON(http.StatusBadRequest, grpcError.Message())
				return
			}
		}
		ctx.JSON(http.StatusBadRequest, err.Error())
		return
	}
	claims := domain.JwtClaims{
		Username:       response.Account.Username,
		Role:           response.Account.Role,
		UserId:         response.Account.Userid,
		StandardClaims: jwt.StandardClaims{},
	}
	token, token_error := domain.GenerateToken(&claims, time.Now().UTC().Add(time.Duration(30)*time.Minute))
	if token_error != nil {
		ctx.JSON(http.StatusBadRequest, token_error)
		return
	}
	ctx.JSON(http.StatusOK, token)
	return
}

func (handler *AuthHandler) Register(ctx *gin.Context) {
	var account auth.RegisterRequest
	error := ctx.ShouldBindJSON(&account)
	if error != nil {
		ctx.JSON(http.StatusBadRequest, error.Error())
		return
	}

	response, err := handler.client.Register(ctx, &account)
	if err != nil {
		grpcError, ok := status.FromError(err)
		if ok {
			switch grpcError.Code() {
			case codes.AlreadyExists:
				ctx.JSON(http.StatusConflict, grpcError.Message())
				return
			default:
				ctx.JSON(http.StatusBadRequest, grpcError.Message())
				return
			}
		}
		ctx.JSON(http.StatusBadRequest, err.Error())
		return
	}
	ctx.JSON(http.StatusCreated, response)
}
