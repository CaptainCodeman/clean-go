package domain

import "errors"

type (
	UserID int64
	User   struct {
		ID     UserID
		Name   string
		Avatar string
		// more fields
	}

	Author struct {
		ID     UserID
		Name   string
		Avatar string
	}
)

func NewUser(name, avatar string) (*User, error) {
	if name == "" {
		return nil, errors.New("invalid user name")
	}
	if avatar == "" {
		return nil, errors.New("invalid user avatar")
	}
	return &User{
		Name:   name,
		Avatar: avatar,
	}, nil
}

func (u *User) Author() *Author {
	return &Author{
		ID:     u.ID,
		Name:   u.Name,
		Avatar: u.Avatar,
	}
}
