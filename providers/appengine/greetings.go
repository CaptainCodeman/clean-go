// +build appengine

package appengine

import (
	"time"

	"golang.org/x/net/context"
	"google.golang.org/appengine/datastore"

	"github.com/captaincodeman/clean/domain"
	"github.com/captaincodeman/clean/engine"
)

type (
	greetingRepository struct {
		context.Context
	}

	greeting struct {
		domain.Greeting
	}
)

var (
	greetingKind = "greeting"
)

func NewGreetingRepository(c context.Context) engine.GreetingRepository {
	return &greetingRepository{c}
}

func (r greetingRepository) Put(g *domain.Greeting) {
	d := &greeting{*g}
	k := datastore.NewIncompleteKey(r.Context, greetingKind, greetingEntityKey(r.Context))
	datastore.Put(r.Context, k, d)
}

func (r greetingRepository) List(query *engine.Query) []*domain.Greeting {
	g := []*greeting{}
	q := translateQuery(greetingKind, query)
	q = q.Ancestor(greetingEntityKey(r.Context))

	k , _ := q.GetAll(r.Context, &g)
	o := make([]*domain.Greeting, len(g))
	for i, _ := range g {
		o[i] = &g[i].Greeting
		o[i].ID = k[i].IntID()
	}
	return o
}

func greetingEntityKey(c context.Context) *datastore.Key {
	return datastore.NewKey(c, "guestbook", "", 1, nil)
}

func (x *greeting) Load(props []datastore.Property) error {
	for _, prop := range props {
		switch prop.Name {
			case "author":
				x.Author = prop.Value.(string)
			case "content":
				x.Content = prop.Value.(string)
			case "date":
				x.Date = prop.Value.(time.Time)
		}
	}
	return datastore.LoadStruct(x, props)
}

func (x *greeting) Save() ([]datastore.Property, error) {
	ps := []datastore.Property{
		datastore.Property{Name: "author", Value: x.Author, NoIndex: true, Multiple: false},
		datastore.Property{Name: "content", Value: x.Content, NoIndex: true, Multiple: false},
		datastore.Property{Name: "date", Value: x.Date, NoIndex: false, Multiple: false},
	}
	return ps, nil
}
