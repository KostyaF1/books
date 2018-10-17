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

type getByID struct {
	service service.Comment
}

func NewGetByID() *getByID {
	return new(getByID)
}

func (g *getByID) Inject(getComm service.Comment) {
	g.service = getComm
}

var _ http.Handler = (*getByID)(nil)
var _ handler.Router = (*getByID)(nil)

func (*getByID) Path() (path string) {
	return "/get_book_id/{id:[0-9]+}/get_comments/get_comment_id/{id:[0-9]+}"
}

func (*getByID) Method() (method string) {
	return http.MethodGet
}

func (g *getByID) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	vars := mux.Vars(r)
	id := vars["id"]
	id64, err := strconv.ParseInt(id, 10, 64)

	if err != nil {
		fmt.Fprintf(w, "error: %s", err.Error())
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	response := g.service.GetByID(ctx, id64)

	if err := json.NewEncoder(w).Encode(response); err != nil {
		fmt.Fprintf(w, "error: %s", err.Error())
		w.WriteHeader(http.StatusBadRequest)
		return
	}
}
