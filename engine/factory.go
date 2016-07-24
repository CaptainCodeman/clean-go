package engine

type (
	// EngineFactory interface allows us to provide
	// other parts of the system with a way to make
	// instances of our use-case / interactors when
	// they need to
	EngineFactory interface {
		// NewGreeter creates a new Greeter interactor
		NewGreeter() Greeter
	}

	// engine factory stores the state of our engine
	// which only involves a storage factory instance
	engineFactory struct {
		StorageFactory
	}
)

// NewEngine creates a new engine factory that will
// make use of the passed in StorageFactory for any
// data persistence needs.
func NewEngine(s StorageFactory) EngineFactory {
	return &engineFactory{s}
}
