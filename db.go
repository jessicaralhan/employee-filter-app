package main

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB

func InitDB() {
	var err error
	dbHost := getEnv("DB_HOST", "localhost")
	dbUser := getEnv("DB_USER", "admin")
	dbPassword := getEnv("DB_PASSWORD", "admin123")
	dbName := getEnv("DB_NAME", "employee_db")

	dsn := fmt.Sprintf("%s:%s@tcp(%s:3306)/%s?parseTime=true",
		dbUser, dbPassword, dbHost, dbName)

	db, err = sql.Open("mysql", dsn)
	if err != nil {
		panic(err)
	}
}

func getEnv(key, defaultValue string) string {
	value := os.Getenv(key)
	if value == "" {
		return defaultValue
	}
	return value
}
