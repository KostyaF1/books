package comment_handler

import (
	"books/server/handler"
	"books/server/service"
	"books/server/service/request"
	"encoding/json"
	"fmt"
	"net/http"
)

type add struct {
	service service.Comment
}

func NewAdd() *add {
	return new(add)
}

func (a *add) Inject(newComment service.Comment) {
	a.service = newComment
}

var _ http.Handler = (*add)(nil)
var _ handler.Router = (*add)(nil)

func (*add) Path() (path string) {
	return "/get_book_id/{id:[0-9]+}/add_comment"
}

func (*add) Method() (method string) {
	return http.MethodPost
}

func (a *add) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	var addNewComment request.AddComment
	defer r.Body.Close()
	ctx := r.Context()

	if err := json.NewDecoder(r.Body).Decode(&addNewComment); err != nil {
		fmt.Fprintf(w, "error: %s", err.Error())
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	response := a.service.Add(ctx, addNewComment)

	if err := json.NewEncoder(w).Encode(response); err != nil {
		fmt.Fprintf(w, "error: %s", err.Error())
		w.WriteHeader(http.StatusBadRequest)
		return
	}
}
