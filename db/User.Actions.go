package db

import "github.com/simple-rest-API/model"

type UserActions struct {
	TaskRecord
	UserRecord
}

func (UA *UserActions) UserCreateTask(t *model.Task) error {
	err := UA.CreateTask(t)

	if err != nil {
		return err
	}

	return nil
}

func (UA *UserActions) UserGetAllTasks() (model.Tasks, error) {
	tasks, err := UA.GetAllTasks()

	if err != nil {
		return tasks, err
	}

	return tasks, nil
}

func (UA *UserActions) UserUpdateTask(ID int, task *model.Task) error {
	User := UA.UpdateTask(ID, task)

	if User != nil {
		return User
	}

	return nil
}

func (UA *UserActions) UserDeleteTask(ID int) error {
	User := UA.DeleteTask(ID)

	if User != nil {
		return User
	}

	return nil
}
