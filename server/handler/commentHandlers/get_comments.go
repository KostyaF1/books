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

type getCommentsHandler struct {
	getComments service.Comment
}

func NewGetCommentsHandler() *getCommentsHandler {
	return new(getCommentsHandler)
}

func (g *getCommentsHandler) Inject(getComments service.Comment) {
	g.getComments = getComments
}

var _ http.Handler = (*getCommentsHandler)(nil)
var _ handler.Router = (*getCommentsHandler)(nil)

func (g *getCommentsHandler) Path() (path string) {
	return "/get_book_id/{id:[0-9]+}/get_comments"
}

func (g *getCommentsHandler) Method() (method string) {
	return http.MethodGet
}

func (g *getCommentsHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	vars := mux.Vars(r)
	id := vars["id"]

	id64, err := strconv.ParseInt(id, 10, 64)

	if err != nil {
		fmt.Fprintf(w, "error: %s", err.Error())
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	response := g.getComments.GetComments(ctx, id64)

	if err := json.NewEncoder(w).Encode(response); err != nil {
		fmt.Fprintf(w, "error: %s", err.Error())
		w.WriteHeader(http.StatusBadRequest)
		return
	}
}
