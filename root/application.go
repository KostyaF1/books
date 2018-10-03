package root

import (
	"books/app"
	"books/db"
	"books/handler"
	"books/repo"
	"books/service"
	"github.com/gorilla/mux"
	"net/http"
)

type Application struct {
}

func (application *Application) Run() {
	config := app.Config{
		Driver: "postgres",
		DbURL:  "",
	}

	query := db.NewQuery()
	query.Inject(config)

	booksRepo := repo.NewBooks()
	booksRepo.Inject(query)

	bookStore := service.NewBookStore()
	bookStore.Inject(booksRepo)

	createBookHandler := handler.NewCreateBook()
	createBookHandler.Inject(bookStore)

	router := mux.NewRouter()

	router.
		Handle(createBookHandler.Path(), createBookHandler).
		Methods(createBookHandler.Method())

	server := http.Server{
		Handler: router,
	}

	server.ListenAndServe()
}
