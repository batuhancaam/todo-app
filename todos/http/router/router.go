package router

import (
	"net/http"
	"todo-app/controller"

	"github.com/gin-gonic/gin"
)

func NewRouter(todoController *controller.TodoController) *gin.Engine {

	router := gin.Default()

	router.GET("", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, "welcome home")
	})

	todosRouter := router.Group("/api/todos")
	todosRouter.GET("", todoController.FindAll)
	todosRouter.GET("/:todoId", todoController.FindByID)
	todosRouter.POST("", todoController.Create)
	todosRouter.PATCH("/:todoId", todoController.Update)
	todosRouter.DELETE("/:todoId", todoController.Delete)

	return router
}
