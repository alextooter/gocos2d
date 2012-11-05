package gocos2d

import (
	"bytes"
	"github.com/jteeuwen/glfw"
)

type texture2d struct {
	texture *glfw.Image
}

func (tex *texture2d) Dispose() {
	tex.texture.Release()
}
func (tex *texture2d) Dimensions() (w, h float32) {
	w = float32(tex.texture.Width())
	h = float32(tex.texture.Height())
	return
}
func (tex *texture2d) Scale(scale float32) {
	w, h := tex.Dimensions()
	tex.texture.SetWidth(int(scale * w))
	tex.texture.SetHeight(int(scale * h))
}

func (tex *texture2d) Copy() (*texture2d, error) {
	var b bytes.Buffer
	data := tex.texture.Data()
	b.Read(data)
	return Texture2d(&b)

}
