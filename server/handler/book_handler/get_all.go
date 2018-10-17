package book_handler

import (
	"books/server/handler"
	"books/server/service"
	"encoding/json"
	"fmt"
	"net/http"
)

type getAll struct {
	service service.Book
}

func NewGetAll() *getAll {
	return new(getAll)
}

func (g *getAll) Inject(getBooks service.Book) {
	g.service = getBooks
}

var _ http.Handler = (*getAll)(nil)
var _ handler.Router = (*getAll)(nil)

func (*getAll) Path() (path string) {
	return "/get_all_books"
}

func (*getAll) Method() (method string) {
	return http.MethodGet
}

func (g *getAll) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	response := g.service.GetAll(ctx)

	if err := json.NewEncoder(w).Encode(response); err != nil {
		fmt.Fprintf(w, "error: %s", err.Error())
		w.WriteHeader(http.StatusBadRequest)
		return
	}
}
