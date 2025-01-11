package handler

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"task-api/config"
)

func InitErrHandler(ein *config.Apps) {
	ein.Tkbai.HTTPErrorHandler = func(err error, ctx echo.Context) {
		config.Log.Debug().Msg(err.Error())

		ctx.JSON(http.StatusInternalServerError, map[string]string{
			"status":  "success",
			"message": err.Error(),
		})
	}
}
