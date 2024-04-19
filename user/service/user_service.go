package service

import (
	"github.com/batuhancaam/todo-app/model"
	"github.com/batuhancaam/todo-app/user/data/request"
	"github.com/dgrijalva/jwt-go/v4"
)

type UserService interface {
	SignUp(createUserReq request.CreateUserRequest)
	Login(signInReq request.LoginRequest) (string, error)
	ParseToken(accessToken string) (*jwt.Token, error)
	GetCurrentUser(accessToken string) (user *model.User, err error)
}
