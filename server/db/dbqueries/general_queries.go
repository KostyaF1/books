package dbqueries

//GeneralQuerie ...
type GeneralQuerie struct {
	ID        int
	LastName  string
	FirstName string
	BookName  string
	PageCount int
	Genre     string
	BookType  string
}

//GetAllUnitsQuerie ...
var GetAllUnitsQuerie = `SELECT authors.last_name AS lastName,
authors.first_name AS firstName,
book_products.name AS bookName,
page_count AS pageCount,
(SELECT name AS genreName
 FROM genres,
	  product_genres
 WHERE product_genres.genre_id = genres.id
   AND product_genres.book_product_id = author_products.book_product_id),
product_types.name AS productType
FROM authors,
book_products,
author_products,
product_types
WHERE author_products.book_product_id = book_products.id
AND author_products.author_id = authors.id
AND product_types.id = book_products.type`
