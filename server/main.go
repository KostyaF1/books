package main

import (
	"books/server/db"
	"books/server/db/dbservice"
	"books/server/handler"
	"log"
)

func main() {
	dataConfig := db.NewDataConfig("postgres", "postgres://postgres:1111@localhost/books")

	dbConn, err := dbservice.NewConn(dataConfig)
	defer dbConn.DBConn.Close()
	if err != nil {
		log.Println("DB Open Error" + err.Error())
	}

	generalRepo := dbservice.NewGeneralRepo(dbConn)

	handler.Run(generalRepo)
}
