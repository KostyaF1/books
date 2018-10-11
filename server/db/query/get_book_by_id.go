package query

const GetBookByID = `SELECT id,
       (SELECT name FROM book_products),
       array(SELECT genres.name
             FROM genres
             WHERE genres.id IN
                   (SELECT product_genres.genre_id
                    FROM product_genres
                    WHERE product_genres.book_product_id = store_units.book_product_id)) AS genre,
       (SELECT product_types.name FROM product_types WHERE id IN
                                                           (SELECT type FROM book_products)) AS type,
       (SELECT page_count FROM book_products),
       array(SELECT authors.first_name || ' ' || authors.last_name
             FROM authors
             WHERE authors.id IN (SELECT author_products.author_id
                                  FROM author_products
                                  WHERE author_products.book_product_id = store_units.book_product_id)) AS author,
       price,
       array(SELECT body || '-author:' || author FROM coments) AS comments
FROM store_units
WHERE id = $1;`
