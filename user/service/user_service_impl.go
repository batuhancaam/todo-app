package service

import (
	"time"

	"github.com/batuhancaam/todo-app/helper"
	"github.com/batuhancaam/todo-app/model"
	"github.com/batuhancaam/todo-app/user/data/request"
	"github.com/batuhancaam/todo-app/user/repository"
	"github.com/batuhancaam/todo-app/user/utils"
	"github.com/golang-jwt/jwt/v5"

	"github.com/go-playground/validator"
	"github.com/spf13/viper"
)

type UserServiceImpl struct {
	userRepo  repository.UserRepository
	validator *validator.Validate
}

// SingIn implements UserService.

func NewUserServiceImpl(userRepository repository.UserRepository, validate *validator.Validate) UserService {

	return &UserServiceImpl{
		userRepo:  userRepository,
		validator: validate,
	}
}

// SignUp implements UserService.
func (u *UserServiceImpl) SignUp(createUserReq request.CreateUserRequest) {
	err := u.validator.Struct(createUserReq)
	helper.ErrorPanic(err)

	pwd, err := utils.HashPassword(createUserReq.Password)

	helper.ErrorPanic(err)

	userModel := model.User{
		Email:    createUserReq.Email,
		Password: pwd,
	}

	u.userRepo.CreateUser(userModel)
}

func (u *UserServiceImpl) SingIn(signInReq request.SignInRequest) (string, error) {
	err := u.validator.Struct(signInReq)
	helper.ErrorPanic(err)

	user, err := u.userRepo.GetUser(signInReq.Email)
	helper.ErrorPanic(err)

	if !utils.VerifyPassword(signInReq.Password, user.Password) {
		return "", err
	}
	var secretKey = []byte("secret-key")
	expireDuration := time.Second * viper.GetDuration("jwt.expire_time")

	token := jwt.NewWithClaims(jwt.SigningMethodHS256,
		jwt.MapClaims{
			"email": user.Email,
			"exp":   time.Now().Add(expireDuration),
		})

	return token.SignedString(secretKey)

}
