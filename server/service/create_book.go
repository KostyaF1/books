package service

import (
	"books/server/db/dbo"
	"books/server/db/repo"
	"context"
)

type (
	CreateBookReq struct {
		ID   int64  `json:"id"`
		Name string `json:"name"`
	}
	CreateBookResp struct {
		ID    int64  `json:"id"`
		Name  string `json:"name"`
		Error error  `json:"error"`
	}
)

type CreateBook interface {
	CreateBook(ctx context.Context, req CreateBookReq) CreateBookResp
}

//createBook ...
type createBook struct {
	books repo.Create
}

func NewCreateBook() *createBook {
	return new(createBook)
}

var _ CreateBook = (*createBook)(nil)

func (cb *createBook) Inject(books repo.Create) {
	cb.books = books
}

func (cb *createBook) CreateBook(ctx context.Context, req CreateBookReq) CreateBookResp {
	book := dbo.Book{
		Name: req.Name,
	}
	if err := cb.books.Create(ctx, book); err != nil {
		return CreateBookResp{
			Error: err,
		}
	}
	return CreateBookResp{
		ID: book.ID,
	}
}
