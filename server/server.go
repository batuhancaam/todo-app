package server

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/batuhancaam/todo-app/helper"
	"github.com/batuhancaam/todo-app/todo/http/controller"
	todorouter "github.com/batuhancaam/todo-app/todo/http/router"
	"github.com/batuhancaam/todo-app/todo/repository"
	"github.com/batuhancaam/todo-app/todo/service"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator"
	"github.com/joho/godotenv"
	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Server struct {
	httpServer *http.Server

	tdController controller.TodoController
}

func NewServer() *Server {

	validate := validator.New()
	db := initDB()

	tdRepo := repository.NewTodoRepositoryImpl(db, viper.GetString("mysql.todo_table"))
	tdService := service.NewTodoServiceImpl(tdRepo, validate)

	return &Server{
		tdController: *controller.NewTodoController(tdService),
	}

}

func (s *Server) RunServer(port string) error {
	router := gin.Default()
	router.Use(
		gin.Recovery(),
		gin.Logger(),
	)
	api := router.Group("/api")
	todorouter.RegisterEndpoints(api, &s.tdController)

	s.httpServer = &http.Server{
		Addr:    ":" + viper.GetString("port"),
		Handler: router,
	}

	go func() {
		if err := s.httpServer.ListenAndServe(); err != nil {
			log.Fatalf("Failed to start server: %+v", err)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, os.Interrupt)

	<-quit

	ctx, shutdown := context.WithTimeout(context.Background(), 5*time.Second)
	defer shutdown()

	return s.httpServer.Shutdown(ctx)
}

func initDB() *gorm.DB {

	err := godotenv.Load(".env")
	helper.ErrorPanic(err)

	connectionString := os.Getenv("DB_CONNECTION_STRING")
	dsn := connectionString

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	helper.ErrorPanic(err)

	return db
}
