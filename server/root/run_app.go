package root

import (
	"books/server/conf"
	"books/server/db"
	"books/server/db/repo"
	"books/server/handler/bookHandler"
	"books/server/handler/commentHandlers"
	"books/server/handler/ratingHandler"
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

	createBookHandler := bookHandler.NewCreate()
	createBookHandler.Inject(book)

	deleteBookHandler := bookHandler.NewDeleteBook()
	deleteBookHandler.Inject(book)

	getAllBooksHandler := bookHandler.NewGetAllBooks()
	getAllBooksHandler.Inject(book)

	getBookIDHandler := bookHandler.NewGetBookByID()
	getBookIDHandler.Inject(book)

	commentRepo := repo.NewComments()
	commentRepo.Inject(dbConn)

	comment := service.NewAddComment()
	comment.Inject(commentRepo)

	addCommentHandler := commentHandlers.NewAddComment()
	addCommentHandler.Inject(comment)

	addCommentAnswerHandler := commentHandlers.NewAddCommentAnswer()
	addCommentAnswerHandler.Inject(comment)

	getCommentsHandler := commentHandlers.NewGetCommentsHandler()
	getCommentsHandler.Inject(comment)

	getCommentByIDHandler := commentHandlers.NewGetCommentByID()
	getCommentByIDHandler.Inject(comment)

	ratingRepo := repo.NewRating()
	ratingRepo.Inject(dbConn)

	rating := service.NewRating()
	rating.Inject(ratingRepo)

	addRatingHandler := ratingHandler.NewAddRating()
	addRatingHandler.Inject(rating)

	router := mux.NewRouter()

	router.
		Handle(createBookHandler.Path(), createBookHandler).
		Methods(createBookHandler.Method())

	router.
		Handle(getAllBooksHandler.Path(), getAllBooksHandler).
		Methods(getAllBooksHandler.Method())

	router.
		Handle(getBookIDHandler.Path(), getBookIDHandler).
		Methods(getBookIDHandler.Method())

	router.
		Handle(deleteBookHandler.Path(), deleteBookHandler).
		Methods(deleteBookHandler.Method())

	router.
		Handle(addCommentHandler.Path(), addCommentHandler).
		Methods(addCommentHandler.Method())

	router.
		Handle(addCommentAnswerHandler.Path(), addCommentAnswerHandler).
		Methods(addCommentAnswerHandler.Method())

	router.
		Handle(getCommentsHandler.Path(), getCommentsHandler).
		Methods(getCommentsHandler.Method())

	router.
		Handle(getCommentByIDHandler.Path(), getCommentByIDHandler).
		Methods(getCommentByIDHandler.Method())

	router.
		Handle(addRatingHandler.Path(), addRatingHandler).
		Methods(addRatingHandler.Method())

	server := http.Server{
		Handler: router,
		Addr:    "localhost:8081",
	}

	server.ListenAndServe()
}
