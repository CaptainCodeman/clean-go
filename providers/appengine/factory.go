// +build appengine

package appengine

import (
	"github.com/captaincodeman/clean/engine"
)

type (
	storageFactory struct{}
)

func NewStorage() engine.StorageFactory {
	return &storageFactory{}
}

func (f *storageFactory) NewGreetingRepository() engine.GreetingRepository {
	return NewGreetingRepository()
}
