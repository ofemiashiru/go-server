package user

import (
	"acme/model"
)

type UserRepository interface {
	GetUsers() ([]model.User, error)
	GetUser(id int) ([]model.User, error)
	AddUser(user []model.User) (id int, err error)
	UpdateUser(id int, user []model.User) ([]model.User, error)
	DeleteUser(id int) error
	Close()
}
