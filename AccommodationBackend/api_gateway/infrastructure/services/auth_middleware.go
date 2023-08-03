package services

import (
	"api_gateway/domain/model"
	"github.com/gin-gonic/gin"
)

func ValidateToken() gin.HandlerFunc {
	return func(context *gin.Context) {
		tokenString := context.Request.Header.Get("Authorization")
		valid, claims := model.VerifyToken(tokenString)
		if !valid {
			context.AbortWithStatus(401)
		}
		if len(context.Keys) == 0 {
			context.Keys = make(map[string]interface{})
		}
		context.Keys["Username"] = claims.Username
		context.Keys["Role"] = claims.Role
		context.Keys["UserId"] = claims.UserId

	}
}

func AuthorizeRole(allowedRole string) gin.HandlerFunc {
	return func(context *gin.Context) {
		if len(context.Keys) == 0 {
			context.AbortWithStatus(403)
		}

		role := context.Keys["Role"]
		if role != allowedRole {
			context.AbortWithStatus(403)
		}
	}
}

func AuthorizeId(id string, context *gin.Context) bool {
	if len(context.Keys) == 0 {
		return true
	}
	UserId := context.Keys["UserId"]
	if id != UserId {
		return true
	}
	return false
}
