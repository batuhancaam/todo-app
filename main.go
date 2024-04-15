package main

import (
	"net/http"
	"todo-app/config"
	"todo-app/helper"
	"todo-app/model"

	"github.com/go-playground/validator/v10"
	"github.com/rs/zerolog/log"
)

func main() {

	log.Info().Msg("Started server!")
	// Database Connection
	db := config.DatabaseConnection()
	db.Table("todos").AutoMigrate(&model.Todo{})

	validate := validator.New()

	// Repository Initialize
	todoRepository := repository.NewTodoRepositoryImpl(db)

	// Service Initialize
	todoService := service.NewTodoServiceImpl(todoRepository, validate)

	//Controller Initialize
	todoController := controller.NewTodoController(todoService)

	//Router Initialize
	routes := router.NewRouter(todoController)

	server := http.Server{
		Addr:    ":9090",
		Handler: routes,
	}

	err := server.ListenAndServe()
	helper.ErrorPanic(err)
}
