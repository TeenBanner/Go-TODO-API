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

	log.Println("Server Listening at http://127.0.0.1:1234/")
	log.Printf("Routes available: \n")
	log.Println("Server availabe routes http://127.0.0.1:1234/API/v1/todo/newtasks")

	err := e.Start(":1234")
	if err != nil {
		log.Fatal(err)
	}

}
