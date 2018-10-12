package repo

import (
	"books/server/db"
	"books/server/db/dbo"
	"books/server/db/query"
	"context"
	"log"
)

type Comments interface {
	Add(ctx context.Context, coment dbo.Comment) (*dbo.AddCommentRepo, error)
	AddAnswer(ctx context.Context, coment dbo.Comment) (*dbo.AddCommentRepo, error)
	GetCommentsRepo(ctx context.Context, bookID int64) []*dbo.GetCommentsRepo
	GetByID(ctx context.Context, commID int64) *dbo.GetCommIDRepo
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

func (c *comments) Add(ctx context.Context, comment dbo.Comment) (*dbo.AddCommentRepo, error) {
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

	return &dbo.AddCommentRepo{
		ID:     commentID,
		Author: comment.Author,
	}, tx.Commit()
}

func (c *comments) AddAnswer(ctx context.Context, comment dbo.Comment) (*dbo.AddCommentRepo, error) {
	dbConn := c.dbConn.Connect()
	tx, err := dbConn.BeginTx(ctx, nil)

	if err != nil {
		return nil, err
	}
	var commentID int64

	err = tx.QueryRowContext(ctx, query.AddCommentAnswer,
		comment.BookID, comment.Body, comment.Author, comment.Father).Scan(&commentID)
	if err != nil {
		return nil, err
	}

	return &dbo.AddCommentRepo{
		ID:     commentID,
		Author: comment.Author,
		Father: comment.Father,
	}, tx.Commit()
}

func (c *comments) GetCommentsRepo(ctx context.Context, bookID int64) []*dbo.GetCommentsRepo {
	dbConn := c.dbConn.Connect()
	rows, err := dbConn.QueryContext(ctx, query.GetComments, bookID)
	if err != nil {
		log.Println("Error when GetAll" + err.Error())
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

	return comments
}

func (c *comments) GetByID(ctx context.Context, commID int64) *dbo.GetCommIDRepo {
	dbConn := c.dbConn.Connect()
	rows, err := dbConn.QueryContext(ctx, query.GetCommByID, commID)
	if err != nil {
		log.Println("Error when GetAll" + err.Error())
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
	}
}

func (c *comments) Delete(ctx context.Context, id int64) error {
	dbConn := c.dbConn.Connect()
	dbConn.QueryContext(ctx, "Insert Into")

	return nil
}
