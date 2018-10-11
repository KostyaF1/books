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
		Father string `json:"father"`
		Error  error
	}
)

type Comment interface {
	AddComment(ctx context.Context, req AddCommentReq) AddCommentResp
}

type comment struct {
	addComm repo.Comment
}

func NewAddComment() *comment {
	return new(comment)
}

var _ Comment = (*comment)(nil)

func (cs *comment) Inject(addComm repo.Comment) {
	cs.addComm = addComm
}

func (cs *comment) AddComment(ctx context.Context, req AddCommentReq) AddCommentResp {
	comment := dbo.Comment{
		BookID: req.BookID,
		Author: req.Author,
		Body:   req.Body,
	}
	resp, err := cs.addComm.Add(ctx, comment)
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
