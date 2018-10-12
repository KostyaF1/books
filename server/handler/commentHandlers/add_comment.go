package commentHandlers

import (
	"books/server/handler"
	"books/server/service"
	"encoding/json"
	"fmt"
	"net/http"
)

type addComment struct {
	addComment service.Comment
}

func NewAddComment() *addComment {
	return new(addComment)
}

func (ac *addComment) Inject(newComment service.Comment) {
	ac.addComment = newComment
}

var _ http.Handler = (*addComment)(nil)
var _ handler.Router = (*addComment)(nil)

func (*addComment) Path() (path string) {
	return "/get_book_id/{id:[0-9]+}/add_comment"
}

func (*addComment) Method() (method string) {
	return http.MethodPost
}

func (ac *addComment) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	var addNewComment service.AddCommentReq
	defer r.Body.Close()
	ctx := r.Context()

	if err := json.NewDecoder(r.Body).Decode(&addNewComment); err != nil {
		fmt.Fprintf(w, "error: %s", err.Error())
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	response := ac.addComment.AddComment(ctx, addNewComment)

	if err := json.NewEncoder(w).Encode(response); err != nil {
		fmt.Fprintf(w, "error: %s", err.Error())
		w.WriteHeader(http.StatusBadRequest)
		return
	}
}
