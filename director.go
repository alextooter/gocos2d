package gocos2d

import (
	"github.com/jteeuwen/glfw"
	"time"
)

type director struct {
	stack    sceneStack
	glWindow window
}

func (d *director) Push(s *Scene) {
	d.stack.push(s)
}
func (d *director) Replace(s *Scene) {
	d.stack.replace(s)
}
func (d *director) Pop() {
	d.stack.pop()
}
func (d *director) Get() *Scene {
	return d.stack.currentScene
}
func (d *director) SetWindowParams(title string, width, height int, b Bits, fullscreen bool) {
	if title == "" {
		d.glWindow.title = "Gocos2D"
	} else {
		d.glWindow.title = title
	}
	d.glWindow.width = width
	d.glWindow.height = height
	d.glWindow.redBit = b.red
	d.glWindow.greenBit = b.green
	d.glWindow.blueBit = b.blue
	d.glWindow.alphaBit = b.alpha
	d.glWindow.depthBit = b.depth
	d.glWindow.stencilBit = b.stencil
	d.glWindow.fullscreen = fullscreen
}
func (d *director) Start() {
	d.initGL()
	d.ClearColor(0, 0, 0, 0)
}
func (d *director) Cleanup() {
	println("Shutting Down")
	time.Sleep(500 * time.Millisecond)
	glfw.CloseWindow()
	glfw.Terminate()
	println("Gocos2D ended correctly.")
}
func (d *director) Update() {
	if glfw.WindowParam(glfw.Opened) != 1 {
		Running = false
	}
}
func (d *director) Draw() {
	if glfw.WindowParam(glfw.Opened) == 1 {
		glfw.SwapBuffers()
	}
}
