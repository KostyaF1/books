package bookHandlers

import (
	"books/server/handler"
	"books/server/service"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
)

type getAllBooks struct {
	getAllBooks service.GetAllBooks
}

func NewGetAllBooks() *getAllBooks {
	return new(getAllBooks)
}

func (g *getAllBooks) Inject(getBooks service.GetAllBooks) {
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

func (g *getAllBooks) getAll(ctx context.Context) service.GetAllBooksResp {
	delBook := g.getAllBooks.GetAllBooks(ctx)

	return delBook
}

func (g *getAllBooks) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	var delBookReq service.DeleteBookReq
	defer r.Body.Close()
	ctx := r.Context()

	if err := json.NewDecoder(r.Body).Decode(&delBookReq); err != nil {
		fmt.Fprintf(w, "error: %s", err.Error())
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	response := g.getAll(ctx)

	if err := json.NewEncoder(w).Encode(response); err != nil {
		fmt.Fprintf(w, "error: %s", err.Error())
		w.WriteHeader(http.StatusBadRequest)
		return
	}
}
