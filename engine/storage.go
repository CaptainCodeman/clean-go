package engine

import (
	"golang.org/x/net/context"

	"github.com/captaincodeman/clean/domain"
)

type (
	GreetingRepository interface {
		Put(c context.Context, greeting *domain.Greeting)
		List(c context.Context, query *Query) []*domain.Greeting
	}

	StorageFactory interface {
		NewGreetingRepository() GreetingRepository
	}
)
