package config

import (
	"database/sql"
	"time"

	"github.com/FRFebi/go-movie/helper"
)

func NewDB() *sql.DB {
	db, err := sql.Open("mysql", "root:1234567890@tcp(127.0.0.1:3306)/tht")
	helper.PanicIfError(err)

	db.SetMaxIdleConns(5)
	db.SetMaxOpenConns(20)
	db.SetConnMaxIdleTime(60 * time.Minute)
	db.SetConnMaxLifetime(10 * time.Minute)

	return db
}
