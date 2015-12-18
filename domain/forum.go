package domain

import "errors"

type (
	ForumID int64
	Forum   struct {
		ID                ForumID
		Title             string
		Domain            string
		AllowMessaging    bool
		AllowRegistration bool
		AllowPosting      bool

		// favicon, icon, logo
		// about
		// analytics, adsense, recaptcha, addthis
	}
)

var (
	ErrRegistrationClosed = errors.New("registration closed")
	ErrPostingDisabled    = errors.New("posting is disabled")
)

func NewForum(title, domain string) (*Forum, error) {
	return &Forum{
		Title:             title,
		Domain:            domain,
		AllowMessaging:    true,
		AllowPosting:      true,
		AllowRegistration: true,
	}, nil
}

func (f *Forum) Open() *Forum {
	f.AllowMessaging = true
	f.AllowPosting = true
	f.AllowRegistration = true
	return f
}

func (f *Forum) Close() *Forum {
	f.AllowMessaging = false
	f.AllowPosting = false
	f.AllowRegistration = false
	return f
}

func (f *Forum) CreateUser(name, avatar string) (*User, error) {
	if !f.AllowRegistration {
		return nil, ErrRegistrationClosed
	}
	return NewUser(name, avatar)
}

func (f *Forum) CreateTopic(author Author, title, html string) (*Topic, error) {
	if !f.AllowPosting {
		return nil, ErrPostingDisabled
	}
	return NewTopic(author, title, html)
}
