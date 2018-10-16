package response

import (
	"books/server/db/dbo"
)

type (
	GetBookResp struct {
		AllBooks []*dbo.GetBookRepo `json:"all_books"`
		Error    string             `json:"error"`
	}
	GetBookIDResp struct {
		Book  *dbo.GetBookIDRepo `json:"book"`
		Error string             `json:"error"`
	}
)
