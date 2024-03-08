package db

import "github.com/simple-rest-API/model"

type db struct {
	Tasks TaskRecord
	Users UserRecord
}

func NewDBInstance(TaskMap map[int]model.Task, UserMap map[int]model.User) db {
	taskDbInstance := TaskRecord{
		RecordID: 0,
		Tasks:    TaskMap,
	}

	UserDbInstance := UserRecord{
		UserRecordId: 0,
		Users:        UserMap,
		UsersTasks:   TaskMap,
	}

	DB := db{
		Tasks: taskDbInstance,
		Users: UserDbInstance,
	}
	return DB
}
