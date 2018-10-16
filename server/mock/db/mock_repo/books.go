package mock_repo

import (
	"books/server/db/dbo"
	"books/server/db/repo"
	"context"
)

type Books struct {
	Mock struct {
		Create  func(ctx context.Context, book *dbo.Book) error
		Delete  func(ctx context.Context, id int64) error
		GetAll  func(ctx context.Context) ([]*dbo.GetBookRepo, error)
		GetByID func(ctx context.Context, bookID int64) (*dbo.GetBookIDRepo, error)
	}
}

func (repo *Books) Create(ctx context.Context, book *dbo.Book) error {
	return repo.Mock.Create(ctx, book)
}

func (repo *Books) Delete(ctx context.Context, id int64) error {
	return repo.Mock.Delete(ctx, id)
}

func (repo *Books) GetAll(ctx context.Context) ([]*dbo.GetBookRepo, error) {
	return repo.Mock.GetAll(ctx)
}

func (repo *Books) GetByID(ctx context.Context, bookID int64) (*dbo.GetBookIDRepo, error) {
	return repo.Mock.GetByID(ctx, bookID)
}

var _ repo.Books = (*Books)(nil)
