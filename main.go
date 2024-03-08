package main

import (
	"log"

	"github.com/labstack/echo/v4"
	"github.com/simple-rest-API/db"
	"github.com/simple-rest-API/routes"
)

func main() {
	e := echo.New()
	db := db.InitTaskRecords()

	routes.RouteTask(e, &db)
	routes.UpdateTask(e, &db)
	routes.DeleteTask(e, &db)
	routes.GetByIdTask(e, &db)

	log.Println("Server Listening at http://127.0.0.1:1234/")
	log.Printf("Routes available: \n")
	log.Printf("Server availabe routes GET: http://127.0.0.1:1234/API/v1/todo/tasks \nPOST: http://127.0.0.1:1234/API/v1/todo/tasks/newTasks \nPUT: http://127.0.0.1:1234/API/v1/todo/tasks")

	err := e.Start(":1234")
	if err != nil {
		log.Fatal(err)
	}

}
