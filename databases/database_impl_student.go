package databases

import (
	"database/sql"
	"task-api/config"
)

type StudentData struct {
	ID             sql.NullInt64  `json:"id" db:"id"`
	StudentAddress sql.NullString `json:"studentID" db:"student_address"`
	Name           sql.NullString `json:"name" db:"name"`
	StudentNumber  sql.NullString `json:"studentNumber" db:"student_number" `
	Major          sql.NullString `json:"major" db:"major"`
	InsertDate     sql.NullTime   `json:"insertDate" db:"insert_date"`
}

func (tkbaiDbImpl *AppDbImplement) CreateStudentData(data StudentData) (err error) {
	query := "INSERT INTO tkbai_data (student_address, name, student_number, major) VALUES (?,?,?,?)"
	_, err = tkbaiDbImpl.ConnectTkbaiDB.Exec(query, data.StudentAddress, data.Name, data.StudentNumber, data.Major)
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
