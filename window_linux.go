//http://mortdeus.mit-license.org/
package gocos2d

import (
	glfw "github.com/go-gl/glfw3"
	gl "github.com/mortdeus/egles/es2"
	"log"
	"runtime"
)

type window struct {
	*glfw.Window
	Width, Height int
}

func errorCallback(err glfw.ErrorCode, desc string) {
	log.Printf("%v: %v\n", err, desc)
}
func init() {
	runtime.LockOSThread()
}
func (w *window) init() {
	glfw.SetErrorCallback(errorCallback)
	if !glfw.Init() {
		println("glfw init failure")
	}
	if w.Width < 150 || w.Height == 150 {
		w.Width = 680
		w.Height = 480
	}
	glfw.WindowHint(glfw.ClientApi, glfw.OpenglEsApi)
	glfw.WindowHint(glfw.ContextVersionMajor, 2)
	glfw.WindowHint(glfw.ContextVersionMinor, 0)
	var err error
	w.Window, err = glfw.CreateWindow(w.Width, w.Height, Director.AppId, nil, nil)
	if err != nil {
		panic(err)
	}
	w.MakeContextCurrent()
	glfw.SwapInterval(1)
}

func (w *window) reshape(width, height int) {
	w.Width, w.Height = width, height
	gl.Viewport(0, 0, width, height)
}
func (w *window) cleanup() {
	w.Destroy()
	glfw.Terminate()
}

func (w *window) pollEvents() {
	glfw.PollEvents()
}
