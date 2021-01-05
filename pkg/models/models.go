package models

import (
	"errors"
	"time"
)

// ErrNoRecord : error when no matching record found
var ErrNoRecord = errors.New("models: no matching record found")

// Snippet : object model for snippet
type Snippet struct {
	ID      int
	Title   string
	Content string
	Created time.Time
	Expires time.Time
}
