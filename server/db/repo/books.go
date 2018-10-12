package repo

import (
	"books/server/db"
	"books/server/db/dbo"
	"books/server/db/query"
	"context"
	"fmt"
)

type Books interface {
	Create(ctx context.Context, book dbo.Book) (*dbo.CreateBookRepo, error)
	Delete(ctx context.Context, id int64) (int64, string, error)
	GetAll(ctx context.Context) ([]*dbo.GetBookRepo, error)
	GetByID(ctx context.Context, bookID int64) *dbo.GetBookIDRepo
}

//books...
type books struct {
	dbConn db.Connector
}

var _ Books = (*books)(nil)

func NewBooks() *books {
	return new(books)
}

func (b *books) Inject(conn db.Connector) {
	b.dbConn = conn
}

func (b *books) Create(ctx context.Context, book dbo.Book) (*dbo.CreateBookRepo, error) {
	dbConn := b.dbConn.Connect()
	tx, err := dbConn.BeginTx(ctx, nil)

	if err != nil {
		return nil, err
	}

	var bookID int64
	err = tx.QueryRowContext(
		ctx,
		query.CreateBook,
		book.Name, book.PageCount, book.BookType,
	).Scan(&bookID)

	if err != nil {
		tx.Rollback()
		return nil, err
	}

	var authorID int64

	err = tx.QueryRowContext(ctx,
		query.CreateAuthor,
		book.AuthorName, book.AuthorSurname).Scan(&authorID)

	if err != nil {
		tx.Rollback()
		return nil, err
	}

	_, err = tx.ExecContext(ctx,
		query.AuthorBook, bookID, authorID)
	if err != nil {
		tx.Rollback()
		return nil, err
	}

	_, err = tx.ExecContext(ctx,
		query.GenreBook,
		bookID, book.Genre)

	if err != nil {
		tx.Rollback()
		return nil, err
	}

	var storeUnit int64

	err = tx.QueryRowContext(ctx,
		query.StoreUnit,
		bookID, book.Price).Scan(&storeUnit)

	if err != nil {
		tx.Rollback()
		return nil, err
	}

	return &dbo.CreateBookRepo{
		ID:            storeUnit,
		Name:          book.Name,
		Genre:         book.Genre,
		BookType:      book.BookType,
		PageCount:     book.PageCount,
		AuthorName:    book.AuthorName,
		AuthorSurname: book.AuthorSurname,
	}, tx.Commit()

}

func (b *books) Delete(ctx context.Context, id int64) (int64, string, error) {
	dbConn := b.dbConn.Connect()

	row, err := dbConn.QueryContext(ctx, "SELECT name FROM book_products WHERE id=$1", id)
	if err != nil {
		return 0, "", err
	}
	row.Next()
	var name string
	row.Scan(&name)

	_, err = dbConn.ExecContext(ctx, query.DeleteBookByID, id)

	if err != nil {
		return 0, "", err
	}

	//res.

	return id, name, err
}

//GetAllUnits ...
func (b *books) GetAll(ctx context.Context) ([]*dbo.GetBookRepo, error) {
	dbConn := b.dbConn.Connect()
	rows, err := dbConn.QueryContext(ctx, query.GetAllUnitsQuerie)
	if err != nil {
		return nil, err
	}

	var allBooks []*dbo.GetBookRepo

	for rows.Next() {
		var book dbo.GetBookRepo

		if err = rows.Scan(
			&book.ID,
			&book.Name,
			&book.Genre,
			&book.BookType,
			&book.PageCount,
			&book.Author,
			&book.Price,
			&book.Value,
		); err != nil {
			return nil, err
		}

		allBooks = append(allBooks, &book)
	}
	return allBooks, nil
}

func (b *books) GetByID(ctx context.Context, bookID int64) *dbo.GetBookIDRepo {
	dbConn := b.dbConn.Connect()
	rows, err := dbConn.QueryContext(ctx, query.GetBookByID, bookID)
	if err != nil {
		return &dbo.GetBookIDRepo{
			Error: err,
		}
	}

	rows.Next()

	var book dbo.GetBookIDRepo

	if err = rows.Scan(
		&book.ID,
		&book.BookName,
		&book.Genre,
		&book.BookType,
		&book.PageCount,
		&book.Author,
		&book.Price,
		//&book.Comments,
		&book.Value,
	); err != nil {
		fmt.Println(err)
		return &dbo.GetBookIDRepo{
			Error: err,
		}
	}

	return &book
}
