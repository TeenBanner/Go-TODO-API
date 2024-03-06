package main

import (
	"log"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()

	err := e.Start(":1234")
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Server Listening at http://127.0.0.1:1234/")
	log.Printf("Routes available: \n")
	log.Println("Server Listening at http://127.0.0.1:1234/API/v1/HelloWorld")
}
