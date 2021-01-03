package main

import (
	"flag"
	"net/http"
	"trannguyenhung011086/learn-go-web/pkg/logger"
)

var (
	infoLog  = logger.InfoLog()
	errorLog = logger.ErrorLog()
)

func main() {
	addr := flag.String("addr", ":4000", "HTTP address port")
	flag.Parse()

	mux := http.NewServeMux()
	mux.HandleFunc("/", home)
	mux.HandleFunc("/snippet", showSnippet)
	mux.HandleFunc("/snippet/create", createSnippet)

	serveStatic(mux, "./ui/static")

	infoLog.Printf("Starting server on port %s", *addr)

	server := &http.Server{
		Addr:     *addr,
		ErrorLog: errorLog,
		Handler:  mux,
	}

	err := server.ListenAndServe()
	if err != nil {
		errorLog.Fatal(err)
	}
}
