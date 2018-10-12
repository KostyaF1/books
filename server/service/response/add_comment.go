package response

type AddCommentResp struct {
	ID     int64  `json:"id"`
	Author string `json:"author"`
	Father int64  `json:"father"`
	Error  error
}
