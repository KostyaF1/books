package service

import (
	"books/server/db/dbo"
	"books/server/db/repo"
	"context"
)

type (
	CreateBookReq struct {
		ID            int64  `json:"id"`
		Name          string `json:"name"`
		Genre         string `json:"genre"`
		BookType      string `json:"book_type"`
		PageCount     int    `json:"page_count"`
		AuthorName    string `json:"author_name"`
		AuthorSurname string `json:"author_surname"`
		Price         int    `json:"price"`
	}
	CreateBookResp struct {
		ID            int64  `json:"id"`
		Name          string `json:"name"`
		Genre         string `json:"genre"`
		BookType      string `json:"book_type"`
		PageCount     int    `json:"page_count"`
		AuthorName    string `json:"author_name"`
		AuthorSurname string `json:"author_surname"`
		Price         int    `json:"price"`
		Error         error  `json:"error"`
	}

	DeleteBookReq struct {
		ID int64 `json:"id"`
	}
	DeleteBookResp struct {
		ID    int64  `json:"id"`
		Name  string `json:"name"`
		Error error  `json:"error"`
	}

	GetBookResp struct {
		AllBooks []*repo.GetBookRepo `json:"all_books"`
	}
	GetBookIDResp struct {
		Book *repo.GetBookIDRepo `json:"book"`
	}
)

type Book interface {
	CreateBook(ctx context.Context, req CreateBookReq) CreateBookResp
	DeleteBook(ctx context.Context, req DeleteBookReq) DeleteBookResp
	GetAllBooks(ctx context.Context) GetBookResp
	GetBookByID(ctx context.Context, id int64) GetBookIDResp
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

func (b *book) CreateBook(ctx context.Context, req CreateBookReq) CreateBookResp {
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
		return CreateBookResp{
			Error: err,
		}
	}
	return CreateBookResp{
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

func (b *book) DeleteBook(ctx context.Context, req DeleteBookReq) DeleteBookResp {
	book := dbo.Book{
		ID: req.ID,
	}

	d_id, d_name, err := b.books.Delete(ctx, book.ID)
	if err != nil {
		return DeleteBookResp{
			Error: err,
		}
	}
	return DeleteBookResp{
		ID:   d_id,
		Name: d_name,
	}
}

func (b *book) GetAllBooks(ctx context.Context) GetBookResp {
	allBooks := b.books.GetAll(ctx)

	return GetBookResp{
		AllBooks: allBooks,
	}
}

func (b *book) GetBookByID(ctx context.Context, id int64) GetBookIDResp {
	Book := b.books.GetByID(ctx, id)

	return GetBookIDResp{
		Book: Book,
	}
}
