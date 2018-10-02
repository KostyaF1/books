package db

import (
	"database/sql"
	"log"
	//postgresql driver
	_ "github.com/lib/pq"
)

//Conn ...
type Conn struct {
	DBConn *sql.DB
}

//DBer interface for connection to database
type DBer interface {
	DB() *sql.DB
}

//DB ...
func (con *Conn) DB() *sql.DB {
	return con.DBConn
}

//NewConn function opens new connection with a data base
func NewConn(config *DataConfig) (*Conn, error) {
	dbConn, err := sql.Open(config.Driver, config.DataSource)
	if err != nil {
		return nil, err
	}

	log.Println("DB Open Ok")
	return &Conn{
		DBConn: dbConn,
	}, nil
}
