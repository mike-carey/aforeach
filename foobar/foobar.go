package foobar

// Foo A fake object for testing async foreach
type Foo struct {
	Bar  *Bar
	Name string
}

// Bar A fake object for testing async foreach
type Bar struct {
	Name string
}

// NewFooBarPair Generates a new Foo-Bar pair
func NewFooBarPair(name string) (Foo, Bar) {
	bar := Bar{Name: name}
	foo := Foo{Name: name, Bar: &bar}

	return foo, bar
}
