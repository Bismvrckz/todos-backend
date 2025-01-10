package databases

import (
	"database/sql"
	"task-api/config"
)

type StudentData struct {
	TaskID   sql.NullString `json:"taskId" db:"task_id"`
	TaskName sql.NullString `json:"taskName" db:"task_name"`
	TaskDesc sql.NullString `json:"taskDesc" db:"task_description"`
}

func (tkbaiDbImpl *AppDbImplement) CreateTask(data StudentData) (err error) {
	query := "INSERT INTO todos.task (task_name,task_description) VALUES (?,?)"
	_, err = tkbaiDbImpl.ConnectTkbaiDB.Exec(query, data.TaskName, data.TaskDesc)
	if err != nil {
		config.LogErr(err, "Query Error")
		return err
	}

	return err
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

func (tkbaiDbImpl *AppDbImplement) ViewAllStudentData(start, length string) (result []StudentData, err error) {
	query := "SELECT * FROM tkbai_data LIMIT ? OFFSET ?"
	err = tkbaiDbImpl.ConnectTkbaiDB.Get(&result, query, length, start)
	if err != nil {
		config.LogErr(err, "QUERY ERROR")
		return result, err
	}

	config.LogTrc("SUCCESS")
	return result, err
}

func (tkbaiDbImpl *AppDbImplement) ViewStudentDataByNumberAndName(studentNumber, studentName string) (result StudentData, err error) {
	query := `SELECT * FROM tkbai_data WHERE student_number = ? AND name = ?`
	err = tkbaiDbImpl.ConnectTkbaiDB.Get(&result, query, studentNumber, studentName)
	if err != nil {
		config.LogErr(err, "QUERY ERROR")
		return result, err
	}

	config.LogTrc("SUCCESS")
	return result, err
}

func (tkbaiDbImpl *AppDbImplement) ViewStudentDataByIdOrName(credential string) (result StudentData, err error) {
	query := `SELECT * FROM tkbai_data WHERE student_number = ? OR name = ?`
	err = tkbaiDbImpl.ConnectTkbaiDB.Get(&result, query, credential, credential)
	if err != nil {
		config.LogErr(err, "QUERY ERROR")
		return result, err
	}

	config.LogTrc("SUCCESS")
	return result, err
}

func (tkbaiDbImpl *AppDbImplement) ViewStudentDataBulk() (result []StudentData, err error) {
	err = tkbaiDbImpl.ConnectTkbaiDB.Select(&result, "SELECT * FROM tkbai_data")
	if err != nil {
		config.LogErr(err, "Query Error")
		return result, err
	}

	return result, err
}
