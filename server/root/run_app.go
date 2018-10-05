package root

import (
	"books/server/conf"
	"books/server/db"
	"books/server/db/repo"
	"books/server/handler/bookHandlers"
	"books/server/service"
	"github.com/gorilla/mux"
	"net/http"
)

type App struct {
}

func (a *App) Run() {
	config := conf.Config{
		Driver: "postgres",
		DBUrl:  "postgres://postgres:1111@localhost/books",
	}

	dbConn := db.NewConn()
	dbConn.Inject(config)

	booksRepo := repo.NewBooks()
	booksRepo.Inject(dbConn)

	book := service.NewBook()
	book.Inject(booksRepo)

	createBookHandler := bookHandlers.NewCreateBook()
	createBookHandler.Inject(book)

	deleteBookHandler := bookHandlers.NewDeleteBook()
	deleteBookHandler.Inject(book)

	getAllBooksHandler := bookHandlers.NewGetAllBooks()
	getAllBooksHandler.Inject(book)

	router := mux.NewRouter()

	router.
		Handle(createBookHandler.Path(), createBookHandler).
		Methods(createBookHandler.Method())

	router.
		Handle(getAllBooksHandler.Path(), getAllBooksHandler).
		Methods(getAllBooksHandler.Method())

	router.
		Handle(deleteBookHandler.Path(), deleteBookHandler).
		Methods(deleteBookHandler.Method())

	server := http.Server{
		Handler: router,
		Addr:    "localhost:8080",
	}

	server.ListenAndServe()
}
