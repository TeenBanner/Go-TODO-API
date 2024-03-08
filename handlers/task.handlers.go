package handlers

import (
	"log"
	"net/http"
	"strconv"

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
func (t *Task) Update(c echo.Context) error {
	IDstr := c.QueryParam("id")
	log.Println(IDstr)

	ID, err := strconv.Atoi(IDstr)
	if err != nil {
		response := newResponse(Error, "Id no valido", IDstr)
		return c.JSON(http.StatusBadRequest, response)
	}

	data := model.Task{}
	err = c.Bind(&data)

	if err != nil {
		response := newResponse(Error, "Peticion mal structurada", err)
		return c.JSON(http.StatusBadRequest, response)
	}

	err = t.db.UpdateTask(ID, &data)

	if err != nil {
		response := newResponse(Error, "Error al actualizar la tarea", err)
		return c.JSON(http.StatusInternalServerError, response)
	}

	response := newResponse(Message, "tarea Actualizada correctamente", data)
	return c.JSON(http.StatusOK, response)
}

func (t *Task) Delete(c echo.Context) error {
	IDstr := c.QueryParam("id")

	ID, err := strconv.Atoi(IDstr)

	if err != nil {
		response := newResponse(Error, "ID no valido", nil)
		return c.JSON(http.StatusOK, response)
	}

	err = t.db.DeleteTask(ID)

	if err != nil {
		response := newResponse(Error, "No se pudo eliminar la tarea", err)
		return c.JSON(http.StatusInternalServerError, response)
	}

	response := newResponse(Message, "Tarea Eliminda", nil)
	return c.JSON(http.StatusOK, response)
}

func (t *Task) GetById(c echo.Context) error {
	IDstr := c.QueryParam("id")

	ID, err := strconv.Atoi(IDstr)

	if err != nil {
		response := newResponse(Error, "ID no valido", nil)
		return c.JSON(http.StatusBadRequest, response)
	}

	data, err := t.db.GetTaskByID(ID)

	if err != nil {
		response := newResponse(Error, "No se pudo borrar la tarea", err)
		return c.JSON(http.StatusInternalServerError, response)
	}

	response := newResponse(Message, "Tarea elimindada satisfactoriamente", data)
	return c.JSON(http.StatusOK, response)
}
