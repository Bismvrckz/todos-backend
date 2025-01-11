package databases

import (
	"database/sql"
	"task-api/config"
)

type TaskDB struct {
	TaskID     sql.NullInt64  `db:"task_id"`
	TaskName   sql.NullString `db:"task_name"`
	TaskDesc   sql.NullString `db:"task_description"`
	TaskStatus sql.NullString `db:"task_status"`
}

func (tkbaiDbImpl *AppDbImplement) CreateTask(data TaskDB) (err error) {
	query := "INSERT INTO todos.task (task_name,task_description, task_status) VALUES (?,?,?)"
	_, err = tkbaiDbImpl.ConnectTkbaiDB.Exec(query, data.TaskName, data.TaskDesc, data.TaskStatus)
	if err != nil {
		config.LogErr(err, "Query Error")
		return err
	}

	return err
}

func (tkbaiDbImpl *AppDbImplement) GetTasks() (result []TaskDB, err error) {
	err = tkbaiDbImpl.ConnectTkbaiDB.Select(&result, "SELECT * FROM todos.task")
	if err != nil {
		config.LogErr(err, "Query Error")
		return result, err
	}

	return result, err
}

func (tkbaiDbImpl *AppDbImplement) UpdateTask(data TaskDB) (rowsAffected int64, err error) {
	query := "UPDATE todos.task SET task_name = ?, task_description = ?, task_status = ? WHERE task_id = ?"
	resQuery, err := tkbaiDbImpl.ConnectTkbaiDB.Exec(query, data.TaskName, data.TaskDesc, data.TaskStatus, data.TaskID)
	if err != nil {
		config.LogErr(err, "Query Error")
		return rowsAffected, err
	}

	rowsAffected, err = resQuery.RowsAffected()
	if err != nil {
		return rowsAffected, err
	}

	return rowsAffected, err
}

func (tkbaiDbImpl *AppDbImplement) DeleteTask(taskID int64) (rowsAffected int64, err error) {
	query := "DELETE FROM todos.task WHERE task_id = ?"
	resQuery, err := tkbaiDbImpl.ConnectTkbaiDB.Exec(query, taskID)
	if err != nil {
		config.LogErr(err, "Query Error")
		return rowsAffected, err
	}

	rowsAffected, err = resQuery.RowsAffected()
	if err != nil {
		return rowsAffected, err
	}

	return rowsAffected, err
}
