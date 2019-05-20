# Abstract Async

## Usage
This package makes use of [genny](https://github.com/cheekybits/genny) to make use of generics.  Use the genny executable to generate a new async function, for example ForEach:

### Command line
```bash
genny -in=$GOPATH/src/github.com/mike-carey/async/foreach.go -out=myforeach.go -pkg mypkg gen "Input=ObjectIn Output=ObjectOut"
```

### Go generate
```golang
//go:generate genny -in=$GOPATH/src/github.com/mike-carey/async/foreach.go -out=myforeach.go -pkg mypkg gen "Input=ObjectIn Output=ObjectOut"
```

## Running tests
The `foobar` is the dummy package while will generate all of the async functions using basic objects so that tests can be built around a concrete object.

```bash
ginkgo ./foobar
```
