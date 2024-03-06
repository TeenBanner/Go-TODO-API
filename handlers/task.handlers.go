package handlers

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/simple-rest-API/model"
)

type Task struct {
	db DB
}

type response struct {
	Status  string      `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func newResponse(status, message string, data interface{}) response {
	return response{
		Status:  status,
		Message: message,
		Data:    data,
	}
}

var (
	Message string = "OK"
	Error   string = "Error"
)

func NewTask(db DB) Task {
	return Task{db}
}

func (t *Task) Create(c echo.Context) error {
	data := model.Task{}
	err := c.Bind(&data)
	if err != nil {
		response := newResponse(Error, "peticion Mal estructurada", data)
		return c.JSON(http.StatusBadRequest, response)
	}

	err = t.db.CreateTask(&data)
	if err != nil {
		response := newResponse(Error, "Hubo un error al crear la persona ", err)
		return c.JSON(http.StatusInternalServerError, response)
	}

	response := newResponse(Message, "Persona Creada Correctamente", data)
	return c.JSON(http.StatusCreated, response)
}

func (t *Task) GetAll(c echo.Context) error {
	data, err := t.db.GetAllTasks()
	if err != nil {
		response := newResponse(Error, "1", data)
		return c.JSON(http.StatusOK, response)
	}
	response := newResponse(Message, "Tareas obtenidas correctamente", data)
	return c.JSON(http.StatusOK, response)
}
