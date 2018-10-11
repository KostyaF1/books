package bookHandlers

import (
	"books/server/handler"
	"books/server/service"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

type getBookByID struct {
	getBook service.Book
}

func NewGetBookByID() *getBookByID {
	return new(getBookByID)
}

func (g *getBookByID) Inject(getBook service.Book) {
	g.getBook = getBook
}

var _ http.Handler = (*deleteBook)(nil)
var _ handler.Router = (*deleteBook)(nil)

func (g *getBookByID) Path() (path string) {
	return "/get_book_id/{id:[0-9]+}"
}

func (g *getBookByID) Method() (method string) {
	return http.MethodGet
}

func (g *getBookByID) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	vars := mux.Vars(r)
	id := vars["id"]
	id64, err := strconv.ParseInt(id, 10, 64)

	if err != nil {
		fmt.Fprintf(w, "error: %s", err.Error())
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	response := g.getBook.GetBookByID(ctx, id64)

	if err := json.NewEncoder(w).Encode(response); err != nil {
		fmt.Fprintf(w, "error: %s", err.Error())
		w.WriteHeader(http.StatusBadRequest)
		return
	}
}
