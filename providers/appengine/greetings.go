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
	greetingRepository struct{}

	greeting struct {
		domain.Greeting
	}
)

var (
	greetingKind = "greeting"
)

func NewGreetingRepository() engine.GreetingRepository {
	return &greetingRepository{}
}

func (r greetingRepository) Put(c context.Context, g *domain.Greeting) {
	d := &greeting{*g}
	k := datastore.NewIncompleteKey(c, greetingKind, greetingEntityKey(c))
	datastore.Put(c, k, d)
}

func (r greetingRepository) List(c context.Context, query *engine.Query) []*domain.Greeting {
	g := []*greeting{}
	q := translateQuery(greetingKind, query)
	q = q.Ancestor(greetingEntityKey(c))

	k, _ := q.GetAll(c, &g)
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
