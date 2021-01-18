package main

import (
	"flag"
	"html/template"
	"log"
	"net/http"
	"time"
	"trannguyenhung011086/learn-go-web/pkg/logger"
	"trannguyenhung011086/learn-go-web/pkg/models/mysql"

	_ "github.com/go-sql-driver/mysql"
	"github.com/golangcollege/sessions"
)

type application struct {
	errorLog      *log.Logger
	infoLog       *log.Logger
	session       *sessions.Session
	snippets      *mysql.SnippetModel
	templateCache map[string]*template.Template
}

func main() {
	addr := flag.String("addr", ":4000", "HTTP address port")
	dsn := flag.String("dsn", "web:123456789@/snippetbox?parseTime=true", "MySQL database")
	secret := flag.String("secret", "s6Ndh+pPbnzHbS*+9Pk8qGWhTzbpa@ge", "Session Secret")
	flag.Parse()

	db, err := openDb(*dsn)
	if err != nil {
		logger.ErrorLog().Fatal(err)
	}
	defer db.Close()

	templateCache, err := newTemplateCache("./ui/html/")
	if err != nil {
		logger.ErrorLog().Fatal(err)
	}

	session := sessions.New([]byte(*secret))
	session.Lifetime = 12 * time.Hour

	app := &application{
		errorLog:      logger.ErrorLog(),
		infoLog:       logger.InfoLog(),
		session:       session,
		snippets:      &mysql.SnippetModel{DB: db},
		templateCache: templateCache,
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
