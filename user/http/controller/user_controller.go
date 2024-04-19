package controller

import (
	"net/http"

	"github.com/batuhancaam/todo-app/helper"
	"github.com/batuhancaam/todo-app/user/data/request"
	"github.com/batuhancaam/todo-app/user/data/response"
	"github.com/batuhancaam/todo-app/user/service"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

type UserController struct {
	userService service.UserService
}

func NewUserController(userService service.UserService) *UserController {

	return &UserController{
		userService: userService,
	}
}

func (c *UserController) SingUp(ctx *gin.Context) {

	createUserRequest := request.CreateUserRequest{}

	err := ctx.ShouldBindJSON(&createUserRequest)

	helper.ErrorPanic(err)

	c.userService.SignUp(createUserRequest)

	ctx.Header("Content-Type", "application/json")
	ctx.Status(http.StatusCreated)
}

func (c *UserController) Login(ctx *gin.Context) {

	loginRequest := request.LoginRequest{}

	err := ctx.ShouldBindJSON(&loginRequest)

	token, err := c.userService.Login(loginRequest)

	helper.ErrorPanic(err)

	ctx.SetCookie("token", token, viper.GetInt("jwt.exp_time"), "/", "localhost", false, true)

	ctx.JSON(http.StatusOK, response.SignInResponse{Token: token})
}

func (c *UserController) Logout(ctx *gin.Context) {

	ctx.SetCookie("token", "", -1, "/", "localhost", false, true)

}
