package handler

import (
	"database/sql"
	"github.com/labstack/echo/v4"
	"net/http"
	"task-api/databases"
	"task-api/models"
)

func CreateTask(ctx echo.Context) (err error) {
	var payload models.TaskJson
	err = ctx.Bind(&payload)
	if err != nil {
		return err
	}

	err = databases.DbInterface.CreateTask(databases.TaskDB{
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
	result, err := databases.DbInterface.GetTasks()
	if err != nil {
		return err
	}

	var tasks []models.TaskJson
	for _, each := range result {
		tasks = append(tasks, models.TaskJson{
			TaskID:   each.TaskID.Int64,
			TaskName: each.TaskName.String,
			TaskDesc: each.TaskDesc.String,
		})
	}

	return ctx.JSON(http.StatusOK, tasks)
}

func UpdateTask(ctx echo.Context) (err error) {
	var payload models.TaskJson
	err = ctx.Bind(&payload)
	if err != nil {
		return err
	}

	rowsAffected, err := databases.DbInterface.UpdateTask(databases.TaskDB{
		TaskName: sql.NullString{
			String: payload.TaskName,
			Valid:  true,
		},
		TaskDesc: sql.NullString{
			String: payload.TaskDesc,
			Valid:  true,
		},
		TaskID: sql.NullInt64{
			Int64: payload.TaskID,
			Valid: true,
		},
	})

	if rowsAffected == 0 {
		return ctx.JSON(http.StatusNotFound, map[string]string{
			"status":  "failed",
			"message": "task not found",
		})
	}

	if err != nil {
		return err
	}

	return ctx.JSON(http.StatusOK, map[string]string{
		"status": "success",
	})
}

func DeleteTask(ctx echo.Context) (err error) {
	var payload models.TaskJson
	err = ctx.Bind(&payload)
	if err != nil {
		return err
	}

	rowsAffected, err := databases.DbInterface.DeleteTask(payload.TaskID)

	if rowsAffected == 0 {
		return ctx.JSON(http.StatusNotFound, map[string]string{
			"status":  "failed",
			"message": "task not found",
		})
	}

	if err != nil {
		return err
	}

	return ctx.JSON(http.StatusOK, map[string]string{
		"status": "success",
	})
}
