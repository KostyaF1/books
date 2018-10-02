package main

import (
	"books/server/db"
	"books/server/db/dbservice"
	"books/server/handler"
	"books/server/routing"
	"books/server/service"
	"log"
)

func main() {
	dataConfig := db.NewDataConfig("postgres", "postgres://postgres:1111@localhost/books")

	dbConn, err := db.NewConn(dataConfig)
	defer dbConn.DBConn.Close()
	if err != nil {
		log.Println("DB Open Error" + err.Error())
	}
	generalRepo := dbservice.NewGeneralRepo(dbConn)

	service := service.NewGeneralService(generalRepo)
	routerIn := routing.NewRouterIn(service)
	newHandler := handler.NewGeneralHandler(routerIn)
	handler.Run(*newHandler)
}
