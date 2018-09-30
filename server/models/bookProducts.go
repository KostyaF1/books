package models

//BookProducts is a model of book/magazine/newspaper in DB
type BookProducts struct {
	id              uint
	name            string
	pageCount       uint
	bookProductType uint
}
