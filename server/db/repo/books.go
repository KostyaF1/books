package repo

import (
	"books/server/db"
	"books/server/db/dbo"
	"books/server/db/query"
	"context"
	"log"
)

type Books interface {
	Create(ctx context.Context, book dbo.Book) error
	Delete(ctx context.Context, id int64) error
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

func (b *books) Create(ctx context.Context, book dbo.Book) error {
	return nil
}

func (b *books) Delete(ctx context.Context, id int64) error {
	dbConn := b.dbConn.Connect()

	_, err := dbConn.ExecContext(ctx, query.DeleteBookByID, id)

	if err != nil {
		return err
	}

	//res.

	return nil
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
			bookName  string
			genre     string
			bookType  string
			pageCount int
			author    string
		)
		rows.Scan(&bookName, &genre, &bookType, &pageCount, &author)
		allBooks = append(allBooks, &dbo.Book{
			Name:      bookName,
			BookType:  bookType,
			Genre:     genre,
			PageCount: pageCount,
			Author:    author,
		})
	}
	return allBooks
}
