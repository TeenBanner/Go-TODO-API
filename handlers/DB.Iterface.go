package handlers

import "github.com/simple-rest-API/model"

type DB interface {
	// user methods
	GetAllUsers() (model.Users, error)
	GetUserById(ID int) (model.User, error)

	CreateUser(User *model.User) error
	UpdateUser(ID int, User *model.User) error
	DeleteUser(ID int) error

	// UserActionsOnTasks
	UserCreateTask(t *model.Task) error
	UserGetAllTasks() (model.Tasks, error)
	UserUpdateTask(ID int, task *model.Task) error
	UserDeleteTask(ID int) error
}
