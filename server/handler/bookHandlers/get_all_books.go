package bookHandlers

import (
	"books/server/handler"
	"books/server/service"
	"encoding/json"
	"fmt"
	"net/http"
)

type getAllBooks struct {
	getAllBooks service.Book
}

func NewGetAllBooks() *getAllBooks {
	return new(getAllBooks)
}

func (g *getAllBooks) Inject(getBooks service.Book) {
	g.getAllBooks = getBooks
}

var _ http.Handler = (*deleteBook)(nil)
var _ handler.Router = (*deleteBook)(nil)

func (g *getAllBooks) Path() (path string) {
	return "/get_all_books"
}

func (g *getAllBooks) Method() (method string) {
	return http.MethodGet
}

func (g *getAllBooks) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	response := g.getAllBooks.GetAllBooks(ctx)

	if err := json.NewEncoder(w).Encode(response); err != nil {
		fmt.Fprintf(w, "error: %s", err.Error())
		w.WriteHeader(http.StatusBadRequest)
		return
	}
}
