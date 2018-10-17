package request

type AddComment struct {
	BookID          int64  `json:"book_id"`
	Body            string `json:"body"`
	Author          string `json:"author"`
	CommentParentID int64  `json:"father"`
}
