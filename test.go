package main

const Foo int = 0

type X struct {
}

type Y interface {
	Foo(s string)
}

type Z interface {
	Bar(s string)
}

func Bar() {
	type XXX interface{}
}
