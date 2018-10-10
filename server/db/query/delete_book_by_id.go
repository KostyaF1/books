package query

//GeneralQuerie ...
type GeneralQuerie struct {
	BookName  string
	Genre     string
	BookType  string
	PageCount int
	Author    string
}

//GetAllUnitsQuerie ...

const DeleteBookByID = `
DELETE FROM book_products
WHERE id = $1;`
