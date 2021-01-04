package main

import (
	"flag"
	"net/http"
	"trannguyenhung011086/learn-go-web/pkg/logger"
)

func main() {
	addr := flag.String("addr", ":4000", "HTTP address port")
	flag.Parse()

	app := &application{
		errorLog: logger.ErrorLog(),
		infoLog:  logger.InfoLog(),
	}

	server := &http.Server{
		Addr:     *addr,
		ErrorLog: app.errorLog,
		Handler:  app.routes(),
	}

	app.infoLog.Printf("Starting server on port %s", *addr)

	err := server.ListenAndServe()
	if err != nil {
		app.errorLog.Fatal(err)
	}
}
