package commentHandlers

import (
	"books/server/handler"
	"books/server/service"
	"books/server/service/request"
	"encoding/json"
	"fmt"
	"net/http"
)

type addAnswer struct {
	service service.Comment
}

func NewAddCommentAnswer() *addAnswer {
	return new(addAnswer)
}

func (a *addAnswer) Inject(newAnswer service.Comment) {
	a.service = newAnswer
}

var _ http.Handler = (*addAnswer)(nil)
var _ handler.Router = (*addAnswer)(nil)

func (*addAnswer) Path() (path string) {
	return "/get_book_id/{id:[0-9]+}/get_comments/get_comment_id/{id:[0-9]+}/add_answer"
}

func (*addAnswer) Method() (method string) {
	return http.MethodPost
}

func (a *addAnswer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	var addComment request.AddCommentReq
	defer r.Body.Close()
	ctx := r.Context()

	if err := json.NewDecoder(r.Body).Decode(&addComment); err != nil {
		fmt.Fprintf(w, "error: %s", err.Error())
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	response := a.service.AddCommentAnswer(ctx, addComment)

	if err := json.NewEncoder(w).Encode(response); err != nil {
		fmt.Fprintf(w, "error: %s", err.Error())
		w.WriteHeader(http.StatusBadRequest)
		return
	}
}
