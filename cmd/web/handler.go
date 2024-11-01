package main

import(
	"errors"
	"net/http"
	"log"
	"html/template"
	"strconv"

	"github.com/udbhav-rohatgi/library/internal/models"
	"github.com/gorilla/mux"
)

func (app *application) home(w http.ResponseWriter, r *http.Request){
	if r.URL.Path != "/"{
		app.notFound(w)
		return
	}

	// books, err := app.books.Latest()
	// if err != nil {
	// 	app.serverError(w, err)
	// 	return
	// }

	// data := app.newTemplateData(r)
	// data.Books = books

	// app.render(w, http.StatusOK, "home.tmpl", data)

		// Initialize a slice containing the paths to the two files. It's important
		// to note that the file containing our base template must be the *first*
		// file in the slice.
		files := []string{
		"./ui/html/base.tmpl",
		"./ui/html/pages/home.tmpl",
		}
		// Use the template.ParseFiles() function to read the files and store the
		// templates in a template set. Notice that we can pass the slice of file
		// paths as a variadic parameter?
		ts, err := template.ParseFiles(files...)
		if err != nil {
		log.Println(err.Error())
		http.Error(w, "Internal Server Error", 500)
		return
		}
		// Use the ExecuteTemplate() method to write the content of the "base"
		// template as the response body.
		err = ts.ExecuteTemplate(w, "base", nil)
		if err != nil {
		log.Println(err.Error())
		http.Error(w, "Internal Server Error", 500)
		}
}

func(app *application) book(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	params := mux.Vars(r)
	idStr := params["id"]

	myId, err := strconv.Atoi(idStr)
	if err != nil{
		http.Error(w, "Invalid Book ID", http.StatusBadRequest)
		return
	}

	book, err := app.books.Get(myId)
	if err != nil {
		if errors.Is(err, models.ErrNoRecord){
			app.notFound(w)
		} else{
			app.serverError(w, err)
		}
		return
	}

	data := app.newTemplateData(r)
	data.Book = book

	// app.render(w, http.StatusOK, "view.tmpl", data)
}

func(app *application) homeCreate(w http.ResponseWriter, r *http.Request){
	if r.Method != http.MethodPost{
		w.Header().Set("Allow", http.MethodPost)
		http.Error(w, "Method not Allowed", http.StatusMethodNotAllowed)
		return
	}

	err := r.ParseForm()
	if err != nil {
		app.serverError(w, err)
		return
	}

	taskName := r.FormValue("task")

	if taskName == "" {
		http.Error(w, "Task name is required", http.StatusBadRequest)
		return
	}

	_, err = app.books.Insert(taskName)
	if err != nil{
		app.serverError(w, err)
		return
	}

	http.Redirect(w, r, "/", http.StatusSeeOther)
}

func (app *application) homeDelete(w http.ResponseWriter, r *http.Request){

	if r.Method == http.MethodPost && r.FormValue("_method") == "DELETE"{
		vars := mux.Vars(r)
		idStr := vars["id"]
	

		id,err := strconv.Atoi(idStr)
		if err != nil{
			http.Error(w, "Invalid task ID", http.StatusBadRequest)
			return 
		}

		err = app.books.Delete(id)
		if err != nil{
			app.serverError(w, err)
			return
		}

		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}

	http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
}