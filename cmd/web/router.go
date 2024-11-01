package main

import (
	"net/http"

	"github.com/gorilla/mux"
)

func(app *application) routes() *mux.Router{
	router := mux.NewRouter()

	// router.Use(secureHeaders)
	// router.Use(app.logRequest)
	// router.Use(app.recoverPanic)

	fileServer := http.FileServer(http.Dir("./ui/static"))
	router.PathPrefix("/static/").Handler(http.StripPrefix("/static", fileServer))

	router.HandleFunc("/", app.home).Methods("GET")
	// router.HandleFunc("/{id}", app.book).Methods("GET")
	// router.HandleFunc("/add", app.addBook).Methods("POST")
	// router.HandleFunc("/{id}", app.deleteBook).Methods("GET")

	return router
}