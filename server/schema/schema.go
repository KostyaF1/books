package schema

import (
	"books/server/db"
	"books/server/db/query"
	"fmt"
)

type Schema struct {
	dbConn db.Connector
}

func NewSchema() *Schema {
	return new(Schema)
}

func (s *Schema) Inject(dbConn db.Connector) {
	s.dbConn = dbConn
}

func (s *Schema) CreateTables() {
	dbConn := s.dbConn.Connect()

	_, err := dbConn.Exec(query.Schema)

	if err != nil {
		fmt.Println(err.Error())
	}
}
