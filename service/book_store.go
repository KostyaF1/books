package service

import (
	"books/dbo"
	"books/repo"
	"context"
)

type (
	CreateBookRequest struct {
		Name string `json:"name"`
	}

	CreateBookResponse struct {
		ID    int    `json:"id"`
		Error string `json:"error, omitempty"`
	}
)

type BookStore interface {
	CreateBook(ctx context.Context, req CreateBookRequest) CreateBookResponse
}

var _ BookStore = (*bookStore)(nil)

func NewBookStore() *bookStore {
	return new(bookStore)
}

func (service *bookStore) Inject(books repo.Books) {
	service.books = books
}

type bookStore struct {
	books repo.Books
}

func (service *bookStore) CreateBook(ctx context.Context, req CreateBookRequest) CreateBookResponse {
	book := &dbo.Book{
		Name: req.Name,
	}

	if err := service.books.Create(ctx, book); err != nil {
		return CreateBookResponse{
			Error: err.Error(),
		}
	}

	return CreateBookResponse{
		ID: book.ID,
	}
}
