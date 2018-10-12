package service

import (
	"books/server/db/dbo"
	"books/server/db/repo"
	"books/server/service/request"
	"books/server/service/response"
	"context"
)

type Book interface {
	CreateBook(ctx context.Context, req request.CreateBookReq) response.CreateBookResp
	DeleteBook(ctx context.Context, req request.DeleteBookReq) response.DeleteBookResp
	GetAllBooks(ctx context.Context) response.GetBookResp
	GetBookByID(ctx context.Context, id int64) response.GetBookIDResp
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

func (b *book) CreateBook(ctx context.Context, req request.CreateBookReq) response.CreateBookResp {
	book := dbo.Book{
		Name:          req.Name,
		Genre:         req.Genre,
		BookType:      req.BookType,
		PageCount:     req.PageCount,
		AuthorName:    req.AuthorName,
		AuthorSurname: req.AuthorSurname,
		Price:         req.Price,
	}

	bookResp, err := b.books.Create(ctx, book)
	if err != nil {
		return response.CreateBookResp{
			Error: err,
		}
	}
	return response.CreateBookResp{
		ID:            bookResp.ID,
		Name:          bookResp.Name,
		Genre:         bookResp.Genre,
		BookType:      bookResp.BookType,
		PageCount:     bookResp.PageCount,
		AuthorName:    bookResp.AuthorName,
		AuthorSurname: bookResp.AuthorSurname,
		Price:         bookResp.Price,
		Error:         err,
	}
}

func (b *book) DeleteBook(ctx context.Context, req request.DeleteBookReq) response.DeleteBookResp {
	book := dbo.Book{
		ID: req.ID,
	}

	d_id, d_name, err := b.books.Delete(ctx, book.ID)
	if err != nil {
		return response.DeleteBookResp{
			Error: err,
		}
	}
	return response.DeleteBookResp{
		ID:   d_id,
		Name: d_name,
	}
}

func (b *book) GetAllBooks(ctx context.Context) response.GetBookResp {
	allBooks, err := b.books.GetAll(ctx)
	if err != nil {
		return response.GetBookResp{
			Error: err,
		}
	}
	return response.GetBookResp{
		AllBooks: allBooks,
	}
}

func (b *book) GetBookByID(ctx context.Context, id int64) response.GetBookIDResp {
	Book := b.books.GetByID(ctx, id)

	return response.GetBookIDResp{
		Book: Book,
	}
}
