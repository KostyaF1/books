package repo

import (
	"books/server/db"
	"books/server/db/dbo"
	"books/server/db/query"
	"context"
)

type Books interface {
	Create(ctx context.Context, book *dbo.Book) error
	Delete(ctx context.Context, id int64) error
	GetAll(ctx context.Context) ([]*dbo.GetBookRepo, error)
	GetByID(ctx context.Context, bookID int64) (*dbo.GetBookIDRepo, error)
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

// only error
func (b *books) Create(ctx context.Context, book *dbo.Book) error {
	dbConn := b.dbConn.Connect()
	tx, err := dbConn.BeginTx(ctx, nil)

	var bookID int64
	err = tx.QueryRowContext(
		ctx,
		query.CreateBook,
		book.Name, book.PageCount, book.BookType,
	).Scan(&bookID)

	var authorID int64

	err = tx.QueryRowContext(ctx,
		query.CreateAuthor,
		book.AuthorName, book.AuthorSurname).Scan(&authorID)

	_, err = tx.ExecContext(ctx,
		query.AuthorBook, bookID, authorID)
	if err != nil {
		tx.Rollback()
		return err
	}

	_, err = tx.ExecContext(ctx,
		query.GenreBook,
		bookID, book.Genre)

	_, err = tx.ExecContext(ctx,
		query.StoreUnit,
		bookID, book.Price)

	if err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit()

}

// named parameters
// error
func (b *books) Delete(ctx context.Context, id int64) error {
	dbConn := b.dbConn.Connect()

	row, err := dbConn.QueryContext(ctx, "SELECT name FROM book_products WHERE id=$1", id)
	row.Next()
	var name string
	row.Scan(&name)

	_, err = dbConn.ExecContext(ctx, query.DeleteBookByID, id)

	if err != nil {
		return err
	}

	return nil
}

// naming
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
			//&book.Name,
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

// return err
func (b *books) GetByID(ctx context.Context, bookID int64) (*dbo.GetBookIDRepo, error) {
	dbConn := b.dbConn.Connect()
	rows, err := dbConn.QueryContext(ctx, query.GetBookByID, bookID)
	//if err != nil {
	//	return nil, err
	//}

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
		&book.Comments,
		&book.Value,
	); err != nil {
		return nil, err
	}

	return &book, nil
}
