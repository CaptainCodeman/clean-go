package engine

import (
	"golang.org/x/net/context"

	"github.com/captaincodeman/clean/domain"
)

type (
	Greeter interface {
		Add(c context.Context, r *AddGreetingRequest) *AddGreetingResponse
		List(c context.Context, r *ListGreetingsRequest) *ListGreetingsResponse
	}

	greeter struct {
		repository GreetingRepository
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

func NewGreeter(repository GreetingRepository) Greeter {
	return &greeter{
		repository: repository,
	}
}
func (f *engineFactory) NewGreeter() Greeter {
	return NewGreeter(f.NewGreetingRepository())
}

func (g *greeter) Add(c context.Context, r *AddGreetingRequest) *AddGreetingResponse {
	greeting := domain.NewGreeting(r.Author, r.Content)
	g.repository.Put(c, greeting)
	return &AddGreetingResponse{
		ID: greeting.ID,
	}
}

func (g *greeter) List(c context.Context, r *ListGreetingsRequest) *ListGreetingsResponse {
	q := NewQuery("greeting").Order("date", Descending).Slice(0, r.Count) // .Filter("author", core.NotEqual, "")
	return &ListGreetingsResponse{
		Greetings: g.repository.List(c, q),
	}
}
