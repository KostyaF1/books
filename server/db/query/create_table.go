package query

const CreateDB = `CREATE DATABASE books_test OWNER postgres;`

const Schema = `
CREATE TABLE IF NOT EXISTS product_types (
	id integer PRIMARY KEY NOT NULL,
	name varchar(40) NOT NULL
);

CREATE TABLE IF NOT EXISTS genres (
	id integer PRIMARY KEY NOT NULL,
	name varchar(40) NOT NULL
);

CREATE TABLE IF NOT EXISTS book_products (
	id serial PRIMARY KEY,
	name varchar(40) NOT NULL,
	page_count integer NOT NULL,
	type integer REFERENCES product_types(id)
);

CREATE TABLE IF NOT EXISTS authors (
	id serial PRIMARY KEY,
	first_name varchar(40) NOT NULL,
	last_name varchar(40) NOT NULL
);
CREATE TABLE IF NOT EXISTS authors_products (
	id serial PRIMARY KEY,
	book_product_id integer REFERENCES book_products(id),
	author_id integer REFERENCES authors(id)
);
CREATE TABLE IF NOT EXISTS product_genres (
	id serial PRIMARY KEY,
	book_product_id integer REFERENCES book_products(id),
	genre_id integer REFERENCES genres(id)
);
CREATE TABLE IF NOT EXISTS store_units (
	id serial PRIMARY KEY,
	book_product_id integer REFERENCES book_products(id),
	price integer NOT NULL
);

CREATE TABLE IF NOT EXISTS ratings (
	id serial PRIMARY KEY,
	store_unit_id integer REFERENCES store_units(id),
	value integer
);

CREATE TABLE IF NOT EXISTS comments (
	id serial PRIMARY KEY,
	body text,
	book_unit_id integer REFERENCES store_units(id)
);
`
