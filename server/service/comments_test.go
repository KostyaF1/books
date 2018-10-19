package service

import (
	"books/server/db/dbo"
	"books/server/mock/db/mock_repo"
	"books/server/service/request"
	"context"
	"github.com/stretchr/testify/assert"
	"testing"
)

//
func TestComments_Add(t *testing.T) {
	t.Run("AddComment check on CommentBody and returned error", func(t *testing.T) {

		mock := mock_repo.Comments{}

		mock.Mock.Add = func(ctx context.Context, comment dbo.Comment) error {
			assert.NotEmpty(t, comment.Body)

			return nil
		}

		comment := NewComment()
		comment.Inject(&mock)

		resp := comment.Add(context.Background(), request.AddComment{
			Body: "df",
		})

		assert.Empty(t, resp.Error)
	})

	t.Run("check no panic on nil error", func(t *testing.T) {

		mock := mock_repo.Comments{}

		mock.Mock.Add = func(ctx context.Context, comment dbo.Comment) error {
			return nil
		}

		comment := NewComment()
		comment.Inject(&mock)

		resp := comment.Add(context.Background(), request.AddComment{})

		assert.Empty(t, resp.Error)
	})
}

//
func TestComments_AddAnswer(t *testing.T) {

}

//
func TestComments_GetAll(t *testing.T) {
}

//
func TestComments_GetByID(t *testing.T) {
}

//
func TestComments_Delete(t *testing.T) {
}
