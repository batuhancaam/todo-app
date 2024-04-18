package service

import (
	"github.com/batuhancaam/todo-app/model"
	"github.com/batuhancaam/todo-app/todo/data/request"
	"github.com/batuhancaam/todo-app/todo/data/response"
)

type TodoService interface {
	Create(todo request.CreateTodoRequest, user *model.User)
	Update(todo request.UpdateTodoRequest, user *model.User)
	Delete(todoId uint, user *model.User)
	FindByID(todoId uint, user *model.User) response.TodoResponse
	FindAll(user *model.User) []response.TodoResponse
}
