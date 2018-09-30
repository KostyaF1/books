package handler

import (
	"books/server/db/dbservice"
	"fmt"
	"log"
	"net/http"
)

func myHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello golang")
	auth := dbservice.AuthorDB{}
	auth.ASelect()

}

//Handler is a http handler function
func Handler() {
	go http.HandleFunc("/", myHandler)
	http.ListenAndServe(":8080", nil)
	log.Println("Server listening")
}
