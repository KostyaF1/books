package book_handler

import (
	"books/server/handler"
	"books/server/service"
	"books/server/service/request"
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"net/http"
)

type create struct {
	service service.Book
}

func NewCreate() *create {
	return new(create)
}

func (c *create) Inject(newBook service.Book) {
	c.service = newBook
}

var (
	_ http.Handler   = (*create)(nil)
	_ handler.Router = (*create)(nil)
)

func (*create) Path() (path string) {
	return "/book"
}

func (*create) Method() (method string) {
	return http.MethodPost
}

func (c *create) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	var bookReq request.CreateBook
	defer r.Body.Close()
	ctx := r.Context()

	if err := json.NewDecoder(r.Body).Decode(&bookReq); err != nil {
		//if _, err := fmt.Fprintf(w, "error: %s", err.Error()); err != nil {
		//	w.WriteHeader(http.StatusBadRequest)
		//	return
		//}
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	response := c.service.Create(ctx, bookReq)
	if ok := assert.Empty(nil, response, nil); ok == true {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if err := json.NewEncoder(w).Encode(response); err != nil {
		//if _, err = fmt.Fprintf(w, "error: %s", err.Error()); err != nil {
		//	w.WriteHeader(http.StatusBadRequest)
		//	return
		//}
		w.WriteHeader(http.StatusBadRequest)
		return
	}
}
