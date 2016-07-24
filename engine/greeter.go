package engine

import (
	"golang.org/x/net/context"
)

type (
	// Greeter is the interface for our interactor
	Greeter interface {
		// Add is the add-a-greeting use-case
		Add(c context.Context, r *AddGreetingRequest) *AddGreetingResponse

		// List is the list-the-greetings use-case
		List(c context.Context, r *ListGreetingsRequest) *ListGreetingsResponse
	}

	// greeter implementation
	greeter struct {
		repository GreetingRepository
	}
)

// NewGreeter creates a new Greeter interactor wired up
// to use the greeter repository from the storage provider
// that the engine has been setup to use.
func (f *engineFactory) NewGreeter() Greeter {
	return &greeter{
		repository: f.NewGreetingRepository(),
	}
}
