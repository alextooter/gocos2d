package gocos2d

import (
	"code.google.com/p/gocos2d/gl"
	"github.com/jteeuwen/glfw"
)

type window struct {
	title string
	redBit,
	greenBit,
	blueBit,
	alphaBit,
	width,
	height,
	depthBit,
	stencilBit int
	fullscreen bool
}

func (d *director) initGL() {
	if d.glWindow.title == "" {
		d.glWindow.title = "gocos2d"
	}

	var windowed int

	if d.glWindow.fullscreen == true {
		windowed = glfw.Fullscreen
	} else {
		windowed = glfw.Windowed
	}

	if err := glfw.Init(); err != nil {
		panic("GLFW failed to initialize")
	}

	if err := glfw.OpenWindow(
		d.glWindow.width,
		d.glWindow.height,
		d.glWindow.redBit,
		d.glWindow.greenBit,
		d.glWindow.blueBit,
		d.glWindow.alphaBit,
		d.glWindow.depthBit,
		d.glWindow.stencilBit,
		windowed); err != nil {
		panic("GLFW failed to create window.")
	}
	glfw.SetSwapInterval(1)
	glfw.SetWindowTitle(d.glWindow.title)

	if err := gl.Init(); err != nil {
		println("GL failed to initialize")
		Running = false
		return
	}
	glfw.SetWindowSizeCallback(onResize)
}

func (d *director) ClearColor(r, g, b, a float32) {
	gl.ClearColor(gl.Clampf(r), gl.Clampf(g), gl.Clampf(b), gl.Clampf(a))
	gl.Clear(gl.COLOR_BUFFER_BIT)
}
func onResize(w, h int) {
	gl.Viewport(0, 0, gl.Sizei(w), gl.Sizei(h))
	gl.MatrixMode(gl.PROJECTION)
	gl.LoadIdentity()
	gl.Ortho(0.0, 1.0, 0.0, 1.0, -1.0, 1.0)
}
