package main

import (
	"fmt"
	"net/http"
	"path/filepath"
	"strconv"
	"trannguyenhung011086/learn-go-web/pkg/models"
)

func (app *application) home(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		app.notFoundError(w)
		return
	}

	s, err := app.snippets.Latest(10)
	if err != nil {
		app.serverError(w, err)
		return
	}

	app.render(w, r, "home.page.html", &templateData{Snippets: s})
}

func (app *application) showSnippet(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil || id < 1 {
		app.notFoundError(w)
		return
	}

	s, err := app.snippets.Get(id)
	if err == models.ErrNoRecord {
		app.notFoundError(w)
		return
	} else if err != nil {
		app.serverError(w, err)
		return
	}

	app.render(w, r, "show.page.html", &templateData{Snippet: s})
}

func (app *application) createSnippet(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Allow", "POST")
	w.Header().Set("Content-Type", "application/json")
	if r.Method != "POST" {
		app.clientError(w, http.StatusMethodNotAllowed)
		return
	}

	title := "O snail"
	content := "O snail\nClimb Mount Fuji,\nBut slowly, slowly!\n\n– Kobayashi"
	expires := "7"

	id, err := app.snippets.Insert(title, content, expires)
	if err != nil {
		app.serverError(w, err)
		return
	}

	http.Redirect(w, r, fmt.Sprintf("/snippet?id=%d", id), http.StatusSeeOther)
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
