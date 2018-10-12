package commentHandlers

import (
	"books/server/handler"
	"books/server/service"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

type getCommentByID struct {
	getComment service.Comment
}

func NewGetCommentByID() *getCommentByID {
	return new(getCommentByID)
}

func (g *getCommentByID) Inject(getComm service.Comment) {
	g.getComment = getComm
}

var _ http.Handler = (*getCommentByID)(nil)
var _ handler.Router = (*getCommentByID)(nil)

func (g *getCommentByID) Path() (path string) {
	return "/get_book_id/{id:[0-9]+}/get_comments/get_comment_id/{id:[0-9]+}"
}

func (g *getCommentByID) Method() (method string) {
	return http.MethodGet
}

func (g *getCommentByID) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	vars := mux.Vars(r)
	id := vars["id"]
	id64, err := strconv.ParseInt(id, 10, 64)

	if err != nil {
		fmt.Fprintf(w, "error: %s", err.Error())
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	response := g.getComment.GetCommByID(ctx, id64)

	if err := json.NewEncoder(w).Encode(response); err != nil {
		fmt.Fprintf(w, "error: %s", err.Error())
		w.WriteHeader(http.StatusBadRequest)
		return
	}
}