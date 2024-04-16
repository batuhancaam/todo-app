package controller

import (
	"net/http"

	"github.com/batuhancaam/todo-app/helper"
	"github.com/batuhancaam/todo-app/user/data/request"
	"github.com/batuhancaam/todo-app/user/data/response"
	"github.com/batuhancaam/todo-app/user/service"
	"github.com/gin-gonic/gin"
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

func (c *UserController) SignIn(ctx *gin.Context) {

	signInRequest := request.SignInRequest{}

	err := ctx.ShouldBindJSON(&signInRequest)

	token, err := c.userService.SingIn(signInRequest)

	helper.ErrorPanic(err)

	ctx.JSON(http.StatusOK, response.SignInResponse{Token: token})
}
