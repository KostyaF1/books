package request

type AddCommentReq struct {
	BookID int64  `json:"book_id"`
	Body   string `json:"body"`
	Author string `json:"author"`
	Father int64  `json:"father"`
}
