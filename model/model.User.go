package model

type User struct {
	ID           int    `json:"id"`
	Name         string `json:"name"`
	Password     int    `json:"password"`
	Email        string `json:"email"`
	TasksCreated Tasks  `json:"currentTasks"`
}

type Users []User
