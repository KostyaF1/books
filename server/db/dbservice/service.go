package dbservice

import (
	"fmt"
	"log"
)

type AuthorDB struct {
	LastName string
}

func (author *AuthorDB) ASelect() {
	rows, err := dbConn.Query("SELECT * FROM authors")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	//bks := make([]*Book, 0)
	bks := make([]*AuthorDB, 0)
	for rows.Next() {

		err := rows.Scan(&author.LastName)
		if err != nil {
			log.Fatal(err)
		}
		bks = append(bks, author)
	}
	if err = rows.Err(); err != nil {
		log.Fatal(err)
	}

	for _, author := range bks {
		fmt.Printf("%s\n", &author.LastName)
	}
}
