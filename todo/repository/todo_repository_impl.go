package repository

import (
	"github.com/batuhancaam/todo-app/model"

	"github.com/batuhancaam/todo-app/helper"
	"github.com/batuhancaam/todo-app/todo/data/request"
	"gorm.io/gorm"
)

// TODO: Refactor
type TodoRepositoryImpl struct {
	DB *gorm.DB
}

func NewTodoRepositoryImpl(DB *gorm.DB) TodoRepository {
	return &TodoRepositoryImpl{DB: DB}
}

// Save implements TodoRepository.
func (t *TodoRepositoryImpl) Save(todo model.Todo, user *model.User) {
	err := t.DB.Model(&user).Association("Todos").Append(&todo)
	helper.ErrorPanic(err)
}

// Update implements TodoRepository.
func (t *TodoRepositoryImpl) Update(todo model.Todo) {
	var updateTodo = request.UpdateTodoRequest{
		ID:        todo.ID,
		Task:      todo.Task,
		StartTime: todo.StartTime,
		EndTime:   todo.EndTime,
		Completed: todo.Completed,
	}
	result := t.DB.Model(&todo).Updates(updateTodo)
	helper.ErrorPanic(result.Error)
}

// Delete implements TodoRepository.
func (t *TodoRepositoryImpl) Delete(todoId uint, user *model.User) {
	todo, err := t.FindByID(todoId, user)
	helper.ErrorPanic(err)
	t.DB.Model(&user).Association("Todos").Delete(todo)
}

// FindByID implements TodoRepository.
func (t *TodoRepositoryImpl) FindByID(todoId uint, user *model.User) (todo model.Todo, err error) {
	var todoModel model.Todo

	err = t.DB.Model(&user).Where("id=?", todoId).Association("Todos").Find(&todoModel)

	if err != nil {
		return todoModel, err
	}

	return todoModel, nil
}

// FindAll implements TodoRepository.
func (t *TodoRepositoryImpl) FindAll(user *model.User) []model.Todo {
	var todos []model.Todo
	err := t.DB.Model(&user).Association("Todos").Find(&todos)
	helper.ErrorPanic(err)
	return todos
}
