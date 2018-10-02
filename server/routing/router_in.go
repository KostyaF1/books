package routing

import (
	"books/server/service"
	"context"
)

//RouterIn ...
type RouterIn struct {
	service *service.GeneralService
}

//NewRouterIn ...
func NewRouterIn(service *service.GeneralService) *RouterIn {
	return &RouterIn{
		service: service,
	}
}

//Rout ...
func (r *RouterIn) Rout(ctx context.Context, uri string) string {
	switch {
	case uri == "/getall":
		return r.service.GetAllUnits(ctx)
	}
	return "Welcome Home =)"
}
