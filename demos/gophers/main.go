package main

import (
	"github.com/mortdeus/gocos2d"
)

var (
	director  = new(gocos2d.Director)
	lvl1      = new(Level)
	groundhog = new(Groundhog)
)

func main() {
	Init()
	defer Cleanup()
	for director.Running {
		Update()
		Draw()
	}

}
func Init() {
	director.Init()
	lvl1.Init("lvl1")
	lvl1.AddChild(groundhog)
	director.Push(lvl1)
}
func Update() {
	director.Update()
}
func Draw() {
	director.Draw()
}
func Cleanup() {
	director.Cleanup()
	lvl1.RemoveChild("Groundhog")

}
