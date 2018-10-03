package service

import (
	"books/dbo"
	"books/mocks/mock_repo"
	"context"
	"errors"
	"testing"
)

func TestBookStore_CreateBook(t *testing.T) {
	t.Run("should return error", func(t *testing.T) {
		booksRepoMock := &mock_repo.BooksMock{}
		booksRepoMock.Mock.Create = func(ctx context.Context, book *dbo.Book) error {
			return errors.New("ooops!")
		}

		service := NewBookStore()
		service.Inject(booksRepoMock)

		response := service.CreateBook(context.Background(), CreateBookRequest{})

		if response.Error != "ooops!" {
			t.Fail()
		}
	})
}
