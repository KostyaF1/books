package dbo

type Comment struct {
	ID              int64  `json:"id"`
	BookID          int64  `json:"book_id"`
	Body            string `json:"body"`
	Author          string `json:"author"`
	CommentParentID int64  `json:"father"`
}

type (
	GetComments struct {
		Body   string `json:"body"`
		Author string `json:"author"`
		Error  error  `json:"error"`
	}

	GetCommID struct {
		Comment string `json:"comment"`
		Answer  string `json:"answer"`
		Error   error  `json:"error"`
	}
)
