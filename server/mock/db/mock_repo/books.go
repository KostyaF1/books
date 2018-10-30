package mock_repo

import (
	"books/server/db"
	"books/server/db/dbo"
	"books/server/db/query"
	"books/server/db/repo"
	"context"
)

type Books struct {
	DBConn db.Connector
	Mock   struct {
		Create  func(ctx context.Context, book *dbo.Book) error
		Delete  func(ctx context.Context, id int64) error
		GetAll  func(ctx context.Context) ([]*dbo.GetBook, error)
		GetByID func(ctx context.Context, bookID int64) (*dbo.GetBookID, error)
	}
}

func NewBooks() *Books {
	return new(Books)
}

func (repo *Books) Inject(conn db.Connector) {
	repo.DBConn = conn
}

func (repo *Books) CreateIntegrationTest(ctx context.Context, book *dbo.Book) (error, int64) {
	dbConn := repo.DBConn.Connect()
	tx, err := dbConn.BeginTx(ctx, nil)

	var bookID int64

	if err = tx.QueryRowContext(ctx, query.CreateBook,
		book.Name, book.PageCount, book.BookType,
	).Scan(&bookID); err != nil {
		tx.Rollback()
		return err, 0
	}
	var authorID int64

	if err = tx.QueryRowContext(ctx, query.CreateAuthor,
		book.AuthorName, book.AuthorSurname).Scan(&authorID); err != nil {
		tx.Rollback()
		return err, 0
	}

	_, err = tx.ExecContext(ctx, query.AuthorBook, bookID, authorID)
	if err != nil {
		tx.Rollback()
		return err, 0
	}

	_, err = tx.ExecContext(ctx, query.GenreBook,
		bookID, book.Genre)
	if err != nil {
		tx.Rollback()
		return err, 0
	}

	var unitID int64

	err = tx.QueryRowContext(ctx, query.StoreUnit,
		bookID, book.Price).Scan(unitID)

	if err != nil {
		tx.Rollback()
		return err, unitID
	}

	return tx.Commit(), unitID
}

func (repo *Books) Create(ctx context.Context, book *dbo.Book) error {
	return repo.Mock.Create(ctx, book)
}

func (repo *Books) Delete(ctx context.Context, id int64) error {
	return repo.Mock.Delete(ctx, id)
}

func (repo *Books) GetAll(ctx context.Context) ([]*dbo.GetBook, error) {
	return repo.Mock.GetAll(ctx)
}

func (repo *Books) GetByID(ctx context.Context, bookID int64) (*dbo.GetBookID, error) {
	return repo.Mock.GetByID(ctx, bookID)
}

var _ repo.Books = (*Books)(nil)
