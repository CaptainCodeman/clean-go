package engine

import (
	"golang.org/x/net/context"
)

type (
	EngineFactory interface {
		GetGreeter(c context.Context) *Greeter
	}

	engineFactory struct {
		StorageFactory
	}
)

func NewEngine(s StorageFactory) EngineFactory {
	return &engineFactory{s}
}
