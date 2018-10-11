package repo

import (
	"books/server/db"
	"books/server/db/dbo"
	"books/server/db/query"
	"context"
	"log"
)

type (
	CreateBookRepo struct {
		ID            int64  `json:"id"`
		Name          string `json:"name"`
		Genre         string `json:"genre"`
		BookType      string `json:"book_type"`
		PageCount     int    `json:"page_count"`
		AuthorName    string `json:"author_name"`
		AuthorSurname string `json:"author_surname"`
	}

	GetBookRepo struct {
		ID        int64  `json:"id"`
		Name      string `json:"name"`
		Genre     string `json:"genre"`
		BookType  string `json:"book_type"`
		PageCount int    `json:"page_count"`
		Author    string `json:"author"`
	}
)

type Books interface {
	Create(ctx context.Context, book dbo.Book) (*CreateBookRepo, error)
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

func (b *books) Create(ctx context.Context, book dbo.Book) (*CreateBookRepo, error) {
	dbConn := b.dbConn.Connect()
	tx, err := dbConn.BeginTx(ctx, nil)

	if err != nil {
		return nil, err
	}

	var bookID int64
	//var typeID int64 = 1
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

	return &CreateBookRepo{
		ID:            bookID,
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
func (b *books) GetAll(ctx context.Context) []*GetBookRepo {
	dbConn := b.dbConn.Connect()
	rows, err := dbConn.QueryContext(ctx, query.GetAllUnitsQuerie)
	if err != nil {
		log.Println("Error when GetAll" + err.Error())
	}

	var allBooks []*GetBookRepo

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
		allBooks = append(allBooks, &GetBookRepo{
			ID:        id,
			Name:      bookName,
			Genre:     genre,
			BookType:  bookType,
			PageCount: pageCount,
			Author:    author,
		})
	}
	return allBooks
}
