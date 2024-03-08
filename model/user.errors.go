package model

import "errors"

var (
	ErrUserDoNotExist  = errors.New("El usuario No existe")
	ErrUserCanNotBeNil = errors.New("El Usuario puede ser nulo")
)
