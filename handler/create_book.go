package handler

import (
	"books/service"
	"encoding/json"
	"fmt"
	"net/http"
)

func NewCreateBook() *createBook {
	return new(createBook)
}

var (
	_ http.Handler = (*createBook)(nil)
	_ Router       = (*createBook)(nil)
)

type createBook struct {
	bookStore service.BookStore
}

func (handler *createBook) Path() (path string) {
	return "/books/create"
}

func (handler *createBook) Method() (method string) {
	return http.MethodPost
}

func (handler *createBook) Inject(bookStore service.BookStore) {
	handler.bookStore = bookStore
}

func (handler *createBook) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	var bookReq service.CreateBookRequest
	defer req.Body.Close()

	if err := json.NewDecoder(req.Body).Decode(&bookReq); err != nil {
		fmt.Fprintf(res, "error: %s", err.Error())
		res.WriteHeader(http.StatusBadRequest)
		return
	}

	response := handler.bookStore.CreateBook(req.Context(), bookReq)

	if err := json.NewEncoder(res).Encode(response); err != nil {
		fmt.Fprintf(res, "error: %s", err.Error())
		res.WriteHeader(http.StatusBadRequest)
		return
	}
}
