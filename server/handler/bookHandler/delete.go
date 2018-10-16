package bookHandler

import (
	"books/server/handler"
	"books/server/service"
	"books/server/service/request"
	"encoding/json"
	"fmt"
	"net/http"
)

type delete struct {
	service service.Book
}

func NewDeleteBook() *delete {
	return new(delete)
}

func (d *delete) Inject(deleteBook service.Book) {
	d.service = deleteBook
}

var _ http.Handler = (*delete)(nil)
var _ handler.Router = (*delete)(nil)

func (*delete) Path() (path string) {
	return "/delete_book"
}

func (*delete) Method() (method string) {
	return http.MethodPost
}

func (d *delete) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	var bookReq request.DeleteBookReq
	defer r.Body.Close()
	ctx := r.Context()

	if err := json.NewDecoder(r.Body).Decode(&bookReq); err != nil {
		fmt.Fprintf(w, "error: %s", err.Error())
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	resp := d.service.DeleteBook(ctx, bookReq)

	if err := json.NewEncoder(w).Encode(resp); err != nil {
		fmt.Fprintf(w, "error: %s", err.Error())
		w.WriteHeader(http.StatusBadRequest)
		return
	}
}
