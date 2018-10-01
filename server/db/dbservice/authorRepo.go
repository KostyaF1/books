package dbservice

import (
	"books/server/models"
	"context"
	"fmt"
	"log"
)

type AuthorRepo struct {
}

func NewAuthorRepo() *AuthorRepo {
	return &AuthorRepo{}
}

func (author *AuthorRepo) GetAllAuthors(cxt context.Context) {
	//dbConn := author.DBer.DB()
	rows, err := dbConn.QueryContext(cxt, "SELECT last_name FROM authors")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	authors := []*models.Authors{}
	for rows.Next() {
		var (
			last_name string
		)

		rows.Scan(&last_name)

		authors = append(authors, &models.Authors{
			LastName: last_name,
		})
	}
	for _, auth := range authors {
		fmt.Println(auth.LastName)
	}
}
