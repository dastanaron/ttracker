package database

import (
	"database/sql"
	"fmt"
	"gui-mini-ttracker/helpers"
	"os"

	_ "github.com/mattn/go-sqlite3"
)

var connect *sql.DB

func GetConnection() *sql.DB {
	var err error

	cwd, err := os.Getwd()

	helpers.CheckError("Error receiving cwd", err)

	if connect == nil {
		connect, err = sql.Open("sqlite3", fmt.Sprintf("%s/data.db", cwd))
		helpers.CheckError("Error database connection", err)
		checkTables()
	}

	return connect
}

func checkTables() {
	createTaskTable()
	createServiceDataTable()
}

func createTaskTable() {
	sql := `CREATE TABLE IF NOT EXISTS "tasks" (
		"id"	INTEGER NOT NULL,
		"name"	TEXT,
		"startTime"	INTEGER,
		"duration"	INTEGER,
		"project"	TEXT,
		"endTime"	INTEGER,
		PRIMARY KEY("id" AUTOINCREMENT)
	);`

	GetConnection().Exec(sql)
}

func createServiceDataTable() {
	sql := `CREATE TABLE IF NOT EXISTS "service_data" (
		"name" TEXT NOT NULL UNIQUE,
		"data" TEXT NOT NULL,
		PRIMARY KEY("name")
	);`

	GetConnection().Exec(sql)
}
