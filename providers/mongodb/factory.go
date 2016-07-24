package mongodb

import (
	"time"

	"gopkg.in/mgo.v2"

	"github.com/captaincodeman/clean/engine"
)

type (
	storageFactory struct {
		session *mgo.Session
	}
)

// NewStorage creates a new instance of this mongodb storage factory
func NewStorage(url string) engine.StorageFactory {
	session, _ := mgo.DialWithTimeout(url, 10*time.Second)
	ensureIndexes(session)
	return &storageFactory{session}
}

// NewGreetingRepository creates a new datastore greeting repository
func (f *storageFactory) NewGreetingRepository() engine.GreetingRepository {
	return newGreetingRepository(f.session)
}

func ensureIndexes(s *mgo.Session) {
	index := mgo.Index{
		Key:        []string{"date"},
		Background: true,
	}
	c := s.DB("").C(greetingCollection)
	c.EnsureIndex(index)
}
