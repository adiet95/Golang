package auth

import (
	"github.com/adiet95/Golang/GoRent/src/database/orm/models"
	"github.com/adiet95/Golang/GoRent/src/helpers"
	"github.com/adiet95/Golang/GoRent/src/interfaces"
)

type auth_service struct {
	repo interfaces.UserRepo
}
type token_response struct {
	Tokens string `json:"token"`
}

func NewService(reps interfaces.UserRepo) *auth_service {
	return &auth_service{reps}
}

func (a auth_service) Login(body models.User) *helpers.Response {
	user, err := a.repo.FindByEmail(body.Email)
	if err != nil {
		return helpers.New("email not registered", 401, true)
	}
	if !helpers.CheckPass(user.Password, body.Password) {
		return helpers.New("wrong password", 401, true)
	}
	token := helpers.NewToken(body.Email, user.Role)
	theToken, err := token.Create()
	if err != nil {
		return helpers.New(err.Error(), 401, true)
	}
	return helpers.New(token_response{Tokens: theToken}, 200, false)
}
