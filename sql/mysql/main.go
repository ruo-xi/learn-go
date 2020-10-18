package main

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

const DriverName = "mysql"
const dataSourceName = "root:cao19981128@/study"

func start(driverName string, dataSourceName string) *sql.DB {
	db, err := sql.Open(driverName, dataSourceName)
	if err != nil {
		log.Fatal(err)
	}
	return db
}

func main() {
	db := start(DriverName, dataSourceName)
	defer db.Close()
}
