package mock_repo

import (
	"books/server/db/dbo"
	"books/server/db/repo"
	"context"
)

type Comments struct {
	Mock struct {
		Add       func(ctx context.Context, comment dbo.Comment) error
		AddAnswer func(ctx context.Context, comment dbo.Comment) error
		GetAll    func(ctx context.Context, bookID int64) ([]*dbo.GetComments, error)
		GetByID   func(ctx context.Context, commID int64) (*dbo.GetCommID, error)
		Delete    func(ctx context.Context, id int64) error
	}
}

func (c *Comments) Add(ctx context.Context, comment dbo.Comment) error {
	return c.Mock.Add(ctx, comment)
}

func (c *Comments) AddAnswer(ctx context.Context, comment dbo.Comment) error {
	return c.Mock.AddAnswer(ctx, comment)
}

func (c *Comments) GetAll(ctx context.Context, bookID int64) ([]*dbo.GetComments, error) {
	return c.Mock.GetAll(ctx, bookID)
}

func (c *Comments) GetByID(ctx context.Context, commID int64) (*dbo.GetCommID, error) {
	return c.Mock.GetByID(ctx, commID)
}

func (c *Comments) Delete(ctx context.Context, id int64) error {
	return c.Mock.Delete(ctx, id)
}

var _ repo.Comments = (*Comments)(nil)
