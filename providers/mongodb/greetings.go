package mongodb

import (
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"

	"github.com/captaincodeman/clean/domain"
	"github.com/captaincodeman/clean/engine"
)

type (
	greetingRepository struct {
		session *mgo.Session
	}

	greeting struct {
		domain.Greeting
	}
)

var (
	greetingCollection = "greeting"
)

func NewGreetingRepository(session *mgo.Session) engine.GreetingRepository {
	return &greetingRepository{session}
}

func (r greetingRepository) Put(g *domain.Greeting) {
	s := r.session.Clone()
	defer s.Close()

	c := s.DB("").C(greetingCollection)
	if g.ID == 0 {
		g.ID = getNextSequence(s, greetingCollection)
	}
	c.Upsert(bson.M{"_id": g.ID}, g)
}

func (r greetingRepository) List(query *engine.Query) []*domain.Greeting {
	s := r.session.Clone()
	defer s.Close()

	c := s.DB("").C(greetingCollection)
	g := []*domain.Greeting{}
	q := translateQuery(c, query)
	q.All(&g)

	return g
}
