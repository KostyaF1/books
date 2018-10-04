package service

import (
	"books/server/db/dbo"
	"books/server/db/repo"
	"context"
)

type (
	DeleteBookReq struct {
		ID   int64  `json:"id"`
		Name string `json:"name"`
	}
	DeleteBookResp struct {
		ID    int64  `json:"id"`
		Name  string `json:"name"`
		Error error  `json:"error"`
	}
)

type DeleteBook interface {
	DeleteBook(ctx context.Context, req DeleteBookReq) DeleteBookResp
}

//deleteBook ...
type deleteBook struct {
	books repo.Delete
}

func NewDeleteBook() *deleteBook {
	return new(deleteBook)
}

var _ DeleteBook = (*deleteBook)(nil)

func (db *deleteBook) Inject(books repo.Delete) {
	db.books = books
}

func (db *deleteBook) DeleteBook(ctx context.Context, req DeleteBookReq) DeleteBookResp {
	book := dbo.Book{
		Name: req.Name,
	}
	if err := db.books.Delete(ctx, book.Name); err != nil {
		return DeleteBookResp{
			Error: err,
		}
	}
	return DeleteBookResp{
		ID: book.ID,
	}
}
