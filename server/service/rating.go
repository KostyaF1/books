package service

import (
	"books/server/db/repo"
	"books/server/service/request"
	"books/server/service/response"
	"context"
)

type Rating interface {
	Add(ctx context.Context, req request.Rating) response.Rating
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

func (r *rating) Add(ctx context.Context, req request.Rating) response.Rating {
	if err := r.addRating.Add(ctx, req.ID, req.Value); err != nil {
		return response.Rating{
			Error: err.Error(),
		}
	}
	return response.Rating{
		Error: "nil",
	}
}
