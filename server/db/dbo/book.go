package dbo

import "github.com/lib/pq"

//Book is a model of book/magazine/newspaper in DB
type Book struct {
	ID            int64  `json:"id"`
	Name          string `json:"name"`
	Genre         string `json:"genre"`
	BookType      string `json:"book_type"`
	PageCount     int    `json:"page_count"`
	AuthorName    string `json:"author_name"`
	AuthorSurname string `json:"author_surname"`
	Price         int    `json:"price"`
}

type (
	GetBook struct {
		ID        int64          `json:"id"`
		Name      string         `json:"name"`
		Genre     pq.StringArray `json:"genre"`
		BookType  string         `json:"book_type"`
		PageCount int            `json:"page_count"`
		Author    string         `json:"author"`
		Price     int            `json:"price"`
		Value     float64        `json:"value"`
	}
	GetBookID struct {
		ID        int64   `json:"id"`
		BookName  string  `json:"book_name"`
		Genre     string  `json:"genre"`
		BookType  string  `json:"book_type"`
		PageCount int     `json:"page_count"`
		Author    string  `json:"author"`
		Price     int     `json:"price"`
		Value     float64 `json:"value"`
		Comments  string  `json:"comments"`
	}
)
