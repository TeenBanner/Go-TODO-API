package model

import "errors"

var ErrtaskDoNotExist = errors.New("El id de la tarea O el usuario no existen")
var ErrCanotUpatetask = errors.New("No se pudo actualizar la tarea")
var ErrTaskCanNotBeNil = errors.New("La persona no puede ser nula")
