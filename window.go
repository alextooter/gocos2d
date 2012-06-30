package gocos2d

import (
//	"errors"
	"glfw"
	gl "gogl/gl21"
)

type window struct {
	title      string
	width      int
	height     int
	rgba       [4]int
	depthBit   int
	stencilBit int
	fullscreen   bool
}

func (w *window) setTitle(title string) {
	if title == "" {
		w.title = "Gocos2D"
	} else {
		w.title = title
	}
}
func (w *window) setWidthAndHeight(width, height int) {
	w.width = width
	w.height = height
}
func (w *window) setRGBA(red, green, blue, alpha int) {
	w.rgba[0] = red
	w.rgba[1] = green
	w.rgba[2] = blue
	w.rgba[3] = alpha
}
func (w *window) init(title string, width, height int,  rgba [4]int, isFullscreen bool) error{
	if w = new(window); w != nil{
		panic("Window initialization failed.")
	}
	w.setTitle(title)
	w.setWidthAndHeight(width, height)
	w.setRGBA(rgba[0], rgba[1], rgba[2], rgba[3])
	if isFullscreen{
		w.fullscreen = true
	}
	w.initGL()
	return nil
}
func (w *window) initGL() error{
	var windowed int

	if w.fullscreen == true{
		windowed = glfw.Fullscreen
	}else{
		windowed = glfw.Windowed
		}
	if err := glfw.Init(); err != nil {
		panic("GLFW failed to initialize")
	}
	defer glfw.Terminate()

	glfw.OpenWindowHint(glfw.WindowNoResize, 1)
	if err := glfw.OpenWindow(
				w.width,
				w.height,
				w.rgba[0],
				w.rgba[1],
				w.rgba[2],
				w.rgba[3],
				w.depthBit,
				w.stencilBit,
				windowed); err != nil {
		panic("GLFW failed to create window.")	
	}
	defer glfw.CloseWindow()
	
	glfw.SetSwapInterval(1)
	glfw.SetWindowTitle(w.title)

	if err := gl.Init(); err != nil {
		panic("GL failed to initialize")
	}
	return nil
}
