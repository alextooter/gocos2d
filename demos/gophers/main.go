package main

import (
	"fmt"
	"gocos2d.org"
)

var (
	director  = new(gocos2d.Director)
	lvl1      = new(Level)
	groundhog = new(Groundhog)
)

func main() {
	Init()
	for director.Running {
		Update()
		Draw()
	}
	Cleanup()

}
func Init() {
	gocos2d.AppID = "Gophers"
	director.Init()
	lvl1.Init("lvl1")
	groundhog.Init("groundhog0")
	fmt.Println(lvl1.AddChild(groundhog))
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
	lvl1.RemoveChild("groundhog0")

}
