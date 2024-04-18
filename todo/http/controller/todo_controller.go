package controller

import (
	"net/http"
	"strconv"

	"github.com/batuhancaam/todo-app/helper"
	"github.com/batuhancaam/todo-app/model"
	"github.com/batuhancaam/todo-app/todo/data/request"
	"github.com/batuhancaam/todo-app/todo/service"
	"github.com/gin-gonic/gin"
)

type TodoController struct {
	todoService service.TodoService
}

func NewTodoController(service service.TodoService) *TodoController {
	return &TodoController{
		todoService: service,
	}
}

func (controller *TodoController) Create(ctx *gin.Context) {
	createTodoRequest := request.CreateTodoRequest{}

	err := ctx.ShouldBindJSON(&createTodoRequest)
	user := ctx.MustGet("user").(*model.User)
	helper.ErrorPanic(err)

	controller.todoService.Create(createTodoRequest, user)

	ctx.Header("Content-Type", "application/json")
	ctx.Status(http.StatusCreated)
}

func (controller *TodoController) Update(ctx *gin.Context) {
	updateTodoRequest := request.UpdateTodoRequest{}

	err := ctx.ShouldBindJSON(&updateTodoRequest)
	helper.ErrorPanic(err)
	todoId := ctx.Param("todoId")
	user := ctx.MustGet("user").(*model.User)
	id, err := strconv.ParseUint(todoId, 10, 32)
	helper.ErrorPanic(err)
	updateTodoRequest.ID = uint(id)
	controller.todoService.Update(updateTodoRequest, user)

	ctx.Header("Content-Type", "application/json")
	ctx.Status(http.StatusOK)

}

func (controller *TodoController) Delete(ctx *gin.Context) {
	todoId := ctx.Param("todoId")
	id, err := strconv.ParseUint(todoId, 10, 32)
	helper.ErrorPanic(err)
	user := ctx.MustGet("user").(*model.User)

	controller.todoService.Delete(uint(id), user)

	ctx.Header("Content-Type", "application/json")
	ctx.Status(http.StatusOK)
}

func (controller *TodoController) FindByID(ctx *gin.Context) {
	todoId := ctx.Param("todoId")
	id, err := strconv.ParseUint(todoId, 10, 32)

	helper.ErrorPanic(err)
	user := ctx.MustGet("user").(*model.User)

	todoResponse := controller.todoService.FindByID(uint(id), user)

	ctx.Header("Content-Type", "application/json")
	ctx.JSON(http.StatusOK, todoResponse)
}

func (controller *TodoController) FindAll(ctx *gin.Context) {
	user := ctx.MustGet("user").(*model.User)

	todosResponse := controller.todoService.FindAll(user)

	ctx.Header("Content-Type", "application/json")
	ctx.JSON(http.StatusOK, todosResponse)

}
