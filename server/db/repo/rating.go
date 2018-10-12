package repo

import (
	"books/server/db"
	"books/server/db/query"
	"context"
)

type RatingRepo interface {
	Add(ctx context.Context, bookID int64, value int) (int64, error)
}

type ratingRepo struct {
	dbConn db.Connector
}

func NewRating() *ratingRepo {
	return new(ratingRepo)
}

var _ RatingRepo = (*ratingRepo)(nil)

func (r *ratingRepo) Inject(conn db.Connector) {
	r.dbConn = conn
}

func (r *ratingRepo) Add(ctx context.Context, bookID int64, value int) (int64, error) {
	dbConn := r.dbConn.Connect()
	tx, err := dbConn.BeginTx(ctx, nil)

	if err != nil {
		return 0, err
	}

	var ratingID int64

	err = tx.QueryRowContext(ctx, query.AddRating, bookID, value).Scan(&ratingID)

	if err != nil {
		return 0, err
	}

	return ratingID, tx.Commit()
}
