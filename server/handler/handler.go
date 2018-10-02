package handler

import (
	"books/server/db/dbservice"
	"fmt"
	"net/http"
)

//GeneralHandler ...
type GeneralHandler struct {
	generalRepo *dbservice.GeneralRepo
}

//ServeHTTP ...
func (gh GeneralHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello golang")
	ctx := r.Context()
	gh.generalRepo.GetAll(ctx)
}

var _ http.Handler = (*GeneralHandler)(nil)

//Run is a http handler function
func Run(generalRepo *dbservice.GeneralRepo) {
	http.HandleFunc("/", GeneralHandler{generalRepo: generalRepo}.ServeHTTP)
	http.ListenAndServe(":8081", nil)
}
