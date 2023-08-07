package api

import (
	"auth_service/domain/model"
	auth "common/proto/auth_service"
	"golang.org/x/crypto/bcrypt"
)

// Func za mapiranje objekata iz modela na proto message
func RegisterMapper(request *auth.RegisterRequest) *model.Account {
	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(request.Password), bcrypt.DefaultCost)
	account := model.Account{
		Username: request.Username,
		Password: string(hashedPassword),
		Role:     request.Role,
		UserId:   request.Userid,
	}
	return &account
}

func LoginMapper(account *model.Account) *auth.LoginResponse {
	acc := auth.LoginResponse{
		Username: account.Username,
		Role:     account.Role,
		Userid:   account.UserId,
	}
	return &acc
}

func UpdateMapper(req *auth.UpdateRequest) *model.Account {
	acc := model.Account{
		Username: req.Username,
		Password: req.Password,
		UserId:   req.Userid,
	}
	return &acc
}
