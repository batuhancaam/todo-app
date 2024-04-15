package repository

import (
	"todo-app/helper"
	"todo-app/model"

	"gorm.io/gorm"
)

type TodoRepositoryImpl struct {
	DB *gorm.DB
}

func NewTodoRepositoryImpl(DB *gorm.DB) TodoRepository {
	return &TodoRepositoryImpl{DB: DB}
}

// Delete implements TodoRepository.
func (t *TodoRepositoryImpl) Delete(todoId uint) {
	var todo model.Todo
	result := t.DB.Where("id =?", todoId).Delete(&todo)
	helper.ErrorPanic(result.Error)
}

// FindAll implements TodoRepository.
func (t *TodoRepositoryImpl) FindAll() []model.Todo {
	var todos []model.Todo
	result := t.DB.Find(&todos)
	helper.ErrorPanic(result.Error)
	return todos
}

// FindByID implements TodoRepository.
func (t *TodoRepositoryImpl) FindByID(todoId uint) (todo model.Todo, err error) {
	var todoModel model.Todo

	result := t.DB.First(&todoModel, todoId)

	if result.Error != nil {
		return todo, result.Error
	}

	return todoModel, nil
}

// Save implements TodoRepository.
func (t *TodoRepositoryImpl) Save(todo model.Todo) {
	result := t.DB.Create(&todo)
	helper.ErrorPanic(result.Error)
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
