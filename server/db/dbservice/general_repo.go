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

	allStuff := []*dbqueries.GeneralQuerie{}

	for rows.Next() {
		var (
			lastName    string
			firstName   string
			bookName    string
			pageCount   int
			genreName   string
			productType string
		)

		rows.Scan(&lastName, &firstName, &bookName, &pageCount, &productType, &genreName)

		allStuff = append(allStuff, &dbqueries.GeneralQuerie{
			LastName:  lastName,
			FirstName: firstName,
			BookName:  bookName,
			PageCount: pageCount,
			Genre:     genreName,
			BookType:  productType,
		})
	}
	return allStuff
}
