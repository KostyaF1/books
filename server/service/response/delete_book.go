package response

type DeleteBookResp struct {
	ID    int64  `json:"id"`
	Name  string `json:"name"`
	Error error  `json:"error"`
}
