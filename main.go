package main

import (
	"log"

	"github.com/batuhancaam/todo-app/config"
	"github.com/batuhancaam/todo-app/server"
	"github.com/spf13/viper"
)

func main() {

	if err := config.Init(); err != nil {
		log.Fatalf("%s", err.Error())
	}

	server := server.NewServer()

	if err := server.RunServer(viper.GetString("port")); err != nil {
		log.Fatalf("%s", err.Error())
	}
}
