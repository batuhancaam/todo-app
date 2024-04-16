package server

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/batuhancaam/todo-app/helper"
	tdcont "github.com/batuhancaam/todo-app/todo/http/controller"
	tdrouter "github.com/batuhancaam/todo-app/todo/http/router"
	tdrepo "github.com/batuhancaam/todo-app/todo/repository"
	tdservice "github.com/batuhancaam/todo-app/todo/service"

	usrcont "github.com/batuhancaam/todo-app/user/http/controller"
	usrrouter "github.com/batuhancaam/todo-app/user/http/router"
	usrrepo "github.com/batuhancaam/todo-app/user/repository"
	usrservice "github.com/batuhancaam/todo-app/user/service"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator"
	"github.com/joho/godotenv"
	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Server struct {
	httpServer *http.Server

	tdController  tdcont.TodoController
	usrController usrcont.UserController
}

func NewServer() *Server {

	validate := validator.New()
	db := initDB()

	// Todo process injections
	tdRepo := tdrepo.NewTodoRepositoryImpl(db, viper.GetString("mysql.todo_table"))
	tdService := tdservice.NewTodoServiceImpl(tdRepo, validate)

	// User process injections
	usrRepo := usrrepo.NewUserRepositoryImpl(db, viper.GetString("mysql.user_table"))
	usrService := usrservice.NewUserServiceImpl(usrRepo, validate)

	return &Server{
		tdController:  *tdcont.NewTodoController(tdService),
		usrController: *usrcont.NewUserController(usrService),
	}

}

func (s *Server) RunServer(port string) error {
	router := gin.Default()
	router.Use(
		gin.Recovery(),
		gin.Logger(),
	)

	usrrouter.RegisterEndpoints(router, &s.usrController)
	api := router.Group("/api")
	tdrouter.RegisterEndpoints(api, &s.tdController)

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
