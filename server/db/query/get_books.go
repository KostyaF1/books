package query

const GetBooks = `
SELECT store_units.id ,
      book_products.name AS book_name,
       array(SELECT genres.name
             FROM genres
             WHERE genres.id IN
                   (SELECT product_genres.genre_id
                    FROM product_genres
                    WHERE product_genres.book_product_id = store_units.book_product_id)) AS genre,
       product_types.name AS product_type,
       book_products.page_count,
       array(SELECT authors.first_name || ' ' || authors.last_name
             FROM authors
             WHERE authors.id IN (SELECT authors_products.author_id
                                  FROM authors_products
                                  WHERE authors_products.book_product_id = store_units.book_product_id)) AS author,
      store_units.price AS price,
       (SELECT AVG(value) FROM ratings
		WHERE store_unit_id = store_units.id AND store_units.book_product_id = book_products.id) AS value
FROM store_units
INNER JOIN book_products ON store_units.book_product_id = book_products.id
INNER JOIN product_types ON book_products.type = product_types.id;`

const GetBooks1 = `
SELECT name FROM book_products`
