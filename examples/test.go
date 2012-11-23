package main

import (
	"github.com/mortdeus/gocos2d"
)

var director = new(gocos2d.Director)

type A struct {
	*gocos2d.Node
}
type B struct {
	*gocos2d.Node
}

func main() {
	a, b := new(A), new(B)
	a.Init()
	b.Init()
	a.AddChild(b)
	a.Cleanup()
}
