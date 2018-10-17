package request

type CreateBookReq struct {
	ID            int64  `json:"id"`
	Name          string `json:"name"`
	Genre         string `json:"genre"`
	BookType      string `json:"book_type"`
	PageCount     int    `json:"page_count"`
	AuthorName    string `json:"author_name"`
	AuthorSurname string `json:"author_surname"`
	Price         int    `json:"price"`
}
