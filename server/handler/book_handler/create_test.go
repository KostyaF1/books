package book_handler

import (
	"books/server/mock/mock_service"
	"books/server/service/request"
	"books/server/service/response"
	"context"
	"github.com/stretchr/testify/assert"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestCreate_ServeHTTP(t *testing.T) {
	t.Run("check req structure", func(t *testing.T) {

		mock := mock_service.Books{}

		mock.Mock.Create = func(ctx context.Context, req request.CreateBook) response.CreateBook {
			assert.NotEqual(t, req.Name, "")
			assert.NotEqual(t, req.Genre, "")
			assert.NotEqual(t, req.BookType, "")
			assert.NotEqual(t, req.PageCount, 0)
			assert.NotEqual(t, req.AuthorName, "")
			assert.NotEqual(t, req.AuthorSurname, "")
			assert.NotEqual(t, req.Price, 0)

			return response.CreateBook{
				Error: "nil",
			}
		}

		handler := NewCreate()
		handler.Inject(&mock)

		js := `{
			"name":"mock",
			"genre":"criminal",
			"book_type":"book",
			"page_count":300,
			"author_name":"mock",
			"author_surname":"handler",
			"price":200
		}`

		bodyReader := strings.NewReader(js)

		r := httptest.NewRequest("POST", "/book", bodyReader)

		w := httptest.ResponseRecorder{}

		handler.ServeHTTP(&w, r)
	})

	t.Run("check resp", func(t *testing.T) {

		mock := mock_service.Books{}

		mock.Mock.Create = func(ctx context.Context, req request.CreateBook) response.CreateBook {
			return response.CreateBook{
				Error: "nil",
			}
		}

		handler := NewCreate()
		handler.Inject(&mock)

		js := `
		}`

		bodyReader := strings.NewReader(js)

		r := httptest.NewRequest("POST", "/book", bodyReader)

		w := httptest.ResponseRecorder{}

		handler.ServeHTTP(&w, r)

		assert.Equal(t, 400, w.Code)

	})

	t.Run("check service err err", func(t *testing.T) {

		mock := mock_service.Books{}

		mock.Mock.Create = func(ctx context.Context, req request.CreateBook) response.CreateBook {
			return response.CreateBook{
				Error: "nil",
			}
		}

		handler := NewCreate()
		handler.Inject(&mock)

		js := `
		}`

		bodyReader := strings.NewReader(js)

		r := httptest.NewRequest("POST", "/book", bodyReader)

		w := httptest.ResponseRecorder{}

		handler.ServeHTTP(&w, r)

		assert.Equal(t, 400, w.Code)

	})

}
