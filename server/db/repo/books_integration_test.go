// +build integration

package repo

import (
	"books/server/conf"
	"books/server/db"
	"books/server/db/dbo"
	"books/server/db/query"
	"context"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestBooks(t *testing.T) {
	config := conf.Config{
		Driver: "postgres",
		DBUrl:  "postgresql://postgres:1111@localhost:5432/books_test?sslmode=disable",
	}
	conn := db.NewConn()
	conn.Inject(config)

	books := NewBooks()
	books.Inject(conn)

	dbConn := conn.Connect()
	if _, err := dbConn.ExecContext(context.Background(), query.Schema); err != nil {
		t.Fail()
	}

	//dbConn.Close()

	t.Run("Create", func(t *testing.T) {
		ctx := context.Background()
		book := &dbo.Book{
			Name:          "book2",
			Genre:         "History",
			BookType:      "book",
			PageCount:     12,
			AuthorName:    "AuthorN",
			AuthorSurname: "AuthorS",
			Price:         250,
		}

		err, id := func(ctx context.Context, book *dbo.Book) (error, int64) {
			dbConn := conn.Connect()
			tx, err := dbConn.BeginTx(ctx, nil)
			if err != nil {
				return err, 0
			}

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
				bookID, book.Price).Scan(&unitID)

			if err != nil {
				tx.Rollback()
				return err, unitID
			}

			_, err = tx.ExecContext(ctx, query.DefRating,
				unitID)
			if err != nil {
				tx.Rollback()
				return err, 0
			}

			return tx.Commit(), unitID
		}(ctx, book)
		assert.NoError(t, err)

		createdBook, err := books.GetByID(ctx, id)
		if err != nil {
			t.Fail()
		}

		assert.NoError(t, err)

		assert.Equal(t, createdBook.Price, book.Price)
		//dbConn.Close()
	})

	t.Run("Delete", func(t *testing.T) {
		ctx := context.Background()
		book := &dbo.Book{
			Name:          "book2",
			PageCount:     12,
			Genre:         "History",
			BookType:      "book",
			AuthorSurname: "nnn",
			AuthorName:    "fdvdd",
			Price:         250,
		}

		err, id := func(ctx context.Context, book *dbo.Book) (error, int64) {
			dbConn := conn.Connect()
			tx, err := dbConn.BeginTx(ctx, nil)
			if err != nil {
				return err, 0
			}

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
				bookID, book.Price).Scan(&unitID)

			if err != nil {
				tx.Rollback()
				return err, unitID
			}

			_, err = tx.ExecContext(ctx, query.DefRating,
				unitID)
			if err != nil {
				tx.Rollback()
				return err, 0
			}

			return tx.Commit(), unitID
		}(ctx, book)

		assert.NoError(t, err)

		err = books.Delete(ctx, id)
		if err != nil {
			t.Fail()
		}
		assert.NoError(t, err)

		//_, err = books.GetByID(ctx, book.ID)
		//if err != nil {
		//	t.Fail()
		//}
		//assert.NotEmpty(t, err)
		//assert.Contains(t, err.Error(), "could not find the book")
	})

}

//func test(tFunc func(ctx context.Context, t *testing.T)) func(t *testing.T) {
//	ctx := context.Background()
//	ctx, cancel := context.WithCancel(ctx)
//
//	return func(t *testing.T) {
//		defer cancel()
//
//		tFunc(ctx, t)
//	}
//}
