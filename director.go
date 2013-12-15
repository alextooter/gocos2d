//http://mortdeus.mit-license.org/
package gocos2d

import (
	"sync"
)

/*The Director is a singleton that is designed to simplify how you manage your game.
It is used to initialize a openGL context and It maintains a stack of scenes.
You are able to update and draw your scenes directly by calling Director.Update
and Director.Draw in your game loop.*/
var Director *director

type director struct {
	AppId string
	sync.Mutex
	window
	Running bool
	ActionManager
	Scheduler
	curScene Scene
	stack    []Scene
}

func (d *director) Init() {
	d.Lock()
	defer d.Unlock()

	if d.AppId == "" {
		d.AppId = "gocos2d"
	}
	if d.ActionManager == nil {
		d.ActionManager = NewActionManager()
	}
	if d.Scheduler == nil {
		d.Scheduler = NewScheduler()
	}
	d.Running = true
	d.stack = make([]Scene, 0)
	d.window.init()

}
func (d *director) Push(s Scene) {
	d.Lock()
	defer d.Unlock()
	d.stack = append(d.stack, s)
	d.curScene = s

}
func (d *director) Pop() Scene {
	d.Lock()
	defer d.Unlock()
	d.stack = d.stack[:len(d.stack)-1]
	defer func() {
		if (len(d.stack) - 1) > 0 {
			d.curScene = d.stack[len(d.stack)-1]
		}
	}()
	return d.curScene

}
func (d *director) Destroy(n Node) {
	d.Lock()
	defer d.Unlock()
	n.Cleanup()
}
func (d *director) Pause() {
	d.Lock()
	defer d.Unlock()
}
func (d *director) Unpause() {
	d.Lock()
	defer d.Unlock()
}
func (d *director) Cleanup() {
	d.Lock()
	defer d.Unlock()

	d.window.cleanup()
}
func (d *director) Update() {
	d.Lock()
	defer d.Unlock()
	d.curScene.Update()
	if d.window.ShouldClose() {
		d.Running = false
	}
	d.window.pollEvents()
}
func (d *director) Draw() {
	d.Lock()
	defer d.Unlock()
}
func (d *director) SetActionManager(am ActionManager) {
	d.Lock()
	defer d.Unlock()
	d.ActionManager = am
}
func (d *director) SetScheduler(s Scheduler) {
	d.Lock()
	defer d.Unlock()
	d.Scheduler = s
}
