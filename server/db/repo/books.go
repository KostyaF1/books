package repo

import (
	"books/server/db"
	"books/server/db/dbo"
	"books/server/db/query"
	"context"
	"log"
)

type Books interface {
	Create(ctx context.Context, book dbo.Book) (int64, error)
	Delete(ctx context.Context, id int64) (int64, string, error)
	GetAll(ctx context.Context) []*dbo.Book
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

func (b *books) Create(ctx context.Context, book dbo.Book) (int64, error) {
	dbConn := b.dbConn.Connect()
	tx, err := dbConn.BeginTx(ctx, nil)

	if err != nil {
		return 0, err
	}

	var bookID int64
	var typeID int64 = 1
	err = tx.QueryRowContext(
		ctx,
		"INSERT INTO book_products (name, page_count, type) VALUES ($1, $2, $3) RETURNING id",
		book.Name, book.PageCount, typeID,
	).Scan(&bookID)

	if err != nil {
		tx.Rollback()
		return 0, err
	}

	var authorID int64

	err = tx.QueryRowContext(ctx,
		`
			INSERT INTO authors (first_name, last_name)
			VALUES ($1, $2) RETURNING id;`, book.AuthorName, book.AuthorSurname).Scan(&authorID)

	if err != nil {
		tx.Rollback()
		return 0, err
	}

	_, err = tx.ExecContext(ctx,
		`INSERT INTO author_products (book_product_id, author_id)
							VALUES ($1, $2) RETURNING id;`, bookID, authorID)
	if err != nil {
		tx.Rollback()
		return 0, err
	}

	return authorID, tx.Commit()

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
func (b *books) GetAll(ctx context.Context) []*dbo.Book {
	dbConn := b.dbConn.Connect()
	rows, err := dbConn.QueryContext(ctx, query.GetAllUnitsQuerie)
	if err != nil {
		log.Println("Error when GetAll" + err.Error())
	}

	var allBooks []*dbo.Book

	for rows.Next() {
		var (
			id        int64
			bookName  string
			genre     string
			bookType  string
			pageCount int
			author    string
		)
		rows.Scan(&id, &bookName, &genre, &bookType, &pageCount, &author)
		allBooks = append(allBooks, &dbo.Book{
			ID:        id,
			Name:      bookName,
			BookType:  bookType,
			Genre:     genre,
			PageCount: pageCount,
			Author:    author,
		})
	}
	return allBooks
}
