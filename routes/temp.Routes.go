package routes

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/simple-rest-API/handlers"
	Middlewares "github.com/simple-rest-API/middlewares"
)

func RouteTask(e *echo.Echo, db handlers.DB) {

	h := handlers.NewTask(db)
	tasks := e.Group("/API/v1/todo")
	tasks.Use(middleware.Recover(), Middlewares.Log)

	tasks.POST("/newTasks", h.Create)
	tasks.GET("/Tasks", h.GetAll)

}
