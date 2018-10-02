package interfaces

import (
	"books/server/db/dbqueries"
	"context"
)

//Geter ...
type Geter interface {
	GetAllUnits(cxt context.Context) []*dbqueries.GeneralQuerie
}
