package service

import (
	"books/server/db/dbo"
	"books/server/db/repo"
	"books/server/service/request"
	"books/server/service/response"
	"context"
)

type Comment interface {
	AddComment(ctx context.Context, req request.AddCommentReq) response.AddCommentResp
	AddCommentAnswer(ctx context.Context, req request.AddCommentReq) response.AddCommentResp
	GetComments(ctx context.Context, bookID int64) response.GetCommentsResp
	GetCommByID(ctx context.Context, id int64) response.GetCommIDResp
}

type comment struct {
	comments repo.Comments
}

func NewAddComment() *comment {
	return new(comment)
}

var _ Comment = (*comment)(nil)

func (cs *comment) Inject(comm repo.Comments) {
	cs.comments = comm
}

func (cs *comment) AddComment(ctx context.Context, req request.AddCommentReq) response.AddCommentResp {
	comment := dbo.Comment{
		BookID: req.BookID,
		Author: req.Author,
		Body:   req.Body,
	}
	resp, err := cs.comments.Add(ctx, comment)
	if err != nil {
		return response.AddCommentResp{
			Error: err,
		}
	}

	return response.AddCommentResp{
		ID:     resp.ID,
		Author: resp.Author,
	}
}

func (cs *comment) AddCommentAnswer(ctx context.Context, req request.AddCommentReq) response.AddCommentResp {
	comment := dbo.Comment{
		BookID: req.BookID,
		Author: req.Author,
		Body:   req.Body,
		Father: req.Father,
	}
	resp, err := cs.comments.AddAnswer(ctx, comment)
	if err != nil {
		return response.AddCommentResp{
			Error: err,
		}
	}

	return response.AddCommentResp{
		ID:     resp.ID,
		Author: resp.Author,
		Father: resp.Father,
	}
}

func (cs *comment) GetComments(ctx context.Context, bookID int64) response.GetCommentsResp {
	comments := cs.comments.GetCommentsRepo(ctx, bookID)

	return response.GetCommentsResp{
		Comments: comments,
	}
}

func (c *comment) GetCommByID(ctx context.Context, id int64) response.GetCommIDResp {
	comment := c.comments.GetByID(ctx, id)

	return response.GetCommIDResp{
		Comm: comment,
	}
}
