package main

import (
	"flag"
	"log"
	"net/http"
	"trannguyenhung011086/learn-go-web/pkg/logger"
)

type application struct {
	errorLog *log.Logger
	infoLog  *log.Logger
}

func main() {
	addr := flag.String("addr", ":4000", "HTTP address port")
	flag.Parse()

	app := &application{
		errorLog: logger.ErrorLog(),
		infoLog:  logger.InfoLog(),
	}

	mux := http.NewServeMux()
	mux.HandleFunc("/", app.home)
	mux.HandleFunc("/snippet", app.showSnippet)
	mux.HandleFunc("/snippet/create", app.createSnippet)

	app.serveStatic(mux, "./ui/static")

	app.infoLog.Printf("Starting server on port %s", *addr)

	server := &http.Server{
		Addr:     *addr,
		ErrorLog: app.errorLog,
		Handler:  mux,
	}

	err := server.ListenAndServe()
	if err != nil {
		app.errorLog.Fatal(err)
	}
}
