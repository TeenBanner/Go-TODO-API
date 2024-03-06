package handlers

import (
	"github.com/simple-rest-API/model"
)

type DB interface {
	CreateTask(task *model.Task) error
	GetAllTasks() (model.Tasks, error)
	GetTaskByID(id int) (model.Task, error)
	UpdateTask(ID int, title, content string) (model.Task, error)
	DeleteTask(ID int) error
}
