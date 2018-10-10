package bookHandlers

import (
	"books/server/handler"
	"books/server/service"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
)

type deleteBook struct {
	deleteBook service.Book
}

func NewDeleteBook() *deleteBook {
	return new(deleteBook)
}

func (d *deleteBook) Inject(deleteBook service.Book) {
	d.deleteBook = deleteBook
}

var _ http.Handler = (*deleteBook)(nil)
var _ handler.Router = (*deleteBook)(nil)

func (d *deleteBook) Path() (path string) {
	return "/delete_book"
}

func (d *deleteBook) Method() (method string) {
	return http.MethodPost
}

func (d *deleteBook) delete(ctx context.Context, req service.DeleteBookReq) service.DeleteBookResp {
	delBook := d.deleteBook.DeleteBook(ctx, req)

	return delBook
}

func (d *deleteBook) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	var delBookReq service.DeleteBookReq
	defer r.Body.Close()
	ctx := r.Context()

	if err := json.NewDecoder(r.Body).Decode(&delBookReq); err != nil {
		fmt.Fprintf(w, "error: %s", err.Error())
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	response := d.delete(ctx, delBookReq)

	if err := json.NewEncoder(w).Encode(response); err != nil {
		fmt.Fprintf(w, "error: %s", err.Error())
		w.WriteHeader(http.StatusBadRequest)
		return
	}
}
