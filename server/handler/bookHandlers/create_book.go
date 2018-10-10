package bookHandlers

import (
	"books/server/handler"
	"books/server/service"
	"encoding/json"
	"fmt"
	"net/http"
)

type createBook struct {
	createBook service.Book
}

func NewCreateBook() *createBook {
	return new(createBook)
}

func (cb *createBook) Inject(newBook service.Book) {
	cb.createBook = newBook
}

var _ http.Handler = (*createBook)(nil)
var _ handler.Router = (*createBook)(nil)

func (cb *createBook) Path() (path string) {
	return "/create_book"
}

func (cb *createBook) Method() (method string) {
	return http.MethodPost
}

//func (cb *createBook) create(ctx context.Context, req service.CreateBookReq) service.CreateBookResp {
//	newBook := cb.createBook.CreateBook(ctx, req)
//
//	return newBook
//}

func (cb *createBook) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	var newBookReq service.CreateBookReq
	defer r.Body.Close()
	ctx := r.Context()

	if err := json.NewDecoder(r.Body).Decode(&newBookReq); err != nil {
		fmt.Fprintf(w, "error: %s", err.Error())
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	response := cb.createBook.CreateBook(ctx, newBookReq)

	if err := json.NewEncoder(w).Encode(response); err != nil {
		fmt.Fprintf(w, "error: %s", err.Error())
		w.WriteHeader(http.StatusBadRequest)
		return
	}
}
