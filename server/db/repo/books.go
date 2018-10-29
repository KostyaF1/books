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
	GetAll(ctx context.Context) ([]*dbo.GetBook, error)
	GetByID(ctx context.Context, bookID int64) (*dbo.GetBookID, error)
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

	if err = tx.QueryRowContext(ctx, query.CreateBook,
		book.Name, book.PageCount, book.BookType,
	).Scan(&bookID); err != nil {
		tx.Rollback()
		return err
	}

	var authorID int64

	if err = tx.QueryRowContext(ctx, query.CreateAuthor,
		book.AuthorName, book.AuthorSurname).Scan(&authorID); err != nil {
		tx.Rollback()
		return err
	}

	_, err = tx.ExecContext(ctx, query.AuthorBook, bookID, authorID)
	if err != nil {
		tx.Rollback()
		return err
	}

	_, err = tx.ExecContext(ctx, query.GenreBook,
		bookID, book.Genre)
	if err != nil {
		tx.Rollback()
		return err
	}

	_, err = tx.ExecContext(ctx, query.StoreUnit,
		bookID, book.Price)

	if err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit()
}

// error
func (b *books) Delete(ctx context.Context, id int64) error {
	dbConn := b.dbConn.Connect()
	tx, err := dbConn.BeginTx(ctx, nil)

	if err != nil {
		tx.Rollback()
		return err
	}

	_, err = tx.ExecContext(ctx, query.DeleteRating, id)

	if err != nil {
		tx.Rollback()
		return err
	}

	var unitID int64

	err = tx.QueryRowContext(ctx, query.DeleteAuthProd, id).Scan(&unitID)

	if err != nil {
		tx.Rollback()
		return err
	}

	_, err = tx.ExecContext(ctx, query.DeleteProdGenre, id)

	if err != nil {
		tx.Rollback()
		return err
	}

	_, err = tx.ExecContext(ctx, query.DeleteComments, id)

	if err != nil {
		tx.Rollback()
		return err
	}

	_, err = tx.ExecContext(ctx, query.DeleteStoreUnit, id)

	if err != nil {
		tx.Rollback()
		return err
	}

	_, err = tx.ExecContext(ctx, query.DeleteBookByID, unitID)

	if err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit()
}

// naming
func (b *books) GetAll(ctx context.Context) ([]*dbo.GetBook, error) {
	dbConn := b.dbConn.Connect()
	rows, err := dbConn.QueryContext(ctx, query.GetBooks1)
	if err != nil {
		return nil, err
	}

	var allBooks []*dbo.GetBook

	for rows.Next() {
		var book dbo.GetBook

		if err = rows.Scan(
			//&book.ID,
			&book.Name,
			//&book.Genre,
			//&book.BookType,
			//&book.PageCount,
			//&book.Author,
			//&book.Price,
			//&book.Value,
		); err != nil {
			return nil, err
		}

		allBooks = append(allBooks, &book)
	}

	return allBooks, nil
}

// return err
func (b *books) GetByID(ctx context.Context, bookID int64) (*dbo.GetBookID, error) {
	dbConn := b.dbConn.Connect()

	rows, err := dbConn.QueryContext(ctx, query.GetBookByID, bookID)
	if err != nil {
		return nil, err
	}

	rows.Next()

	var book dbo.GetBookID

	if err = rows.Scan(
		&book.ID,
		&book.BookName,
		&book.Genre,
		&book.BookType,
		&book.PageCount,
		&book.Author,
		&book.Price,
		&book.Value,
		&book.Comments,
	); err != nil {
		return nil, err
	}

	return &book, nil
}
