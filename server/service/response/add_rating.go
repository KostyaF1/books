package response

type RatingResp struct {
	ID    int64 `json:"id"`
	Error error `json:"error"`
}
