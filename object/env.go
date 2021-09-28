package object

// Environment represents a local environment that keeps track
// of identifiers and their values within a session.
type Environment struct {
	store map[string]Object
}

// NewEnvironment return new instance of Environment.
func NewEnvironment() *Environment {
	s := make(map[string]Object)

	return &Environment{store: s}
}

func (e *Environment) Get(name string) (Object, bool) {
	obj, ok := e.store[name]

	return obj, ok
}
func (e *Environment) Set(name string, val Object) Object {
	e.store[name] = val

	return val
}
