package logger

import (
	"log"
	"os"
)

// InfoLog : Used to log for INFO level
func InfoLog() *log.Logger {
	return log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
}

// ErrorLog : Used to log for ERROR level
func ErrorLog() *log.Logger {
	return log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)
}
