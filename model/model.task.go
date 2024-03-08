package model

type Task struct {
	Owner   User   `json:"owner"`
	Id      int    `json:"id"`
	Title   string `json:"title"`
	Content string `json:"content"`
}

type Tasks []Task
