package ratingHandler

import (
	"books/server/handler"
	"books/server/service"
	"encoding/json"
	"fmt"
	"net/http"
)

type addRating struct {
	addRating service.Rating
}

func NewAddRating() *addRating {
	return new(addRating)
}

func (ar *addRating) Inject(rating service.Rating) {
	ar.addRating = rating
}

var _ http.Handler = (*addRating)(nil)
var _ handler.Router = (*addRating)(nil)

func (*addRating) Path() (path string) {
	return "/get_book_id/{id:[0-9]+}/add_rating"
}

func (*addRating) Method() (method string) {
	return http.MethodPost
}

func (ar *addRating) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	var addRating service.RatingReq
	defer r.Body.Close()
	ctx := r.Context()

	if err := json.NewDecoder(r.Body).Decode(&addRating); err != nil {
		fmt.Fprintf(w, "error: %s", err.Error())
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	response := ar.addRating.AddRating(ctx, addRating)

	if err := json.NewEncoder(w).Encode(response); err != nil {
		fmt.Fprintf(w, "error: %s", err.Error())
		w.WriteHeader(http.StatusBadRequest)
		return
	}
}
