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
	a, b := new(A).Init("NodeA"), new(B).Init("NodeB")
	a.AddChild(b)
	c := a.Children.Lookup("NodeB").Node
	println(b.Tag() == c.Tag())
	a.RemoveChild("NodeB")
	println(a.Children.Len())
	a.Cleanup()
}
