package main

import (
	"flag"
	"log"
	"net/http"
	"trannguyenhung011086/learn-go-web/pkg/logger"
	"trannguyenhung011086/learn-go-web/pkg/models/mysql"

	_ "github.com/go-sql-driver/mysql"
)

type application struct {
	errorLog *log.Logger
	infoLog  *log.Logger
	snippets *mysql.SnippetModel
}

func main() {
	addr := flag.String("addr", ":4000", "HTTP address port")
	dsn := flag.String("dsn", "web:123456789@/snippetbox?parseTime=true", "MySQL database")
	flag.Parse()

	db, err := openDb(*dsn)
	if err != nil {
		logger.ErrorLog().Fatal(err)
	}
	defer db.Close()

	app := &application{
		errorLog: logger.ErrorLog(),
		infoLog:  logger.InfoLog(),
		snippets: &mysql.SnippetModel{DB: db},
	}

	server := &http.Server{
		Addr:     *addr,
		ErrorLog: app.errorLog,
		Handler:  app.routes(),
	}

	app.infoLog.Printf("Starting server on port %s", *addr)

	err = server.ListenAndServe()
	if err != nil {
		app.errorLog.Fatal(err)
	}
}
