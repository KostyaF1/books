package dbo

type Comment struct {
	ID     int64  `json:"id"`
	BookID int64  `json:"book_id"`
	Body   string `json:"body"`
	Author string `json:"author"`
	Father int64  `json:"father"`
}

type (
	AddCommentRepo struct {
		ID     int64  `json:"id"`
		Author string `json:"author"`
		Father int64  `json:"father"`
		Error  error  `json:"error"`
	}

	GetCommentsRepo struct {
		Body   string `json:"body"`
		Author string `json:"author"`
		Error  error  `json:"error"`
	}

	GetCommIDRepo struct {
		Comment string `json:"comment"`
		Answer  string `json:"answer"`
		Error   error  `json:"error"`
	}
)
