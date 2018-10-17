package root

import (
	"books/server/conf"
	"books/server/db"
	"books/server/db/repo"
	"books/server/handler"
	"books/server/handler/book_handler"
	"books/server/handler/comment_handler"
	"books/server/service"
	"github.com/gorilla/mux"
	"net/http"
	"os"
)

type App struct {
}

func (a *App) Run() {

	config := conf.Config{
		Driver: os.Getenv("DRIVER"),
		DBUrl:  os.Getenv("DBURL"),
	}

	dbConn := db.NewConn()
	dbConn.Inject(config)

	booksRepo := repo.NewBooks()
	booksRepo.Inject(dbConn)

	book := service.NewBook()
	book.Inject(booksRepo)

	createBookHandler := book_handler.NewCreate()
	createBookHandler.Inject(book)

	deleteBookHandler := book_handler.NewDelete()
	deleteBookHandler.Inject(book)

	getAllBooksHandler := book_handler.NewGetAll()
	getAllBooksHandler.Inject(book)

	getBookIDHandler := book_handler.NewGetByID()
	getBookIDHandler.Inject(book)

	commentRepo := repo.NewComments()
	commentRepo.Inject(dbConn)

	comment := service.NewComment()
	comment.Inject(commentRepo)

	addCommentHandler := comment_handler.NewAdd()
	addCommentHandler.Inject(comment)

	addCommentAnswerHandler := comment_handler.NewAddCommentAnswer()
	addCommentAnswerHandler.Inject(comment)

	getCommentsHandler := comment_handler.NewGetAll()
	getCommentsHandler.Inject(comment)

	getCommentByIDHandler := comment_handler.NewGetByID()
	getCommentByIDHandler.Inject(comment)

	deleteCommentHandler := comment_handler.NewDelete()
	deleteCommentHandler.Inject(comment)

	ratingRepo := repo.NewRating()
	ratingRepo.Inject(dbConn)

	rating := service.NewRating()
	rating.Inject(ratingRepo)

	addRatingHandler := handler.NewAdd()
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
		Handle(deleteCommentHandler.Path(), deleteCommentHandler).
		Methods(deleteCommentHandler.Method())

	router.
		Handle(addRatingHandler.Path(), addRatingHandler).
		Methods(addRatingHandler.Method())

	server := http.Server{
		Handler: router,
		Addr:    "localhost:8081",
	}

	server.ListenAndServe()
}
