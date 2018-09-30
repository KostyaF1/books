package dbservice

import (
	"books/server/db"
	"database/sql"
	"log"
	//postgresql driver
	_ "github.com/lib/pq"
)

var dbConn *sql.DB

//OpenConnection function opens new connection with a data base
func OpenConnection(config *db.DataConfig) *sql.DB {
	dbConn, err := sql.Open(config.Driver, config.DataSource)
	if err != nil {
		log.Println("DB Open error" + error.Error(err))
	} else {
		log.Println("DB Open Ok")
	}
	return dbConn
}
