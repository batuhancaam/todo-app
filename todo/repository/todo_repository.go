package repository

import "github.com/batuhancaam/todo-app/model"

type TodoRepository interface {
	Save(todo model.Todo, user *model.User)
	Update(todo model.Todo)
	Delete(todoId uint, user *model.User)
	FindByID(todoId uint, user *model.User) (todo model.Todo, err error)
	FindAll(user *model.User) []model.Todo
}
