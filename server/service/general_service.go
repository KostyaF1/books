package service

import (
	"books/server/service/interfaces"
	"context"
)

//GeneralService ...
type GeneralService struct {
	interfaces.Geter
}

//NewGeneralService ...
func NewGeneralService(geter interfaces.Geter) *GeneralService {
	return &GeneralService{
		Geter: geter,
	}
}

//GetAllUnits ...
func (gs *GeneralService) GetAllUnits(ctx context.Context) string {
	allUnits := gs.Geter.GetAllUnits(ctx)
	response := string(JsonMarshal(allUnits))
	return response
}
