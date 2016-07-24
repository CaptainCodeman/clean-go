package mongodb

import (
	"golang.org/x/net/context"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"

	"github.com/captaincodeman/clean-go/domain"
	"github.com/captaincodeman/clean-go/engine"
)

type (
	greetingRepository struct {
		session *mgo.Session
	}
)

var (
	greetingCollection = "greeting"
)

func newGreetingRepository(session *mgo.Session) engine.GreetingRepository {
	return &greetingRepository{session}
}

func (r greetingRepository) Put(c context.Context, g *domain.Greeting) {
	s := r.session.Clone()
	defer s.Close()

	col := s.DB("").C(greetingCollection)
	if g.ID == 0 {
		g.ID = getNextSequence(s, greetingCollection)
	}
	col.Upsert(bson.M{"_id": g.ID}, g)
}

func (r greetingRepository) List(c context.Context, query *engine.Query) []*domain.Greeting {
	s := r.session.Clone()
	defer s.Close()

	col := s.DB("").C(greetingCollection)
	g := []*domain.Greeting{}
	q := translateQuery(col, query)
	q.All(&g)

	return g
}
