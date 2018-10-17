package response

import (
	"books/server/db/dbo"
)

type (
	GetBook struct {
		AllBooks []*dbo.GetBook `json:"all_books"`
		Error    string         `json:"error"`
	}
	GetBookID struct {
		Book  *dbo.GetBookID `json:"book"`
		Error string         `json:"error"`
	}
)
