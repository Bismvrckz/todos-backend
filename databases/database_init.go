package databases

import (
	"task-api/config"
)

func ConnectAppDatabase() (err error) {
	dbInstance, err := config.DbConnection()
	if err != nil {
		return err
	}

	err = dbInstance.Ping()
	if err != nil {
		return err
	}

	DbInterface = &AppDbImplement{ConnectTkbaiDB: dbInstance}

	return err
}
