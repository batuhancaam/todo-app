package service

import (
	"github.com/batuhancaam/todo-app/helper"
	"github.com/batuhancaam/todo-app/model"
	"github.com/batuhancaam/todo-app/todo/data/request"
	"github.com/batuhancaam/todo-app/todo/data/response"
	"github.com/batuhancaam/todo-app/todo/repository"
	"github.com/go-playground/validator"
)

type TodoServiceImpl struct {
	TodoRepository repository.TodoRepository
	validate       *validator.Validate
}

// Create implements TodoService.
func (t *TodoServiceImpl) Create(todo request.CreateTodoRequest) {
	err := t.validate.Struct(todo)
	helper.ErrorPanic(err)

	todoModel := model.Todo{
		Task:      todo.Task,
		StartTime: todo.StartTime,
		EndTime:   todo.EndTime,
		Completed: todo.Completed,
	}

	t.TodoRepository.Save(todoModel)
}

// Delete implements TodoService.
func (t *TodoServiceImpl) Delete(todoId uint) {
	t.TodoRepository.Delete(todoId)
}

// FindAll implements TodoService.
func (t *TodoServiceImpl) FindAll() []response.TodoResponse {
	result := t.TodoRepository.FindAll()

	var todos []response.TodoResponse

	for _, value := range result {
		todo := response.TodoResponse{
			ID:        value.ID,
			Task:      value.Task,
			StartTime: value.StartTime,
			EndTime:   value.EndTime,
			Completed: value.Completed,
		}
		todos = append(todos, todo)
	}

	return todos
}

// FindByID implements TodoService.
func (t *TodoServiceImpl) FindByID(todoId uint) response.TodoResponse {
	todoData, err := t.TodoRepository.FindByID(todoId)

	helper.ErrorPanic(err)

	todoResponse := response.TodoResponse{
		ID:        todoData.ID,
		Task:      todoData.Task,
		StartTime: todoData.StartTime,
		EndTime:   todoData.EndTime,
		Completed: todoData.Completed,
	}

	return todoResponse

}

// Update implements TodoService.
func (t *TodoServiceImpl) Update(todo request.UpdateTodoRequest) {
	todoData, err := t.TodoRepository.FindByID(todo.ID)
	helper.ErrorPanic(err)

	todoData.Task = todo.Task
	todoData.StartTime = todo.StartTime
	todoData.EndTime = todo.EndTime
	todoData.Completed = todo.Completed

	t.TodoRepository.Update(todoData)
}

func NewTodoServiceImpl(todoRepository repository.TodoRepository, validate *validator.Validate) TodoService {
	return &TodoServiceImpl{
		TodoRepository: todoRepository,
		validate:       validate,
	}
}
