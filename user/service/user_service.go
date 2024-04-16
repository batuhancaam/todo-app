package service

import (
	"github.com/batuhancaam/todo-app/user/data/request"
)

type UserService interface {
	SignUp(createUserReq request.CreateUserRequest)
	SingIn(signInReq request.SignInRequest) (string, error)
}
