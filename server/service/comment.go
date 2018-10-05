package service

import (
	"books/server/db/repo"
	"context"
)

type (
	AddCommentReq struct {
		Body   string
		Father int64
	}

	AddCommentResp struct {
		ID    int64
		Error error
	}
)

type CommentService interface {
	AddComment(ctx context.Context, req AddCommentReq) AddCommentResp
}

type commentService struct {
	addComm repo.Comment
}

func NewAddComment() *commentService {
	return new(commentService)
}

//var _ CommentService = (*commentService)(nil)

func (cs *commentService) Inject(addComm repo.Comment) {
	cs.addComm = addComm
}

/*func (cs *commentService) AddComment(ctx context.Context, req AddCommentReq) AddCommentResp{

}*/
