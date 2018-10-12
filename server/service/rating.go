package service

import (
	"books/server/db/repo"
	"context"
)

type (
	RatingReq struct {
		StoreUnitID int64 `json:"store_unit_id"`
		Value       int   `json:"value"`
	}

	RatingResp struct {
		ID    int64 `json:"id"`
		Error error `json:"error"`
	}
)

type Rating interface {
	AddRating(ctx context.Context, req RatingReq) RatingResp
}

type rating struct {
	addRating repo.RatingRepo
}

func NewRating() *rating {
	return new(rating)
}

var _ Rating = (*rating)(nil)

func (r *rating) Inject(addRating repo.RatingRepo) {
	r.addRating = addRating
}

func (r *rating) AddRating(ctx context.Context, req RatingReq) RatingResp {
	ratingID, err := r.addRating.Add(ctx, req.StoreUnitID, req.Value)

	if err != nil {
		return RatingResp{
			Error: err,
		}
	}
	return RatingResp{
		ID: ratingID,
	}
}
