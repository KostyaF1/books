package response

import "books/server/db/dbo"

type (
	GetComments struct {
		Comments []*dbo.GetComments `json:"comments"`
		Error    string             `json:"error"`
	}

	GetCommID struct {
		Comm  *dbo.GetCommID `json:"comm"`
		Error string         `json:"error"`
	}
)
