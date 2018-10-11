package commentHandlers

import (
	"books/server/handler"
	"books/server/service"
	"encoding/json"
	"fmt"
	"net/http"
)

type addCommentAnswer struct {
	addComment service.Comment
}

func NewAddCommentAnswer() *addCommentAnswer {
	return new(addCommentAnswer)
}

func (ac *addCommentAnswer) Inject(newAnswer service.Comment) {
	ac.addComment = newAnswer
}

var _ http.Handler = (*addComment)(nil)
var _ handler.Router = (*addComment)(nil)

func (*addCommentAnswer) Path() (path string) {
	return "/get_book_id/{id:[0-9]+}/get_comments/get_comment_id/{id:[0-9]+}/add_answer"
}

func (*addCommentAnswer) Method() (method string) {
	return http.MethodPost
}

func (ac *addCommentAnswer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	var addNewComment service.AddCommentReq
	defer r.Body.Close()
	ctx := r.Context()

	if err := json.NewDecoder(r.Body).Decode(&addNewComment); err != nil {
		fmt.Fprintf(w, "error: %s", err.Error())
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	response := ac.addComment.AddCommentAnswer(ctx, addNewComment)

	if err := json.NewEncoder(w).Encode(response); err != nil {
		fmt.Fprintf(w, "error: %s", err.Error())
		w.WriteHeader(http.StatusBadRequest)
		return
	}
}
