package service

import (
	"books/server/db/repo"
	"books/server/service/request"
	"books/server/service/response"
	"context"
)

type Rating interface {
	AddRating(ctx context.Context, req request.RatingReq) response.RatingResp
}

type rating struct {
	addRating repo.Ratings
}

func NewRating() *rating {
	return new(rating)
}

var _ Rating = (*rating)(nil)

func (r *rating) Inject(addRating repo.Ratings) {
	r.addRating = addRating
}

func (r *rating) AddRating(ctx context.Context, req request.RatingReq) response.RatingResp {
	err := r.addRating.Add(ctx, req.StoreUnitID, req.Value)

	if err != nil {
		return response.RatingResp{
			Error: err,
		}
	}
	return response.RatingResp{
		Error: nil,
	}
}
