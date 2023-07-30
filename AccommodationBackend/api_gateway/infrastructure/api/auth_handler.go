package api

import (
	auth "common/proto/auth_service"
	"github.com/gin-gonic/gin"
	"net/http"
)

type AuthHandler struct {
	client auth.AuthServiceClient
}

func NewAuthHandler(client auth.AuthServiceClient) *AuthHandler {
	return &AuthHandler{client: client}
}

func (handler *AuthHandler) Login(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, nil)
}

func (handler *AuthHandler) Register(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, nil)
}
