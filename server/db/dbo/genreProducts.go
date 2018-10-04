package dbo

//GenreProducts is a rilation "many-to-many" between BookProducts and Genres
type GenreProducts struct {
	id            uint
	bookProductID uint
	genreID       uint
}
