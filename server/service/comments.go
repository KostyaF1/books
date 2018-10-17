package service

import (
	"books/server/db/dbo"
	"books/server/db/repo"
	"books/server/service/request"
	"books/server/service/response"
	"context"
)

type Comment interface {
	Add(ctx context.Context, req request.AddComment) response.AddComment
	AddAnswer(ctx context.Context, req request.AddComment) response.AddComment
	GetAll(ctx context.Context, bookID int64) response.GetComments
	GetByID(ctx context.Context, id int64) response.GetCommID
	Delete(ctx context.Context, req request.DeleteComment) response.DeleteComment
}

type comments struct {
	comments repo.Comments
}

func NewComment() *comments {
	return new(comments)
}

var _ Comment = (*comments)(nil)

func (c *comments) Inject(comm repo.Comments) {
	c.comments = comm
}

func (c *comments) Add(ctx context.Context, req request.AddComment) response.AddComment {
	comment := dbo.Comment{
		BookID: req.BookID,
		Author: req.Author,
		Body:   req.Body,
	}
	if err := c.comments.Add(ctx, comment); err != nil {
		return response.AddComment{
			Error: err.Error(),
		}
	}

	return response.AddComment{
		Error: "nil",
	}
}

func (c *comments) AddAnswer(ctx context.Context, req request.AddComment) response.AddComment {
	comment := dbo.Comment{
		BookID:          req.BookID,
		Author:          req.Author,
		Body:            req.Body,
		CommentParentID: req.CommentParentID,
	}

	var resp response.AddComment

	if err := c.comments.AddAnswer(ctx, comment); err != nil {
		resp.Error = err.Error()
		return resp
	}

	return resp
}

func (c *comments) GetAll(ctx context.Context, bookID int64) response.GetComments {
	comments, err := c.comments.GetAll(ctx, bookID)

	if err != nil {
		return response.GetComments{
			Error: err.Error(),
		}
	}

	return response.GetComments{
		Comments: comments,
	}
}

func (c *comments) GetByID(ctx context.Context, id int64) response.GetCommID {
	comment, err := c.comments.GetByID(ctx, id)

	if err != nil {
		return response.GetCommID{
			Error: err.Error(),
		}
	}
	return response.GetCommID{
		Comm: comment,
	}
}

func (c *comments) Delete(ctx context.Context, req request.DeleteComment) response.DeleteComment {
	if err := c.comments.Delete(ctx, req.ID); err != nil {
		return response.DeleteComment{
			Error: err.Error(),
		}
	}

	return response.DeleteComment{
		Error: "nil",
	}
}
