package api

import (
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
