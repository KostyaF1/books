package repo

import (
	"books/server/db"
	"books/server/db/dbo"
	"books/server/db/query"
	"context"
)

type (
	AddCommentRepo struct {
		ID     int64  `json:"id"`
		Author string `json:"author"`
		Father int64  `json:"father"`
		Error  error  `json:"error"`
	}
)

type Comment interface {
	Add(ctx context.Context, coment dbo.Comment) (*AddCommentRepo, error)
	AddAnswer(ctx context.Context, coment dbo.Comment) (*AddCommentRepo, error)
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

func (c *comment) Add(ctx context.Context, comment dbo.Comment) (*AddCommentRepo, error) {
	dbConn := c.dbConn.Connect()
	tx, err := dbConn.BeginTx(ctx, nil)

	if err != nil {
		return nil, err
	}
	var commentID int64

	err = tx.QueryRowContext(ctx, query.AddComment,
		comment.BookID, comment.Body, comment.Author).Scan(&commentID)
	if err != nil {
		return nil, err
	}

	return &AddCommentRepo{
		ID:     commentID,
		Author: comment.Author,
	}, tx.Commit()
}

func (c *comment) AddAnswer(ctx context.Context, comment dbo.Comment) (*AddCommentRepo, error) {
	dbConn := c.dbConn.Connect()
	dbConn.QueryContext(ctx, "Insert Into")

	return nil, nil
}

func (c *comment) Delete(ctx context.Context, id int64) error {
	dbConn := c.dbConn.Connect()
	dbConn.QueryContext(ctx, "Insert Into")

	return nil
}
