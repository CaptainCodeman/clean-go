package engine

import (
	"golang.org/x/net/context"

	"github.com/captaincodeman/clean-go/domain"
)

type (
	AddGreetingRequest struct {
		Author  string
		Content string
	}

	AddGreetingResponse struct {
		ID int64
	}
)

func (g *greeter) Add(c context.Context, r *AddGreetingRequest) *AddGreetingResponse {
	// this is where all our app logic would go - the
	// rules that apply to adding a greeting whether it
	// is being done via the web UI, a console app, or
	// whatever the internet has just been added to ...
	greeting := domain.NewGreeting(r.Author, r.Content)
	g.repository.Put(c, greeting)
	return &AddGreetingResponse{
		ID: greeting.ID,
	}
}
