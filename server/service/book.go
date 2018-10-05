package service

import (
	"books/server/db/dbo"
	"books/server/db/repo"
	"context"
)

type (
	BookReq struct {
		ID   int64  `json:"id"`
		Name string `json:"name"`
	}
	BookResp struct {
		ID       int64       `json:"id"`
		Name     string      `json:"name"`
		AllBooks []*dbo.Book `json:"all_books"`
		Error    error       `json:"error"`
	}
)

type Book interface {
	CreateBook(ctx context.Context, req BookReq) BookResp
	DeleteBook(ctx context.Context, req BookReq) BookResp
	GetAllBooks(ctx context.Context) BookResp
}

//book ...
type book struct {
	books repo.Books
}

func NewBook() *book {
	return new(book)
}

var _ Book = (*book)(nil)

func (b *book) Inject(books repo.Books) {
	b.books = books
}

func (b *book) CreateBook(ctx context.Context, req BookReq) BookResp {
	book := dbo.Book{
		Name: req.Name,
	}
	if err := b.books.Create(ctx, book); err != nil {
		return BookResp{
			Error: err,
		}
	}
	return BookResp{
		ID: book.ID,
	}
}

func (b *book) DeleteBook(ctx context.Context, req BookReq) BookResp {
	book := dbo.Book{
		Name: req.Name,
	}
	if err := b.books.Delete(ctx, book.ID); err != nil {
		return BookResp{
			Error: err,
		}
	}
	return BookResp{
		ID: book.ID,
	}
}

func (b *book) GetAllBooks(ctx context.Context) BookResp {
	allBooks := b.books.GetAll(ctx)
	return BookResp{
		AllBooks: allBooks,
	}
}
