package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/simple-rest-API/middlewares"
)

func HelloWorld(e *echo.Echo) {
	e.Use(middlewares.Log)
	e.GET("/API/v1/HelloWorld", HelloWorldHandler)
}

func HelloWorldHandler(c echo.Context) error {
	c.JSON(http.StatusOK, map[string]string{"Message": "HelloWorld"})
	return nil
}
