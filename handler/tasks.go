package handler

import (
	"database/sql"
	"github.com/labstack/echo/v4"
	"net/http"
	"task-api/databases"
	"task-api/models"
)

func CreateTask(ctx echo.Context) (err error) {
	var payload models.CreateTask
	err = ctx.Bind(&payload)
	if err != nil {
		return err
	}

	err = databases.DbInterface.CreateTask(databases.StudentData{
		TaskName: sql.NullString{
			String: payload.TaskName,
			Valid:  true,
		},
		TaskDesc: sql.NullString{
			String: payload.TaskDesc,
			Valid:  true,
		},
	})

	if err != nil {
		return err
	}

	return ctx.JSON(http.StatusOK, map[string]string{
		"status": "success",
	})
}

func GetAllTask(ctx echo.Context) (err error) {
	return
}
