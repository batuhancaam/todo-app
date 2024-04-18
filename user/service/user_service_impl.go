package service

import (
	"context"
	"fmt"
	"time"

	"github.com/batuhancaam/todo-app/helper"
	"github.com/batuhancaam/todo-app/model"
	"github.com/batuhancaam/todo-app/user/data/request"
	"github.com/batuhancaam/todo-app/user/repository"
	"github.com/batuhancaam/todo-app/user/utils"
	"github.com/dgrijalva/jwt-go/v4"

	"github.com/go-playground/validator"
	"github.com/spf13/viper"
)

// TODO : Add logout and refresh token process
type UserServiceImpl struct {
	userRepo  repository.UserRepository
	validator *validator.Validate
}

type AuthClaims struct {
	jwt.StandardClaims
	User *model.User `json:"user"`
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

	secretKey := []byte(viper.GetString("jwt.secret"))

	claims := AuthClaims{
		User: user,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: jwt.At(time.Now().Add(time.Hour * 24)),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString(secretKey)

}
func (u *UserServiceImpl) ParseToken(ctx context.Context, accessToken string) (*model.User, error) {
	token, err := jwt.ParseWithClaims(accessToken, &AuthClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(viper.GetString("jwt.secret")), nil
	})

	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(*AuthClaims); ok && token.Valid {
		return claims.User, nil
	}

	return nil, err
}
