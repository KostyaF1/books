package db

import (
	"books/app"
	"database/sql"
)

type Query interface {
	DB() *sql.DB
}

func NewQuery() *query {
	return new(query)
}

var _ Query = (*query)(nil)

func (q *query) Inject(config app.Config) {
	q.config = config
	q.init()
}

func (q *query) init() {
	dbConn, err := sql.Open(q.config.Driver, q.config.DbURL)

	if err != nil {
		panic(err)
	}

	q.db = dbConn
}

type query struct {
	db     *sql.DB
	config app.Config
}

func (q *query) DB() *sql.DB {
	return q.db
}
