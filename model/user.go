package model

import "gorm.io/gorm"

type User struct {
	gorm.Model
	ID       uint `gorm:"primaryKey"`
	Email    string
	Password string
	Todos    []Todo
}
