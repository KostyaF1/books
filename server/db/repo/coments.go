package repo

import (
	"books/server/db"
	"books/server/db/dbo"
	"context"
)

type Comment interface {
	Add(ctx context.Context, coment dbo.Coment) error
	Delete(ctx context.Context, id int64) error
}

type comment struct {
	dbConn db.Connector
}

func NewComment() *comment {
	return new(comment)
}

var _ Comment = (*comment)(nil)

func (c *comment) Inject(conn db.Connector) {
	c.dbConn = conn
}

func (c *comment) Add(ctx context.Context, comment dbo.Coment) error {
	dbConn := c.dbConn.Connect()
	dbConn.QueryContext(ctx, "Insert Into")

	return nil
}

func (c *comment) Delete(ctx context.Context, id int64) error {
	dbConn := c.dbConn.Connect()
	dbConn.QueryContext(ctx, "Insert Into")

	return nil
}
