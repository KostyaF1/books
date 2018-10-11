package bookHandlers

import (
	"books/server/handler"
	"books/server/service"
	"encoding/json"
	"fmt"
	"net/http"
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
	return "/get_book_id"
}

func (g *getBookByID) Method() (method string) {
	return http.MethodGet
}

func (g *getBookByID) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	response := g.getBook.GetBookByID(ctx)

	if err := json.NewEncoder(w).Encode(response); err != nil {
		fmt.Fprintf(w, "error: %s", err.Error())
		w.WriteHeader(http.StatusBadRequest)
		return
	}
}
