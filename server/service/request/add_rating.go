package request

type Rating struct {
	ID    int64 `json:"store_unit_id"`
	Value int   `json:"value"`
}
