package engine

type (
	EngineFactory interface {
		NewGreeter() Greeter
	}

	engineFactory struct {
		StorageFactory
	}
)

func NewEngine(s StorageFactory) EngineFactory {
	return &engineFactory{s}
}
