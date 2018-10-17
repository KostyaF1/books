package query

const DeleteRating = `
DELETE
FROM rating
WHERE store_unit_id = $1`

const DeleteAuthProd = `
DELETE
FROM author_products
WHERE book_product_id IN (SELECT store_units.book_product_id
FROM store_units WHERE store_units.id = $1) RETURNING book_product_id;`

const DeleteProdGenre = `
DELETE
FROM product_genres
WHERE book_product_id IN (SELECT store_units.book_product_id FROM store_units WHERE store_units.id = $1);`

const DeleteComments = `
DELETE
FROM comments
WHERE comments.book_unit_id = $1;
`

const DeleteStoreUnit = `
DELETE
FROM store_units
WHERE id = $1;`

const DeleteBookByID = `DELETE
FROM book_products
WHERE id = $1;
`
