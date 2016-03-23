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

func NewStorage(url string) engine.StorageFactory {
	session, _ := mgo.DialWithTimeout(url, 10*time.Second)
	ensureIndexes(session)
	return &storageFactory{session}
}

func (f *storageFactory) NewGreetingRepository() engine.GreetingRepository {
	return NewGreetingRepository(f.session)
}

func ensureIndexes(s *mgo.Session) {
	index := mgo.Index{
		Key:        []string{"date"},
		Background: true,
	}
	c := s.DB("").C(greetingCollection)
	c.EnsureIndex(index)
}
