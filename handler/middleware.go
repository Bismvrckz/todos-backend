package handler

import (
	"github.com/labstack/echo/v4"
)

func AppMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(ctx echo.Context) (err error) {
		return next(ctx)
	}
}
