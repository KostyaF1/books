package comment_handler

import (
	"books/server/handler"
	"books/server/service"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

type getAll struct {
	service service.Comment
}

func NewGetAll() *getAll {
	return new(getAll)
}

func (g *getAll) Inject(getAll service.Comment) {
	g.service = getAll
}

var _ http.Handler = (*getAll)(nil)
var _ handler.Router = (*getAll)(nil)

func (*getAll) Path() (path string) {
	return "/get_book_id/{id:[0-9]+}/get_comments"
}

func (*getAll) Method() (method string) {
	return http.MethodGet
}

func (g *getAll) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	vars := mux.Vars(r)
	id := vars["id"]

	id64, err := strconv.ParseInt(id, 10, 64)

	if err != nil {
		fmt.Fprintf(w, "error: %s", err.Error())
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	response := g.service.GetAll(ctx, id64)

	if err := json.NewEncoder(w).Encode(response); err != nil {
		fmt.Fprintf(w, "error: %s", err.Error())
		w.WriteHeader(http.StatusBadRequest)
		return
	}
}
