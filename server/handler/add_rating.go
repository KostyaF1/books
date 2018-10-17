package handler

import (
	"books/server/service"
	"books/server/service/request"
	"encoding/json"
	"fmt"
	"net/http"
)

type add struct {
	service service.Rating
}

func NewAdd() *add {
	return new(add)
}

func (a *add) Inject(rating service.Rating) {
	a.service = rating
}

var _ http.Handler = (*add)(nil)
var _ Router = (*add)(nil)

func (*add) Path() (path string) {
	return "/get_book_id/{id:[0-9]+}/add_rating"
}

func (*add) Method() (method string) {
	return http.MethodPost
}

func (a *add) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	var addRating request.Rating
	defer r.Body.Close()
	ctx := r.Context()

	if err := json.NewDecoder(r.Body).Decode(&addRating); err != nil {
		fmt.Fprintf(w, "error: %s", err.Error())
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	response := a.service.Add(ctx, addRating)

	if err := json.NewEncoder(w).Encode(response); err != nil {
		fmt.Fprintf(w, "error: %s", err.Error())
		w.WriteHeader(http.StatusBadRequest)
		return
	}
}
