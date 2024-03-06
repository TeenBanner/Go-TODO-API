package model

type Task struct {
	Id      int    `json:"id"`
	Title   string `json:"title"`
	Content string `json:"Content"`
}

type Tasks []Task
