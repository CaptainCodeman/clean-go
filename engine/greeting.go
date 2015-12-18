package engine

import (
	"golang.org/x/net/context"

	"github.com/captaincodeman/clean/domain"
)

type (
	Greeter struct {
		Repository GreetingRepository
	}

	AddGreetingRequest struct {
		Author  string
		Content string
	}

	AddGreetingResponse struct {
		ID int64
	}

	ListGreetingsRequest struct {
		Count int
	}

	ListGreetingsResponse struct {
		Greetings []*domain.Greeting
	}
)

func NewGreeter(repository GreetingRepository) *Greeter {
	return &Greeter{
		Repository: repository,
	}
}
func (f *engineFactory) GetGreeter(c context.Context) *Greeter {
	return NewGreeter(f.GetGreetingRepository(c))
}

func (g *Greeter) Add(r *AddGreetingRequest) *AddGreetingResponse {
	greeting := domain.NewGreeting(r.Author, r.Content)
	g.Repository.Put(greeting)
	return &AddGreetingResponse{
		ID: greeting.ID,
	}
}

func (g *Greeter) List(r *ListGreetingsRequest) *ListGreetingsResponse {
	q := NewQuery("greeting").Order("date", Descending).Slice(0, r.Count) // .Filter("author", core.NotEqual, "")
	return &ListGreetingsResponse{
		Greetings: g.Repository.List(q),
	}
}
