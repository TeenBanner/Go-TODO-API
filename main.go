package main

import (
	"github.com/simple-rest-API/model"
	"log"

	"github.com/labstack/echo/v4"
)

func makeDBinstances() (map[int]model.Task, map[int]model.User) {
	TaskMapInstance := make(map[int]model.Task)
	UserMapInstance := make(map[int]model.User)

	return TaskMapInstance, UserMapInstance
}

func main() {
	e := echo.New()
	/* No Descomentar hasta tener listas las rutas
	dbt, dbu := makeDBinstances()
	db := db.NewDBInstance(dbt, dbu)*/

	log.Println("Server Listening at http://127.0.0.1:1234/")
	log.Printf("Routes available: \n")
	log.Printf("Server availabe routes GET: http://127.0.0.1:1234/API/v1/todo/tasks \nPOST: http://127.0.0.1:1234/API/v1/todo/tasks/newTasks \nPUT: http://127.0.0.1:1234/API/v1/todo/tasks")

	err := e.Start(":1234")
	if err != nil {
		log.Fatal(err)
	}

}
