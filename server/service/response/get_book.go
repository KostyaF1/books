package response

import (
	"books/server/db/dbo"
)

type (
	GetBookResp struct {
		AllBooks []*dbo.GetBookRepo `json:"all_books"`
		Error    error
	}
	GetBookIDResp struct {
		Book *dbo.GetBookIDRepo `json:"book"`
	}
)
