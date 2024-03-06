package Middlewares

import (
	"log"
	"time"

	"github.com/labstack/echo/v4"
)

func Log(f echo.HandlerFunc) echo.HandlerFunc {
	startTime := time.Now()
	return func(c echo.Context) error {
		duration := time.Since(startTime)
		log.Printf("Peticion To: %s, HandlingTime: %q, Method: %q", c.Request().URL, duration, c.Request().Method)
		return f(c)
	}
}
