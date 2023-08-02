package api

import (
	dto "api_gateway/domain/dto"
	"api_gateway/domain/model"
	auth "common/proto/auth_service"
	user "common/proto/user_service"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"net/http"
	"time"
)

type AuthHandler struct {
	authClient auth.AuthServiceClient
	userClient user.UserServiceClient
}

func NewAuthHandler(authClient auth.AuthServiceClient, userClient user.UserServiceClient) *AuthHandler {
	return &AuthHandler{authClient: authClient, userClient: userClient}
}

func (handler *AuthHandler) Login(ctx *gin.Context) {
	var credentials auth.LoginRequest
	err := ctx.ShouldBindJSON(&credentials)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, err.Error())
		return
	}
	response, err := handler.authClient.Login(ctx, &credentials)
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
	claims := model.JwtClaims{
		Username:       response.Username,
		Role:           response.Role,
		UserId:         response.Userid,
		StandardClaims: jwt.StandardClaims{},
	}
	token, tokenError := model.GenerateToken(&claims, time.Now().UTC().Add(time.Duration(30)*time.Minute))
	if tokenError != nil {
		ctx.JSON(http.StatusBadRequest, tokenError)
		return
	}
	ctx.JSON(http.StatusOK, token)
	return
}

func (handler *AuthHandler) Register(ctx *gin.Context) {
	var registerDto dto.RegisterDto
	err := ctx.ShouldBindJSON(&registerDto)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, err.Error())
		return
	}

	userResponse, err := handler.userClient.Create(ctx, dto.RegisterDtoToUser(&registerDto))
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

	response, err := handler.authClient.Register(ctx, dto.RegisterDtoToAccount(&registerDto, userResponse.Id))
	if err != nil {
		_, err := handler.userClient.Delete(ctx, &user.GetRequest{Id: userResponse.Id})
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

func (handler *AuthHandler) Update(ctx *gin.Context) {
	var account auth.UpdateRequest
	err := ctx.ShouldBindJSON(&account)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, err.Error())
		return
	}
	response, err := handler.authClient.Update(ctx, &account)
	if err != nil {
		grpcError, ok := status.FromError(err)
		if ok {
			ctx.JSON(http.StatusBadRequest, grpcError.Message())
			return
		}
		ctx.JSON(http.StatusBadRequest, err.Error())
		return
	}
	ctx.JSON(http.StatusOK, response)
}

func (handler *AuthHandler) Delete(ctx *gin.Context) {
	//TODO izvrsiti provere da li se nalog sme obrisati
	id := ctx.Param("id")
	request := user.GetRequest{Id: id}
	_, err := handler.userClient.Delete(ctx, &request)
	if err != nil {
		grpcError, ok := status.FromError(err)
		if ok {
			ctx.JSON(http.StatusBadRequest, grpcError.Message())
			return
		}
		ctx.JSON(http.StatusBadRequest, err.Error())
		return
	}
	req := auth.DeleteRequest{Id: id}
	response, err := handler.authClient.Delete(ctx, &req)
	if err != nil {
		grpcError, ok := status.FromError(err)
		if ok {
			ctx.JSON(http.StatusBadRequest, grpcError.Message())
			return
		}
		ctx.JSON(http.StatusBadRequest, err.Error())
		return
	}
	ctx.JSON(http.StatusOK, response)
}
