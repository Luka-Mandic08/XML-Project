package api

import (
	"auth_service/domain/model"
	auth "common/proto/auth_service"
)

// Func za mapiranje objekata iz modela na proto message
func RegisterMapper(request *auth.RegisterRequest) *model.Account {
	account := model.Account{
		Username: request.Dto.Username,
		Password: request.Dto.Password,
		Role:     request.Dto.Role,
		UserID:   request.Dto.Userid,
	}
	return &account
}
