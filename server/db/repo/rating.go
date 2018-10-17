package repo

import (
	"books/server/db"
	"books/server/db/query"
	"context"
)

type Ratings interface {
	Add(ctx context.Context, bookID int64, value int) error
}

type ratings struct {
	dbConn db.Connector
}

func NewRating() *ratings {
	return new(ratings)
}

var _ Ratings = (*ratings)(nil)

func (r *ratings) Inject(conn db.Connector) {
	r.dbConn = conn
}

func (r *ratings) Add(ctx context.Context, bookID int64, value int) error {
	dbConn := r.dbConn.Connect()
	tx, err := dbConn.BeginTx(ctx, nil)

	_, err = tx.ExecContext(ctx, query.AddRating, bookID, value)

	if err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit()
}
