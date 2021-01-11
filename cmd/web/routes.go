package main

import "net/http"

func (app *application) routes() http.Handler {
	mux := http.NewServeMux()
	mux.HandleFunc("/", app.home)
	mux.HandleFunc("/snippet", app.showSnippet)
	mux.HandleFunc("/snippet/create", app.createSnippet)

	serveStatic(mux, "./ui/static")

	return app.recoverPanic(app.logRequest(secureHeaders(mux)))
}
