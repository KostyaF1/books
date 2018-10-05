package dbqueries

//GeneralQuerie ...
type GeneralQuerie struct {
	BookName  string
	Genre     string
	BookType  string
	PageCount int
	Author    string
}

//GetAllUnitsQuerie ...
var GetAllUnitsQuerie = `
SELECT book_products.name                                                                           AS bookName,
       array(SELECT genres.name
             FROM genres
             WHERE genres.id IN
                   (SELECT product_genres.genre_id
                    FROM product_genres
                    WHERE product_genres.book_product_id = book_products.id)) AS genre,
    product_types.name AS productType,
       book_products.page_count                                                                     AS pageCount,
       array(SELECT authors.first_name || ' ' || authors.last_name
             FROM authors
             WHERE authors.id IN (SELECT author_products.author_id
                                  FROM author_products
                                  WHERE author_products.book_product_id = book_products.id)) AS author
FROM book_products
INNER JOIN product_types  on book_products.type = product_types.id;`

//GetAllUnitsQuerie1 ...
var GetAllUnitsQuerie1 = `
SELECT authors.last_name AS authorName,
       authors.first_name AS authorSurname,
       book_products.name AS bookName,
       page_count AS pageCount,
       (SELECT name
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
  AND product_types.id = book_products.type;`
