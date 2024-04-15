package controller

import (
	"net/http"
	"strconv"

	"github.com/batuhancaam/todo-app/helper"
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

	helper.ErrorPanic(err)

	controller.todoService.Create(createTodoRequest)

	ctx.Header("Content-Type", "application/json")
	ctx.Status(http.StatusCreated)
}

func (controller *TodoController) Update(ctx *gin.Context) {
	updateTodoRequest := request.UpdateTodoRequest{}

	err := ctx.ShouldBindJSON(&updateTodoRequest)
	helper.ErrorPanic(err)

	todoId := ctx.Param("todoId")
	id, err := strconv.ParseUint(todoId, 10, 32)
	helper.ErrorPanic(err)
	updateTodoRequest.ID = uint(id)
	controller.todoService.Update(updateTodoRequest)

	ctx.Header("Content-Type", "application/json")
	ctx.Status(http.StatusOK)

}

func (controller *TodoController) Delete(ctx *gin.Context) {
	todoId := ctx.Param("todoId")
	id, err := strconv.ParseUint(todoId, 10, 32)
	helper.ErrorPanic(err)

	controller.todoService.Delete(uint(id))

	ctx.Header("Content-Type", "application/json")
	ctx.Status(http.StatusOK)
}

func (controller *TodoController) FindByID(ctx *gin.Context) {
	todoId := ctx.Param("todoId")
	id, err := strconv.ParseUint(todoId, 10, 32)

	helper.ErrorPanic(err)

	todoResponse := controller.todoService.FindByID(uint(id))

	ctx.Header("Content-Type", "application/json")
	ctx.JSON(http.StatusOK, todoResponse)
}

func (controller *TodoController) FindAll(ctx *gin.Context) {
	todosResponse := controller.todoService.FindAll()

	ctx.Header("Content-Type", "application/json")
	ctx.JSON(http.StatusOK, todosResponse)

}
