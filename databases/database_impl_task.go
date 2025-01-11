package databases

import (
	"database/sql"
	"task-api/config"
)

type TaskDB struct {
	TaskID   sql.NullInt64  `db:"task_id"`
	TaskName sql.NullString `db:"task_name"`
	TaskDesc sql.NullString `db:"task_description"`
}

func (tkbaiDbImpl *AppDbImplement) CreateTask(data TaskDB) (err error) {
	query := "INSERT INTO todos.task (task_name,task_description) VALUES (?,?)"
	_, err = tkbaiDbImpl.ConnectTkbaiDB.Exec(query, data.TaskName, data.TaskDesc)
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
	query := "UPDATE todos.task SET task_name = ?, task_description = ? WHERE task_id = ?"
	resQuery, err := tkbaiDbImpl.ConnectTkbaiDB.Exec(query, data.TaskName, data.TaskDesc, data.TaskID)
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

func (tkbaiDbImpl *AppDbImplement) CountAllStudentData() (result int64, err error) {
	query := "SELECT COUNT(*) AS total_rows FROM tkbai_data"
	err = tkbaiDbImpl.ConnectTkbaiDB.Get(&result, query)
	if err != nil {
		config.LogErr(err, "QUERY ERROR")
		return result, err
	}

	config.LogTrc("SUCCESS")
	return result, err
}

func (tkbaiDbImpl *AppDbImplement) DeleteALlStudentData() (err error) {
	query := "DELETE FROM tkbai_data "
	_, err = tkbaiDbImpl.ConnectTkbaiDB.Exec(query)
	if err != nil {
		config.LogErr(err, "Query Error")
		return err
	}

	return err
}

func (tkbaiDbImpl *AppDbImplement) DeleteStudentData(id string) (err error) {
	query := "DELETE FROM tkbai_data WHERE id = ?"
	_, err = tkbaiDbImpl.ConnectTkbaiDB.Exec(query, id)
	if err != nil {
		config.LogErr(err, "Query Error")
		return err
	}

	return err
}

func (tkbaiDbImpl *AppDbImplement) ViewAllStudentData(start, length string) (result []TaskDB, err error) {
	query := "SELECT * FROM tkbai_data LIMIT ? OFFSET ?"
	err = tkbaiDbImpl.ConnectTkbaiDB.Get(&result, query, length, start)
	if err != nil {
		config.LogErr(err, "QUERY ERROR")
		return result, err
	}

	config.LogTrc("SUCCESS")
	return result, err
}

func (tkbaiDbImpl *AppDbImplement) ViewStudentDataByNumberAndName(studentNumber, studentName string) (result TaskDB, err error) {
	query := `SELECT * FROM tkbai_data WHERE student_number = ? AND name = ?`
	err = tkbaiDbImpl.ConnectTkbaiDB.Get(&result, query, studentNumber, studentName)
	if err != nil {
		config.LogErr(err, "QUERY ERROR")
		return result, err
	}

	config.LogTrc("SUCCESS")
	return result, err
}

func (tkbaiDbImpl *AppDbImplement) ViewStudentDataByIdOrName(credential string) (result TaskDB, err error) {
	query := `SELECT * FROM tkbai_data WHERE student_number = ? OR name = ?`
	err = tkbaiDbImpl.ConnectTkbaiDB.Get(&result, query, credential, credential)
	if err != nil {
		config.LogErr(err, "QUERY ERROR")
		return result, err
	}

	config.LogTrc("SUCCESS")
	return result, err
}

func (tkbaiDbImpl *AppDbImplement) ViewStudentDataBulk() (result []TaskDB, err error) {
	err = tkbaiDbImpl.ConnectTkbaiDB.Select(&result, "SELECT * FROM tkbai_data")
	if err != nil {
		config.LogErr(err, "Query Error")
		return result, err
	}

	return result, err
}
