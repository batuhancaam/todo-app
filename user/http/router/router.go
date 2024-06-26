package router

import (
	"github.com/batuhancaam/todo-app/user/http/controller"
	"github.com/gin-gonic/gin"
)

func RegisterEndpoints(router *gin.Engine, userController *controller.UserController) {

	auth := router.Group("/")

	auth.POST("/sign-up", userController.SingUp)
	auth.POST("/login", userController.Login)
	auth.POST("/logout", userController.Logout)

}
