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
