package app

import (
	"database/sql"
	"ridloxyanuar/go-rest-api/helper"
)

func NewDB() *sql.DB {
	db, err := sql.Open("mysql", "root:@tcp(localhost:3306)/go-rest-api-db")
	helper.PanicIfError(err)

	db.SetMaxIdleConns(5)
	db.SetMaxOpenConns(20)
	db.SetConnMaxLifetime(60 * 60 * 10)

	return db
}
