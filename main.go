package main

import (
	"books/server/db"
	"books/server/db/dbservice"
	"books/server/handler"
)

func main() {
	dataConfig := db.NewDataConfig("postgres", "postgres://postgres:books@localhost/books")
	dbConn := dbservice.OpenConnection(dataConfig)
	defer dbConn.Close()
	handler.Handler()
}
