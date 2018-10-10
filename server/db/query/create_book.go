package query

var CreateBook = `
ROLLBACK;
INSERT INTO book_products (name, page_count, type)
VALUES ($1, $2, $3) RETURNING id;
INSERT INTO authors (first_name, last_name)
VALUES ($4, $5) RETURNING id;
INSERT INTO author_products(book_product_id, author_id)
VALUES ($1, $4);
INSERT INTO product_genres (book_product_id, genre_id)
VALUES ($1,$6);
COMMIT;`
