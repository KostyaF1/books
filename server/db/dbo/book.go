package dbo

//Book is a model of book/magazine/newspaper in DB
type Book struct {
	ID            int64  `json:"id"`
	Name          string `json:"name"`
	Genre         string `json:"genre"`
	BookType      string `json:"book_type"`
	PageCount     int    `json:"page_count"`
	AuthorName    string `json:"author_name"`
	AuthorSurname string `json:"author_surname"`
	Author        string `json:"author"`
	Price         int    `json:"price"`
}

type (
	CreateBookRepo struct {
		ID            int64  `json:"id"`
		Name          string `json:"name"`
		Genre         string `json:"genre"`
		BookType      string `json:"book_type"`
		PageCount     int    `json:"page_count"`
		AuthorName    string `json:"author_name"`
		AuthorSurname string `json:"author_surname"`
		Price         int    `json:"price"`
	}

	GetBookRepo struct {
		ID        int64   `json:"id"`
		Name      string  `json:"name"`
		Genre     string  `json:"genre"`
		BookType  string  `json:"book_type"`
		PageCount int     `json:"page_count"`
		Author    string  `json:"author"`
		Price     int     `json:"price"`
		Value     float64 `json:"value"`
		Error     error   `json:"error"`
	}
	GetBookIDRepo struct {
		ID        int64   `json:"id"`
		BookName  string  `json:"book_name"`
		Genre     string  `json:"genre"`
		BookType  string  `json:"book_type"`
		PageCount int     `json:"page_count"`
		Author    string  `json:"author"`
		Price     int     `json:"price"`
		Value     float64 `json:"value"`
		Comments  string  `json:"comments"`
		Error     error   `json:"error"`
	}
)
