package repo

import (
	"books/server/db"
	"books/server/db/dbo"
	"books/server/db/query"
	"context"
	"log"
)

type (
	AddCommentRepo struct {
		ID     int64  `json:"id"`
		Author string `json:"author"`
		Father int64  `json:"father"`
		Error  error  `json:"error"`
	}

	GetCommentsRepo struct {
		Body   string `json:"body"`
		Author string `json:"author"`
		Error  error  `json:"error"`
	}

	GetCommIDRepo struct {
		Body   string `json:"body"`
		Author string `json:"author"`
		Answer string `json:"answer"`
		Error  error  `json:"error"`
	}
)

type Comments interface {
	Add(ctx context.Context, coment dbo.Comment) (*AddCommentRepo, error)
	AddAnswer(ctx context.Context, coment dbo.Comment) (*AddCommentRepo, error)
	GetCommentsRepo(ctx context.Context, bookID int64) []*GetCommentsRepo
	GetByID(ctx context.Context, commID int64) *GetCommIDRepo
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

func (c *comments) Add(ctx context.Context, comment dbo.Comment) (*AddCommentRepo, error) {
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

func (c *comments) AddAnswer(ctx context.Context, comment dbo.Comment) (*AddCommentRepo, error) {
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

	return &AddCommentRepo{
		ID:     commentID,
		Author: comment.Author,
		Father: comment.Father,
	}, tx.Commit()
}

func (c *comments) GetCommentsRepo(ctx context.Context, bookID int64) []*GetCommentsRepo {
	dbConn := c.dbConn.Connect()
	rows, err := dbConn.QueryContext(ctx, query.GetComments, bookID)
	if err != nil {
		log.Println("Error when GetAll" + err.Error())
	}

	var comments []*GetCommentsRepo

	for rows.Next() {
		var (
			body   string
			author string
		)
		rows.Scan(&body, &author)
		comments = append(comments, &GetCommentsRepo{
			Body:   body,
			Author: author,
		})
	}

	return comments
}

func (c *comments) GetByID(ctx context.Context, commID int64) *GetCommIDRepo {
	dbConn := c.dbConn.Connect()
	rows, err := dbConn.QueryContext(ctx, query.GetCommByID, commID)
	if err != nil {
		log.Println("Error when GetAll" + err.Error())
	}

	var (
		body   string
		author string
		answer string
	)

	rows.Next()

	rows.Scan(&body, &author, &answer)

	return &GetCommIDRepo{
		Body:   body,
		Author: author,
		Answer: answer,
	}
}

func (c *comments) Delete(ctx context.Context, id int64) error {
	dbConn := c.dbConn.Connect()
	dbConn.QueryContext(ctx, "Insert Into")

	return nil
}
