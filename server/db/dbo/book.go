package dbo

//Book is a model of book/magazine/newspaper in DB
type Book struct {
	ID        int64
	Name      string
	BookType  string
	Genre     string
	PageCount int
	Author    string
}
