package controller

import (
	"net/http"
	"strconv"
	"todo-app/data/request"
	"todo-app/data/response"
	"todo-app/helper"
	"todo-app/service"

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
	webResponse := response.Response{
		Code:   http.StatusOK,
		Status: "OK",
		Data:   nil,
	}
	ctx.Header("Content-Type", "application/json")
	ctx.JSON(http.StatusOK, webResponse)
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

	webResponse := response.Response{
		Code:   http.StatusOK,
		Status: "OK",
		Data:   nil,
	}
	ctx.Header("Content-Type", "application/json")
	ctx.JSON(http.StatusOK, webResponse)

}

func (controller *TodoController) Delete(ctx *gin.Context) {
	todoId := ctx.Param("todoId")
	id, err := strconv.ParseUint(todoId, 10, 32)
	helper.ErrorPanic(err)

	controller.todoService.Delete(uint(id))

	webResponse := response.Response{
		Code:   http.StatusOK,
		Status: "OK",
		Data:   nil,
	}
	ctx.Header("Content-Type", "application/json")
	ctx.JSON(http.StatusOK, webResponse)
}

func (controller *TodoController) FindByID(ctx *gin.Context) {
	todoId := ctx.Param("todoId")
	id, err := strconv.ParseUint(todoId, 10, 32)

	helper.ErrorPanic(err)

	todoResponse := controller.todoService.FindByID(uint(id))

	webResponse := response.Response{
		Code:   http.StatusOK,
		Status: "OK",
		Data:   todoResponse,
	}
	ctx.Header("Content-Type", "application/json")
	ctx.JSON(http.StatusOK, webResponse)
}

func (controller *TodoController) FindAll(ctx *gin.Context) {
	todosResponse := controller.todoService.FindAll()

	webResponse := response.Response{
		Code:   http.StatusOK,
		Status: "OK",
		Data:   todosResponse,
	}
	ctx.Header("Content-Type", "application/json")
	ctx.JSON(http.StatusOK, webResponse)

}
