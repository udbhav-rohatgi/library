package main

import (
	"database/sql"
	"html/template"
	"log"
	"flag"
	"net/http"
	"os"

	"github.com/udbhav-rohatgi/library/internal/models"
	_ "github.com/lib/pq"
)

type application struct{
	errorLog		*log.Logger
	infoLog 		*log.Logger
	books			*models.BookModel
	templateCache	map[string]*template.Template
}

func main(){
	addr := flag.String("addr", ":8080", "HTTP network address")
	dsn := flag.String("dsn", "postgres://booksdb:pa55word@localhost/booksdb", "PostgreSQL DSN")

	flag.Parse()

	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	errorLog := log.New(os.Stdout, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)

	// databse connection ---------------------------------------
	db, err := openDB(*dsn)
	if err != nil {
		errorLog.Fatal(err)
	}

	defer db.Close()
	// -----------------------------------------------------------

	app := &application{
		errorLog: 	errorLog,
		infoLog:	infoLog,
		books: 		&models.BookModel{DB: db},
	}

	srv := &http.Server{
		Addr: *addr,
		ErrorLog: errorLog,
		Handler: app.routes(),
	}

	infoLog.Printf("Starting Server on %s", *addr)
	err = srv.ListenAndServe()
	errorLog.Fatal(err)

}

func openDB(dsn string) (*sql.DB, error){
	db,err := sql.Open("postgres", dsn)
	if err != nil {
		return nil, err
	}

	if err = db.Ping(); err != nil{
		return nil, err
	}
	return db, nil
}