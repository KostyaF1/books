package dbservice

import (
	"context"
	"fmt"
	"log"
)

//General ...
type General struct {
	ID        int
	LastName  string
	FirstName string
	BookName  string
	PageCount int
	Genre     string
	BookType  string
}

//GeneralRepo ...
type GeneralRepo struct {
	DBer
}

//NewGeneralRepo ...
func NewGeneralRepo(conn *Conn) *GeneralRepo {
	return &GeneralRepo{
		DBer: conn,
	}
}

//GetAll ...
func (generalRepo *GeneralRepo) GetAll(cxt context.Context) {
	dbConn := generalRepo.DBer.DB()
	rows, err := dbConn.QueryContext(cxt, generalQuerie)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	allStuff := []*General{}

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

		allStuff = append(allStuff, &General{
			LastName:  lastName,
			FirstName: firstName,
			BookName:  bookName,
			PageCount: pageCount,
			Genre:     genreName,
			BookType:  productType,
		})
	}
	for _, stuff := range allStuff {
		fmt.Println(stuff.BookName, stuff.LastName, stuff.FirstName,
			stuff.Genre, stuff.BookType, stuff.BookType)
	}
}
