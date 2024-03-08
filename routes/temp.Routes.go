package routes

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/simple-rest-API/handlers"
	Middlewares "github.com/simple-rest-API/middlewares"
)

func RouteTask(e *echo.Echo, db handlers.DB) {

	h := handlers.NewTask(db)
	tasks := e.Group("/API/v1/todo/tasks")
	tasks.Use(middleware.Recover(), Middlewares.Log)

	tasks.POST("/newTasks", h.Create)
	tasks.GET("", h.GetAll)
	tasks.PUT(":id", h.Update)
	tasks.DELETE(":id", h.Delete)
	// get By id
	tasks.GET("task:id", h.GetById)
}

func UpdateTask(e *echo.Echo, db handlers.DB) {
	h := handlers.NewTask(db)

	e.PUT("/:id", h.Update)
}

func DeleteTask(e *echo.Echo, db handlers.DB) {
	h := handlers.NewTask(db)

	e.DELETE("/:id", h.Update)
}

func GetByIdTask(e *echo.Echo, db handlers.DB) {
	h := handlers.NewTask(db)

	e.PUT("/:id", h.Update)
}
