package main

import (
	"crypto/tls"
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
	users         *mysql.UserModel
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
	session.Secure = true

	app := &application{
		errorLog:      logger.ErrorLog(),
		infoLog:       logger.InfoLog(),
		session:       session,
		snippets:      &mysql.SnippetModel{DB: db},
		users:         &mysql.UserModel{DB: db},
		templateCache: templateCache,
	}

	tlsConfig := &tls.Config{
		PreferServerCipherSuites: true,
		CurvePreferences:         []tls.CurveID{tls.X25519, tls.CurveP256},
	}

	server := &http.Server{
		Addr:         *addr,
		ErrorLog:     app.errorLog,
		Handler:      app.routes(),
		TLSConfig:    tlsConfig,
		IdleTimeout:  time.Minute,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	app.infoLog.Printf("Starting server on port %s", *addr)

	err = server.ListenAndServeTLS("./tls/cert.pem", "./tls/key.pem")
	if err != nil {
		app.errorLog.Fatal(err)
	}
}
