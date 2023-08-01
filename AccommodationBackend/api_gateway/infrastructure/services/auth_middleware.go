package services

import (
	"api_gateway/domain"
	"github.com/gin-gonic/gin"
)

func ValidateToken() gin.HandlerFunc {
	return func(context *gin.Context) {
		tokenString := context.Request.Header.Get("Authorization")
		valid, claims := domain.VerifyToken(tokenString)
		if !valid {
			context.AbortWithStatus(401)
		}
		if len(context.Keys) == 0 {
			context.Keys = make(map[string]interface{})
		}
		context.Keys["Username"] = claims.Username
		context.Keys["Role"] = claims.Role
		context.Keys["userId"] = claims.UserId

	}
}

func Authorize(allowedRole string) gin.HandlerFunc {
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
