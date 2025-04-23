package main

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB

func InitDB() {
	var err error
	dsn := "root:password@tcp(localhost:3306)/employee_db?parseTime=true"
	db, err = sql.Open("mysql", dsn)
	if err != nil {
		panic(err)
	}
}
