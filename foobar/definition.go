// +build foobar

package foobar

//go:generate genny -in=../foreach.go -out=foobar_foreach.go -pkg foobar gen "Input=Foo Output=Bar"
//go:generate genny -in=../map-one-to-many.go -out=foobar_map-one-to-many.go -pkg foobar gen "Input=Baz Output=Bar"
