package service

import (
	"books/server/db/dbo"
	"books/server/db/repo"
	"context"
)

type GetAllBooksResp struct {
	allBooks []*dbo.Book `json:"all_books"`
}

type GetAllBooks interface {
	GetAllBooks(ctx context.Context) GetAllBooksResp
}

//GeneralService ...
type getAllBooks struct {
	books repo.GetAll
}

func NewGetAllBooks() *getAllBooks {
	return new(getAllBooks)
}

var _ GetAllBooks = (*getAllBooks)(nil)

func (gab *getAllBooks) Inject(books repo.GetAll) {
	gab.books = books
}

func (gab *getAllBooks) GetAllBooks(ctx context.Context) GetAllBooksResp {
	allBooks := gab.books.GetAll(ctx)
	return GetAllBooksResp{
		allBooks: allBooks,
	}
}
