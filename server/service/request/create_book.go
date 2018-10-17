package request

type CreateBook struct {
	Name          string `json:"name"`
	Genre         string `json:"genre"`
	BookType      string `json:"book_type"`
	PageCount     int    `json:"page_count"`
	AuthorName    string `json:"author_name"`
	AuthorSurname string `json:"author_surname"`
	Price         int    `json:"price"`
}
