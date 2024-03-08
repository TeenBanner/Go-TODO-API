package handlers

import (
	"github.com/labstack/echo/v4"
	"github.com/simple-rest-API/model"
	"net/http"
	"strconv"
)

type User struct {
	User DB
}

func NewUser(db DB) User {
	return User{db}
}

func (U *User) Create(c echo.Context) error {
	data := model.User{}
	err := c.Bind(&data)
	if err != nil {
		response := NewResponse(Error, "Peticion mal Estructurada", err)
		return c.JSON(http.StatusBadRequest, response)
	}

	err = U.User.CreateUser(&data)

	if err != nil {
		response := NewResponse(Error, "No se pudo Crear ala Usuario", err)
		return c.JSON(http.StatusInternalServerError, response)
	}

	response := NewResponse(Message, "Usuario Creada Correctamente", data)
	return c.JSON(http.StatusCreated, response)
}

func (U *User) GetAllUsers(c echo.Context) error {
	data, err := U.User.GetAllUsers()

	if err != nil {
		response := NewResponse(Error, "No se pudieron obtener al los usuarios", err)
		return c.JSON(http.StatusInternalServerError, response)
	}

	response := NewResponse(Message, "Usuarios obtenidos correctamente", data)
	return c.JSON(http.StatusOK, response)
}

func (U *User) GetUserByID(c echo.Context) error {
	Idstr := c.QueryParam("id")

	ID, err := strconv.Atoi(Idstr)

	if err != nil {
		response := NewResponse(Error, "Id no valido", nil)
		return c.JSON(http.StatusBadRequest, response)
	}

	data, err := U.User.GetUserById(ID)
	if err != nil {
		response := NewResponse(Error, "Error al obtener al usuario", err)
		return c.JSON(http.StatusBadRequest, response)
	}

	response := NewResponse(Message, "Usuario Obtenido satisfatorio", data)
	return c.JSON(http.StatusOK, response)
}

func (U *User) Update(c echo.Context) error {
	IDstr := c.QueryParam("id")

	ID, err := strconv.Atoi(IDstr)

	if err != nil {
		response := NewResponse(Error, "Id No valido", err)
		return c.JSON(http.StatusBadRequest, response)
	}

	data := model.User{}
	err = c.Bind(&data)

	err = U.User.UpdateUser(ID, &data)

	if err != nil {
		response := NewResponse(Error, "Error al actualizar el usuario", err)
		return c.JSON(http.StatusInternalServerError, response)
	}

	return nil
}

func (U *User) Delete(c echo.Context) error {
	IDstr := c.QueryParam("id")

	ID, err := strconv.Atoi(IDstr)
	if err != nil {
		response := NewResponse(Error, "Id no valido", err)
		return c.JSON(http.StatusBadRequest, response)
	}

	err = U.User.DeleteUser(ID)

	if err != nil {
		response := NewResponse(Error, "Hubo un error al eliminar al usuario", err)
		return c.JSON(http.StatusBadRequest, response)
	}

	response := NewResponse(Message, "persona Elminada Correctamente", nil)
	return c.JSON(http.StatusOK, response)
}
