package domain

import (
	"time"
)

type (
	// Greeting is the struct that would contain any
	// domain logic if we had any. Because it's simple
	// we're going to send it over the wire directly
	// so we add the JSON serialization tags but we
	// could use DTO specific structs for that
	Greeting struct {
		ID      int64     `json:"id"`
		Author  string    `json:"author"`
		Content string    `json:"content"`
		Date    time.Time `json:"timestamp"`
	}
)

// NewGreeting creates a new Greeting ... who'da thunk it!
func NewGreeting(author, content string) *Greeting {
	return &Greeting{
		Author:  author,
		Content: content,
		Date:    now(),
	}
}
