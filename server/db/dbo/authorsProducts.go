package dbo

//AuthorProducts is a relation "many-to-many" between author and bookProducts
type AuthorProducts struct {
	id            uint
	bookProductID uint
	authorID      uint
}
