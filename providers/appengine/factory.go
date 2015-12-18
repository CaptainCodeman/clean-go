// +build appengine

package appengine

import (
	"golang.org/x/net/context"

	"github.com/captaincodeman/clean/engine"
)

type (
	storageFactory struct{}
)

func NewStorage() engine.StorageFactory {
	return &storageFactory{}
}

func (f *storageFactory) GetGreetingRepository(c context.Context) engine.GreetingRepository {
	return NewGreetingRepository(c)
}
