package service

import (
	"context"

	"github.com/batuhancaam/todo-app/model"
	"github.com/batuhancaam/todo-app/user/data/request"
)

type UserService interface {
	SignUp(createUserReq request.CreateUserRequest)
	SingIn(signInReq request.SignInRequest) (string, error)
	ParseToken(ctx context.Context, accessToken string) (*model.User, error)
}
