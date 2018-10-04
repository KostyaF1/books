package repo

import (
	"books/server/db"
	"books/server/db/dbo"
	"books/server/db/dbqueries"
	"context"
	"log"
)

type Create interface {
	Create(ctx context.Context, book dbo.Book) error
}

type Delete interface {
	Delete(ctx context.Context, name string) error
}

type GetAll interface {
	GetAll(ctx context.Context) []*dbo.Book
}

//books...
type books struct {
	dbConn db.Connector
}

var _ Create = (*books)(nil)
var _ Delete = (*books)(nil)
var _ GetAll = (*books)(nil)

func NewBooks() *books {
	return new(books)
}

func (b *books) Inject(conn db.Connector) {
	b.dbConn = conn
}

func (b *books) Create(ctx context.Context, book dbo.Book) error {
	return nil
}

func (b *books) Delete(ctx context.Context, name string) error {
	return nil
}

//GetAllUnits ...
func (b *books) GetAll(ctx context.Context) []*dbo.Book {
	dbConn := b.dbConn.Connect()
	rows, err := dbConn.QueryContext(ctx, dbqueries.GetAllUnitsQuerie)
	if err != nil {
		log.Println("Error when GetAll" + err.Error())
	}

	allBooks := []*dbo.Book{}

	for rows.Next() {
		var (
			bookName        string
			authorLastName  string
			authorFirstName string
			bookType        string
			genre           string
		)
		rows.Scan(&bookName, &authorLastName, &authorFirstName, &bookType, &genre)
		allBooks = append(allBooks, &dbo.Book{
			Name: bookName,
		})
	}
	return allBooks
}
