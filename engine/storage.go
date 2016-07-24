package engine

import (
	"golang.org/x/net/context"

	"github.com/captaincodeman/clean-go/domain"
)

type (
	// GreetingRepository defines the methods that any
	// data storage provider needs to implement to get
	// and store greetings
	GreetingRepository interface {
		// Put adds a new Greeting to the datastore
		Put(c context.Context, greeting *domain.Greeting)

		// List returns existing greetings matching the
		// query provided
		List(c context.Context, query *Query) []*domain.Greeting
	}

	// StorageFactory is the interface that a storage
	// provider needs to implement so that the engine can
	// request repository instances as it needs them
	StorageFactory interface {
		// NewGreetingRepository returns a storage specific
		// GreetingRepository implementation
		NewGreetingRepository() GreetingRepository
	}
)
