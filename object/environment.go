package object

func NewEnvironment() *Environment {
	s := make(map[string]Object)
	return &Environment{store: s}
}

type Environment struct {
	store map[string]Object
}

func (e *Environment) Get(name string) (Object, bool) {
	object, ok := e.store[name]
	return object, ok
}

func (e *Environment) Set(name string, val Object) Object {
	e.store[name] = val
	return val
}
