// +build foobar

package foobar

//go:generate genny -in=../foreach.go -out=foobar_foreach.go -pkg foobar gen "Input=Foo Output=Bar"
