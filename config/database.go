package config

import (
	"os"
	"todo-app/helper"

	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func DatabaseConnection() *gorm.DB {

	// TODO : Get these variables from env variables

	err := godotenv.Load(".env")
	helper.ErrorPanic(err)

	connectionString := os.Getenv("DB_CONNECTION_STRING")
	dsn := connectionString

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	helper.ErrorPanic(err)

	return db
}
