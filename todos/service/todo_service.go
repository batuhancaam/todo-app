package service

import (
	"todo-app/data/request"
	"todo-app/data/response"
)

type TodoService interface {
	Create(todo request.CreateTodoRequest)
	Update(todo request.UpdateTodoRequest)
	Delete(todoId uint)
	FindByID(todoId uint) response.TodoResponse
	FindAll() []response.TodoResponse
}
