package dbservice

import (
	"books/server/db"
	"database/sql"
	//postgresql driver
	_ "github.com/lib/pq"
)

var dbConn *sql.DB

//OpenConnection function opens new connection with a data base
func NewConn(config *db.DataConfig) *sql.DB {
	dbConn, _ = sql.Open(config.Driver, config.DataSource)
	//if err != nil {
	//	return nil, err
	//} else {
	//	log.Println("DB Open Ok")
	//}
	return dbConn
}
