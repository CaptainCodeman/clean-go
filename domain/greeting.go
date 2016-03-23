package domain

import (
	"time"
)

type (
	Greeting struct {
		ID      int64     `json:"id"`
		Author  string    `json:"author"`
		Content string    `json:"content"`
		Date    time.Time `json:"timestamp"`
	}
)

func NewGreeting(author, content string) *Greeting {
	return &Greeting{
		Author:  author,
		Content: content,
		Date:    now(),
	}
}
