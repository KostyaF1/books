package response

import "books/server/db/dbo"

type (
	GetCommentsResp struct {
		Comments []*dbo.GetCommentsRepo `json:"comments"`
	}

	GetCommIDResp struct {
		Comm *dbo.GetCommIDRepo `json:"comm"`
	}
)
