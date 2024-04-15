package service

import (
	"github.com/batuhancaam/todo-app/todo/data/request"
	"github.com/batuhancaam/todo-app/todo/data/response"
)

type TodoService interface {
	Create(todo request.CreateTodoRequest)
	Update(todo request.UpdateTodoRequest)
	Delete(todoId uint)
	FindByID(todoId uint) response.TodoResponse
	FindAll() []response.TodoResponse
}
