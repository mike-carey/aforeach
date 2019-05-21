package foobar

// Foo ...
type Foo struct {
	Bar  *Bar
	Name string
}

// Bar ...
type Bar struct {
	Name string
}

// Baz ...
type Baz struct {
	Name string
	Foo  []Foo
}

// NewFooBarPair ...
func NewFooBarPair(name string) (Foo, Bar) {
	bar := Bar{Name: name}
	foo := Foo{Name: name, Bar: &bar}

	return foo, bar
}

// NewFooBarBazGroup ...
func NewFooBarBazGroup(name string) (Foo, Bar, Baz) {
	bar := Bar{Name: name}
	foo := Foo{Name: name, Bar: &bar}
	baz := Baz{Name: name, Foo: []Foo{foo}}

	return foo, bar, baz
}
