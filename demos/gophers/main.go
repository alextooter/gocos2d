package main

import (
	"fmt"
	"github.com/mortdeus/gocos2d"
)

var (
	director  = new(gocos2d.Director)
	scene     = new(Scene)
	groundhog = new(Groundhog)
)

func main() {
	Init()
	defer Cleanup()

	x := 5
	for director.Running {
		fmt.Print(".")
		x--
		Update()
		Draw()
		if x <= 0 {
			director.Running = false
			fmt.Println("\nShutting Down")
		}
	}

}

func Init() {
	director.Init()
	scene.Init("InitScene")
	groundhog.Init("groundhog")
	scene.AddChild(groundhog)
	director.Push(scene)
}
func Update() {
	director.Update()
}
func Draw() {
	director.Draw()
}
func Cleanup() {
	director.Cleanup()
	scene.RemoveChild("groundhog")

}
