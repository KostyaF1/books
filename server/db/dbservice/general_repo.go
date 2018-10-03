package dbservice

import (
	"books/server/db"
	"books/server/db/dbqueries"
	"context"
	"log"
)

//GeneralRepo ...
type GeneralRepo struct {
	db.DBer
}

//NewGeneralRepo ...
func NewGeneralRepo(conn *db.Conn) *GeneralRepo {
	return &GeneralRepo{
		DBer: conn,
	}
}

//GetAllUnits ...
func (generalRepo GeneralRepo) GetAllUnits(cxt context.Context) []*dbqueries.GeneralQuerie {
	dbConn := generalRepo.DBer.DB()
	rows, err := dbConn.QueryContext(cxt, dbqueries.GetAllUnitsQuerie)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	var allStuff []*dbqueries.GeneralQuerie

	for rows.Next() {
		var (
			bookName    string
			genreName   string
			productType string
			pageCount   int
			author      string
		)

		rows.Scan(&bookName, &genreName, &productType, &pageCount, &author)

		allStuff = append(allStuff, &dbqueries.GeneralQuerie{
			Author:    author,
			BookName:  bookName,
			PageCount: pageCount,
			Genre:     genreName,
			BookType:  productType,
		})
	}
	return allStuff
}
