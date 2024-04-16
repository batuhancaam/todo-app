package repository

import "github.com/batuhancaam/todo-app/model"

type UserRepository interface {
	CreateUser(user model.User)
	GetUser(email string) (user *model.User, err error)
}
