package engine

import (
	"golang.org/x/net/context"

	"github.com/captaincodeman/clean/domain"
)

type (
	ListGreetingsRequest struct {
		Count int
	}

	ListGreetingsResponse struct {
		Greetings []*domain.Greeting
	}
)

func (g *greeter) List(c context.Context, r *ListGreetingsRequest) *ListGreetingsResponse {
	q := NewQuery("greeting").Order("date", Descending).Slice(0, r.Count)
	return &ListGreetingsResponse{
		Greetings: g.repository.List(c, q),
	}
}
