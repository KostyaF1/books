package handler

import (
	"books/server/routing"
	"fmt"
	"net/http"
)

//GeneralHandler ...
type GeneralHandler struct {
	routIn *routing.RouterIn
}

//NewGeneralHandler ...
func NewGeneralHandler(routIn *routing.RouterIn) *GeneralHandler {
	return &GeneralHandler{
		routIn: routIn,
	}
}

//ServeHTTP ...
func (gh GeneralHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	uri := r.RequestURI
	ctx := r.Context()
	fmt.Fprintf(w, gh.routIn.Rout(ctx, uri))
}

var _ http.Handler = (*GeneralHandler)(nil)

//Run is a http handler function
func Run(generalHandler GeneralHandler) {
	http.HandleFunc("/", generalHandler.ServeHTTP)
	http.ListenAndServe(":8081", nil)
}
