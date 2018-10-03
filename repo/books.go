package repo

import (
	"books/db"
	"books/dbo"
	"context"
)

type Books interface {
	Create(ctx context.Context, book *dbo.Book) error
}

func NewBooks() *books {
	return new(books)
}

var _ Books = (*books)(nil)

type books struct {
	query db.Query
}

func (repo *books) Inject(query db.Query) {
	repo.query = query
}

func (repo *books) Create(ctx context.Context, book *dbo.Book) error {
	return nil
}
