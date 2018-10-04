package db

import (
	"books/server/conf"
	"database/sql"
	//PostgreSQL driver
	_ "github.com/lib/pq"
)

//Connector interface for connection to database
type Connector interface {
	Connect() *sql.DB
}

//conn ...
type Conn struct {
	config *conf.Config
	dbConn *sql.DB
}

var _ Connector = (*Conn)(nil)

//NewConn ...
func NewConn() *Conn {
	return new(Conn)
}

//Inject ...
func (c *Conn) Inject(config conf.Config) {
	c.config = &config
	c.init()
}

//init function opens new connection with a data base
func (c *Conn) init() {
	dbConn, err := sql.Open(c.config.Driver, c.config.DBUrl)
	if err != nil {
		panic(err)
	}
	c.dbConn = dbConn
}

//Connect ...
func (c *Conn) Connect() *sql.DB {
	return c.dbConn
}
