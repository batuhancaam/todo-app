package repository

import "github.com/batuhancaam/todo-app/model"

type TodoRepository interface {
	Save(todo model.Todo)
	Update(todo model.Todo)
	Delete(todoId uint)
	FindByID(todoId uint) (todo model.Todo, err error)
	FindAll() []model.Todo
}
