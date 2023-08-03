package api

import (
	"api_gateway/infrastructure/services"
	user "common/proto/user_service"
	"github.com/gin-gonic/gin"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"net/http"
)

type UserHandler struct {
	client user.UserServiceClient
}

func NewUserHandler(client user.UserServiceClient) *UserHandler {
	return &UserHandler{client: client}
}

func (handler *UserHandler) Get(ctx *gin.Context) {
	id := ctx.Param("id")
	request := user.GetRequest{Id: id}
	response, err := handler.client.Get(ctx, &request)
	if err != nil {
		grpcError, ok := status.FromError(err)
		if ok {
			switch grpcError.Code() {
			case codes.AlreadyExists:
				ctx.JSON(http.StatusConflict, grpcError.Message())
				return
			}
		}
		ctx.JSON(http.StatusBadRequest, err.Error())
		return
	}
	ctx.JSON(http.StatusOK, response)
}

func (handler *UserHandler) Update(ctx *gin.Context) {
	var user user.UpdateRequest
	err := ctx.ShouldBindJSON(&user)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, err.Error())
		return
	}
	if services.AuthorizeId(user.Id, ctx) {
		ctx.JSON(http.StatusUnauthorized, "Not allowed")
		return
	}
	response, err := handler.client.Update(ctx, &user)
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
