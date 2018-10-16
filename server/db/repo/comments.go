package repo

import (
	"books/server/db"
	"books/server/db/dbo"
	"books/server/db/query"
	"context"
)

type Comments interface {
	Add(ctx context.Context, coment dbo.Comment) error
	AddAnswer(ctx context.Context, coment dbo.Comment) error
	GetCommentsRepo(ctx context.Context, bookID int64) ([]*dbo.GetCommentsRepo, error)
	GetByID(ctx context.Context, commID int64) (*dbo.GetCommIDRepo, error)
	Delete(ctx context.Context, id int64) error
}

type comments struct {
	dbConn db.Connector
}

func NewComments() *comments {
	return new(comments)
}

var _ Comments = (*comments)(nil)

func (c *comments) Inject(conn db.Connector) {
	c.dbConn = conn
}

func (c *comments) Add(ctx context.Context, comment dbo.Comment) error {
	dbConn := c.dbConn.Connect()
	tx, err := dbConn.BeginTx(ctx, nil)

	_, err = tx.ExecContext(ctx, query.AddComment,
		comment.BookID, comment.Body, comment.Author)

	return err
}

func (c *comments) AddAnswer(ctx context.Context, comment dbo.Comment) error {
	dbConn := c.dbConn.Connect()
	tx, err := dbConn.BeginTx(ctx, nil)

	_, err = tx.ExecContext(ctx, query.AddCommentAnswer,
		comment.BookID, comment.Body, comment.Author, comment.Father)

	return err
}

func (c *comments) GetCommentsRepo(ctx context.Context, bookID int64) ([]*dbo.GetCommentsRepo, error) {
	dbConn := c.dbConn.Connect()
	rows, err := dbConn.QueryContext(ctx, query.GetComments, bookID)
	if err != nil {
		return nil, err
	}

	var comments []*dbo.GetCommentsRepo

	for rows.Next() {
		var (
			body   string
			author string
		)
		rows.Scan(&body, &author)
		comments = append(comments, &dbo.GetCommentsRepo{
			Body:   body,
			Author: author,
		})
	}

	return comments, err
}

func (c *comments) GetByID(ctx context.Context, commID int64) (*dbo.GetCommIDRepo, error) {
	dbConn := c.dbConn.Connect()
	rows, err := dbConn.QueryContext(ctx, query.GetCommByID, commID)
	if err != nil {
		return nil, err
	}

	var (
		comment string
		answer  string
	)

	rows.Next()

	rows.Scan(&comment, &answer)

	return &dbo.GetCommIDRepo{
		Comment: comment,
		Answer:  answer,
	}, err
}

func (c *comments) Delete(ctx context.Context, id int64) error {
	dbConn := c.dbConn.Connect()
	dbConn.QueryContext(ctx, "Insert Into")

	return nil
}
