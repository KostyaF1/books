package request

type RatingReq struct {
	StoreUnitID int64 `json:"store_unit_id"`
	Value       int   `json:"value"`
}
