package service

import (
	"books/server/db/dbo"
	"books/server/db/repo"
	"books/server/service/request"
	"books/server/service/response"
	"context"
)

type Book interface {
	Create(ctx context.Context, req request.CreateBook) response.CreateBook
	Delete(ctx context.Context, req request.DeleteBook) response.DeleteBook
	GetAll(ctx context.Context) response.GetBook
	GetByID(ctx context.Context, id int64) response.GetBookID
}

//book ...
type books struct {
	books repo.Books
}

func NewBook() *books {
	return new(books)
}

var _ Book = (*books)(nil)

func (b *books) Inject(books repo.Books) {
	b.books = books
}

func (b *books) Create(ctx context.Context, req request.CreateBook) response.CreateBook {
	respErr := response.CreateBook{
		Error: "Empty field",
	}
	switch {
	case req.Name == "":
		return respErr
	case req.Genre == "":
		return respErr
	case req.BookType == "":
		return respErr
	case req.PageCount == 0:
		return respErr
	case req.AuthorName == "":
		return respErr
	case req.AuthorSurname == "":
		return respErr
	case req.Price == 0:
		return respErr
	default:
		book := dbo.Book{
			Name:          req.Name,
			Genre:         req.Genre,
			BookType:      req.BookType,
			PageCount:     req.PageCount,
			AuthorName:    req.AuthorName,
			AuthorSurname: req.AuthorSurname,
			Price:         req.Price,
		}
		if err := b.books.Create(ctx, &book); err != nil {
			return response.CreateBook{
				Error: err.Error(),
			}
		}
	}

	return response.CreateBook{
		Error: "",
	}
}

func (b *books) Delete(ctx context.Context, req request.DeleteBook) response.DeleteBook {
	if err := b.books.Delete(ctx, req.ID); err != nil {
		return response.DeleteBook{
			Error: err.Error(),
		}
	}

	return response.DeleteBook{
		Error: "",
	}
}

func (b *books) GetAll(ctx context.Context) response.GetBook {
	allBooks, err := b.books.GetAll(ctx)
	if err != nil {
		return response.GetBook{
			Error: err.Error(),
		}
	}
	return response.GetBook{
		AllBooks: allBooks,
	}
}

func (b *books) GetByID(ctx context.Context, id int64) response.GetBookID {
	book, err := b.books.GetByID(ctx, id)

	if err != nil {
		return response.GetBookID{
			Error: err.Error(),
		}
	}
	return response.GetBookID{
		Book: book,
	}

}
