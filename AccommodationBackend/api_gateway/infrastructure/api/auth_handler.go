package api

import (
	dto "api_gateway/domain/dto"
	"api_gateway/domain/model"
	"api_gateway/infrastructure/services"
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
		handler.userClient.Delete(ctx, &user.DeleteRequest{Id: userResponse.Id, Role: registerDto.Role})
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
	if services.AuthorizeId(account.Userid, ctx) {
		ctx.JSON(http.StatusUnauthorized, "Not allowed")
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
	id := ctx.Param("userId")
	if services.AuthorizeId(id, ctx) {
		ctx.JSON(http.StatusUnauthorized, "Not allowed")
		return
	}
	role := ctx.Param("role")
	if role != "Host" && role != "Guest" {
		ctx.JSON(http.StatusBadRequest, "Role must be 'Host' or 'Guest'")
		return
	}
	request := user.DeleteRequest{Id: id, Role: role}
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

func (handler *AuthHandler) GetByUserId(ctx *gin.Context) {
	userId := ctx.Param("userId")
	if services.AuthorizeId(userId, ctx) {
		ctx.JSON(http.StatusUnauthorized, "Not allowed")
		return
	}
	request := auth.GetByUserIdRequest{UserId: userId}
	response, err := handler.authClient.GetByUserId(ctx, &request)
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

func (handler *AuthHandler) GenerateAPIKey(ctx *gin.Context) {
	userId := ctx.Param("userId")
	if services.AuthorizeId(userId, ctx) {
		ctx.JSON(http.StatusUnauthorized, "Not allowed to generate APIKey")
		return
	}
	request := auth.GenerateAPIKeyRequest{UserId: userId}
	response, err := handler.authClient.GenerateAPIKey(ctx, &request)
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

func (handler *AuthHandler) LinkAPIKey(ctx *gin.Context) {
	userId := ctx.Param("userId")
	if services.AuthorizeId(userId, ctx) {
		ctx.JSON(http.StatusUnauthorized, "Not allowed to generate APIKey")
		return
	}
	request := auth.LinkAPIKeyRequest{UserId: userId}
	response, err := handler.authClient.LinkAPIKey(ctx, &request)
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
