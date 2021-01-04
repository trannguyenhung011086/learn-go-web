package main

import (
	"fmt"
	"log"
	"net/http"
	"path/filepath"
	"strconv"
	"text/template"
)

type application struct {
	errorLog *log.Logger
	infoLog  *log.Logger
}

func (app *application) home(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		app.notFoundError(w)
		return
	}

	files := []string{
		"./ui/html/home.page.html",
		"./ui/html/base.layout.html",
		"./ui/html/footer.partial.html",
	}

	ts, err := template.ParseFiles(files...)
	if err != nil {
		app.serverError(w, err)
		return
	}
	err = ts.Execute(w, nil)
	if err != nil {
		app.serverError(w, err)
		return
	}
}

func (app *application) showSnippet(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil || id < 0 {
		app.notFoundError(w)
		return
	}
	fmt.Fprintf(w, "Displaying snippet with id %d", id)
}

func (app *application) createSnippet(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Allow", "POST")
	w.Header().Set("Content-Type", "application/json")
	if r.Method != "POST" {
		app.clientError(w, http.StatusMethodNotAllowed)
		return
	}
	w.Write([]byte(`{"message": "Creating snippet"}`))
}

func serveStatic(mux *http.ServeMux, path string) *http.ServeMux {
	fileServer := http.FileServer(http.Dir(path))
	mux.Handle("/static/", http.StripPrefix("/static", fileServer))
	return mux
}

func downloadFile(w http.ResponseWriter, r *http.Request, path string) {
	filepath.Clean(path)
	http.ServeFile(w, r, path)
}
