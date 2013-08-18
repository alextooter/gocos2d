package gocos2d

import glfw "github.com/go-gl/glfw3"

/*The Director is a singleton that is designed to simplify how you manage your game.
It is used to initialize a openGL context and It maintains a stack of scenes.
You are able to update and draw your scenes directly by calling Director.Update
and Director.Draw in your game loop.*/
type Director struct {
	window
	Running bool
	*ActionManager
	*Scheduler
	currentScene IScene
	stack        []IScene
}

func (d *Director) Init() {
	d.Running = true
	d.ActionManager = new(ActionManager)
	d.Scheduler = new(Scheduler)
	d.stack = make([]IScene, 0)
	d.window.init()
}
func (d *Director) Push(s IScene) {
	d.stack = append(d.stack, s)
	d.currentScene = s

}
func (d *Director) Pop() IScene {
	d.stack = d.stack[:len(d.stack)-1]
	defer func() {
		if (len(d.stack) - 1) > 0 {
			d.currentScene = d.stack[len(d.stack)-1]
		}
	}()
	return d.currentScene
}
func (d *Director) Destroy(n INode) {
	n.Cleanup()
}
func (d *Director) Pause() {

}
func (d *Director) Unpause() {

}
func (d *Director) Cleanup() {
	d.window.cleanup()
}
func (d *Director) Update() {
	if d.window.win.ShouldClose() {
		d.Running = false
	}
	glfw.PollEvents()
}
func (d *Director) Draw() {
}
