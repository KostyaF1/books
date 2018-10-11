package service

import (
	"books/server/db/dbo"
	"books/server/db/repo"
	"context"
)

type (
	AddCommentReq struct {
		BookID int64  `json:"book_id"`
		Body   string `json:"body"`
		Author string `json:"author"`
		Father int64  `json:"father"`
	}

	AddCommentResp struct {
		ID     int64  `json:"id"`
		Author string `json:"author"`
		Father int64  `json:"father"`
		Error  error
	}

	GetCommentsResp struct {
		Comments []*repo.GetCommentsRepo `json:"comments"`
	}
	GetCommIDResp struct {
		Comm *repo.GetCommIDRepo `json:"comm"`
	}
)

type Comment interface {
	AddComment(ctx context.Context, req AddCommentReq) AddCommentResp
	AddCommentAnswer(ctx context.Context, req AddCommentReq) AddCommentResp
	GetComments(ctx context.Context, bookID int64) GetCommentsResp
	GetCommByID(ctx context.Context, id int64) GetCommIDResp
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

func (cs *comment) AddComment(ctx context.Context, req AddCommentReq) AddCommentResp {
	comment := dbo.Comment{
		BookID: req.BookID,
		Author: req.Author,
		Body:   req.Body,
	}
	resp, err := cs.comments.Add(ctx, comment)
	if err != nil {
		return AddCommentResp{
			Error: err,
		}
	}

	return AddCommentResp{
		ID:     resp.ID,
		Author: resp.Author,
	}
}

func (cs *comment) AddCommentAnswer(ctx context.Context, req AddCommentReq) AddCommentResp {
	comment := dbo.Comment{
		BookID: req.BookID,
		Author: req.Author,
		Body:   req.Body,
		Father: req.Father,
	}
	resp, err := cs.comments.AddAnswer(ctx, comment)
	if err != nil {
		return AddCommentResp{
			Error: err,
		}
	}

	return AddCommentResp{
		ID:     resp.ID,
		Author: resp.Author,
		Father: resp.Father,
	}
}

func (cs *comment) GetComments(ctx context.Context, bookID int64) GetCommentsResp {
	comments := cs.comments.GetCommentsRepo(ctx, bookID)

	return GetCommentsResp{
		Comments: comments,
	}
}

func (c *comment) GetCommByID(ctx context.Context, id int64) GetCommIDResp {
	comment := c.comments.GetByID(ctx, id)

	return GetCommIDResp{
		Comm: comment,
	}
}
