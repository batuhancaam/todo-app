package repository

import (
	"github.com/batuhancaam/todo-app/helper"
	"github.com/batuhancaam/todo-app/model"
	"gorm.io/gorm"
)

type UserRepositoryImpl struct {
	DB *gorm.DB
}

func NewUserRepositoryImpl(DB *gorm.DB, tableName string) UserRepository {
	DB.Table(tableName).AutoMigrate(&model.User{})
	return &UserRepositoryImpl{
		DB: DB,
	}
}

// CreateUser implements UserRepository.
func (u *UserRepositoryImpl) CreateUser(user model.User) {
	result := u.DB.Create(&user)
	helper.ErrorPanic(result.Error)
}

// GetUser implements UserRepository.
func (u *UserRepositoryImpl) GetUser(email string) (user *model.User, err error) {
	var userRes *model.User

	result := u.DB.Where("email = ?", email).First(&userRes)

	if result.Error != nil {
		return nil, err
	}

	return userRes, nil

}
