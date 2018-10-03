package mock_repo

import (
	"books/dbo"
	"books/repo"
	"context"
)

var _ repo.Books = (*BooksMock)(nil)

type BooksMock struct {
	Mock struct {
		Create func(ctx context.Context, book *dbo.Book) error
	}
}

func (mock *BooksMock) Create(ctx context.Context, book *dbo.Book) error {
	return mock.Mock.Create(ctx, book)
}
