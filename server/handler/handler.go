package handler

import (
	"books/server/db/dbservice"
	"fmt"
	"net/http"
)

//AuthorHandler ...
type AuthorHandler struct {
	authorRepo dbservice.AuthorRepo
}

func (h AuthorHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello golang")
	ctx := r.Context()
	authorRepo := dbservice.NewAuthorRepo()
	authorRepo.GetAllAuthors(ctx)
}

var _ http.Handler = (*AuthorHandler)(nil)

//Run is a http handler function
func Run() {
	http.HandleFunc("/", AuthorHandler{}.ServeHTTP)
	http.ListenAndServe(":8081", nil)
}
