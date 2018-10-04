package dbo

//Authors is a model of a book's author
type Authors struct {
	ID        uint   `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
}
