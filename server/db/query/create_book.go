package query

const CreateBook = `INSERT INTO book_products (name, page_count, type)
					VALUES ($1, $2, (SELECT id FROM product_types WHERE name = $3))
					ON CONFLICT(name) DO UPDATE SET name = EXCLUDED.name RETURNING id;`

const CreateAuthor = `INSERT INTO authors (first_name, last_name)
						VALUES ($1, $2) ON CONFLICT (first_name, last_name) DO UPDATE SET
						first_name = EXCLUDED.first_name RETURNING id;`

const AuthorBook = `INSERT INTO authors_products (book_product_id, author_id)
					VALUES ($1, $2) RETURNING id;`

const GenreBook = `INSERT INTO product_genres (book_product_id, genre_id)
					VALUES ($1, (SELECT id FROM genres WHERE name = $2));`

const DefRating = `INSERT INTO ratings (store_unit_id, value)
					VALUES ($1, 0.0);`

const StoreUnit = `INSERT INTO store_units (book_product_id, price)
					VALUES ($1, $2) ON CONFLICT (book_product_id) DO UPDATE SET
					book_product_id = EXCLUDED.book_product_id RETURNING id;`
