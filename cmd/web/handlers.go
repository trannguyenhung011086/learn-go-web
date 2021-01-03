package main

import (
	"fmt"
	"net/http"
	"path/filepath"
	"strconv"
	"text/template"
)

func home(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}

	files := []string{
		"./ui/html/home.page.html1",
		"./ui/html/base.layout.html",
		"./ui/html/footer.partial.html",
	}

	ts, err := template.ParseFiles(files...)
	if err != nil {
		errorLog.Println(err.Error())
		http.Error(w, "Internal error", http.StatusInternalServerError)
		return
	}
	err = ts.Execute(w, nil)
	if err != nil {
		errorLog.Println(err.Error())
		http.Error(w, "Internal error", http.StatusInternalServerError)
		return
	}
}

func showSnippet(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil || id < 0 {
		http.NotFound(w, r)
		return
	}
	fmt.Fprintf(w, "Displaying snippet with id %d", id)
}

func createSnippet(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Allow", "POST")
	w.Header().Set("Content-Type", "application/json")
	if r.Method != "POST" {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
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
