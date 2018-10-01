package main

import (
	"books/server/db"
	"books/server/db/dbservice"
	"books/server/handler"
)

func main() {
	dataConfig := db.NewDataConfig("postgres", "postgres://postgres:1111@localhost/books")
	dbConn := dbservice.NewConn(dataConfig)
	defer dbConn.Close()
	//if err != nil {
	//	log.Println("DB Open Error" + err.Error())
	//}

	handler.Run()
}
