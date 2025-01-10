package config

import (
	"time"

	"github.com/jmoiron/sqlx"
)

var (
	dbMaxIdleCons = 10
	dbMaxCons     = 100
)

func DbConnection() (db *sqlx.DB, err error) {
	db, err = sqlx.Open("mysql", DbUrl)
	if err != nil {
		return db, err
	}

	err = db.Ping()
	if err != nil {
		return db, err
	}

	db.SetMaxOpenConns(dbMaxCons)
	db.SetMaxIdleConns(dbMaxIdleCons)
	db.SetConnMaxIdleTime(5 * time.Second)
	db.SetConnMaxLifetime(15 * time.Second)
	return db, err
}
