package domain

import (
	"errors"
	"time"
)

type (
	// TopicID identifies a topic
	TopicID int64

	//Topic represents a unique discussion topic within the system
	Topic struct {
		ID      TopicID
		Author  Author
		Title   string
		Html    string
		Started time.Time
	}
)

// NewTopic creates a new topic, validating the passed parameters
func NewTopic(author Author, title, html string) (*Topic, error) {
	if title == "" {
		return nil, errors.New("invalid topic title")
	}
	if html == "" {
		return nil, errors.New("invalid topic html")
	}
	return &Topic{
		Author:  author,
		Title:   title,
		Html:    html,
		Started: Now(),
	}, nil
}
