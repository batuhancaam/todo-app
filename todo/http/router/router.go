package router

import (
	"github.com/batuhancaam/todo-app/todo/http/controller"
	"github.com/gin-gonic/gin"
)

func RegisterEndpoints(router *gin.RouterGroup, todoController *controller.TodoController) {

	todos := router.Group("/todos")
	todos.GET("", todoController.FindAll)
	todos.GET("/:todoId", todoController.FindByID)
	todos.POST("", todoController.Create)
	todos.PATCH("/:todoId", todoController.Update)
	todos.DELETE("/:todoId", todoController.Delete)

}
