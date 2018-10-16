package service

import (
	"books/server/db/dbo"
	"books/server/mock/db/mock_repo"
	"books/server/service/request"
	"context"
	"errors"
	"testing"
)

//TestBook_CreateBook...
func TestBook_CreateBook(t *testing.T) {

	t.Run("check name and returned error", func(t *testing.T) {
		const (
			expectedName = "some name"
		)

		mock := mock_repo.Books{}

		mock.Mock.Create = func(ctx context.Context, book *dbo.Book) error {
			if book.Name != expectedName {
				t.Fail()
			}

			return errors.New("some error")
		}

		book := NewBook()
		book.Inject(&mock)

		resp := book.CreateBook(context.Background(), request.CreateBookReq{
			Name: expectedName,
		})

		if resp.Error == "" {
			t.Fail()
		}
	})

	t.Run("check no panic on nil error", func(t *testing.T) {
		mock := &mock_repo.Books{}

		mock.Mock.Create = func(ctx context.Context, book *dbo.Book) error {
			return nil
		}

		book := NewBook()
		book.Inject(mock)

		resp := book.CreateBook(context.Background(), request.CreateBookReq{})

		if resp.Error != "" {
			t.Fail()
		}
	})

}

//TestBook_DeleteBook...
func TestBook_DeleteBook(t *testing.T) {

	t.Run("check id and returned error", func(t *testing.T) {

		mock := mock_repo.Books{}

		mock.Mock.Delete = func(ctx context.Context, id int64) error {
			if id == 0 {
				t.Fail()
			}

			return errors.New("some error")
		}

		book := NewBook()
		book.Inject(&mock)

		resp := book.DeleteBook(context.Background(), request.DeleteBookReq{
			ID: 5,
		})

		if resp.Error == "" {
			t.Fail()
		}
	})

	t.Run("check no panic on nil error", func(t *testing.T) {
		mock := &mock_repo.Books{}

		mock.Mock.Delete = func(ctx context.Context, id int64) error {
			return nil
		}

		book := NewBook()
		book.Inject(mock)

		resp := book.DeleteBook(context.Background(), request.DeleteBookReq{
			ID: 0,
		})

		if resp.Error != "" {
			t.Fail()
		}
	})

}

//TestBook_GetAll ...
func TestBook_GetAll(t *testing.T) {

	t.Run("check Value and returned error", func(t *testing.T) {

		mock := mock_repo.Books{}

		mock.Mock.GetAll = func(ctx context.Context) ([]*dbo.GetBookRepo, error) {
			var allBooks []*dbo.GetBookRepo
			allBooks = append(allBooks, &dbo.GetBookRepo{
				Name:     "lknk",
				BookType: "book",
			})
			return allBooks, nil
		}
		book := NewBook()
		book.Inject(&mock)

		resp := book.GetAllBooks(context.Background())

		if resp.Error != "" {
			t.Fail()
		}
		for value := range resp.AllBooks {
			if resp.AllBooks[value].Name == "" {
				t.Fail()
			}
		}

	})

	t.Run("check no panic on nil error", func(t *testing.T) {
		mock := &mock_repo.Books{}

		mock.Mock.GetAll = func(ctx context.Context) ([]*dbo.GetBookRepo, error) {

			return nil, errors.New("ERRRORRR")
		}

		book := NewBook()
		book.Inject(mock)

		resp := book.GetAllBooks(context.Background())

		if resp.Error != "" {
			t.Fail()
		}
	})

}

//TestBook_GetBookByID...
func TestBook_GetBookByID(t *testing.T) {

	t.Run("check value and returned error", func(t *testing.T) {

		mock := mock_repo.Books{}

		mock.Mock.GetByID = func(ctx context.Context, id int64) (*dbo.GetBookIDRepo, error) {

			value := &dbo.GetBookIDRepo{
				//BookName: "lknk",
				BookType: "book",
			}
			return value, nil
		}
		book := NewBook()
		book.Inject(&mock)

		resp := book.GetBookByID(context.Background(), 5)
		if resp.Book.BookName == "" {
			t.Fail()
		}

	})

	t.Run("check no panic on nil error", func(t *testing.T) {
		mock := &mock_repo.Books{}

		mock.Mock.GetByID = func(ctx context.Context, id int64) (*dbo.GetBookIDRepo, error) {
			if id == 0 {
				return nil, errors.New("some error")
			}

			return nil, nil
		}

		book := NewBook()
		book.Inject(mock)

		resp := book.GetBookByID(context.Background(), 5)

		if resp.Error != "" {
			t.Fail()
		}
	})

}
