package mock_service

import (
	"books/server/service"
	"books/server/service/request"
	"books/server/service/response"
	"context"
)

type Books struct {
	Mock struct {
		Create  func(ctx context.Context, req request.CreateBook) response.CreateBook
		Delete  func(ctx context.Context, req request.DeleteBook) response.DeleteBook
		GetAll  func(ctx context.Context) response.GetBook
		GetByID func(ctx context.Context, id int64) response.GetBookID
	}
}

func (b *Books) Create(ctx context.Context, req request.CreateBook) response.CreateBook {
	return b.Mock.Create(ctx, req)
}

func (b *Books) Delete(ctx context.Context, req request.DeleteBook) response.DeleteBook {
	return b.Mock.Delete(ctx, req)
}

func (b *Books) GetAll(ctx context.Context) response.GetBook {
	return b.Mock.GetAll(ctx)
}

func (b *Books) GetByID(ctx context.Context, id int64) response.GetBookID {
	return b.Mock.GetByID(ctx, id)
}

var _ service.Book = (*Books)(nil)
