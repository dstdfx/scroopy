package object

// Environment represents a local environment that keeps track
// of identifiers and their values within a session.
type Environment struct {
	store map[string]Object
	outer *Environment
}

// NewEnvironment return new instance of Environment.
func NewEnvironment() *Environment {
	s := make(map[string]Object)

	return &Environment{store: s}
}

// NewEnclosedEnvironment returns new instance of Environment which enclosing
// the new environment.
func NewEnclosedEnvironment(outer *Environment) *Environment {
	env := NewEnvironment()
	env.outer = outer

	return env
}

func (e *Environment) Get(name string) (Object, bool) {
	obj, ok := e.store[name] // look up in inner scope
	if !ok && e.outer != nil {
		obj, ok = e.outer.store[name]
	}

	return obj, ok
}
func (e *Environment) Set(name string, val Object) Object {
	e.store[name] = val

	return val
}
