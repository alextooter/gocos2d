package gocos2d

import (
	_ "code.google.com/p/vp8-go/webp"
	"errors"
	gl "github.com/mortdeus/egles/es2"
	"image"
	_ "image/jpeg"
	_ "image/png"
	"io"
)

type texture struct {
	id         uint
	data       []byte
	shaderProg uint
	pxW, pxH   int
	S, T       float32

	pos, texCoord int
}

func newTexture(r io.Reader) *texture {
	//TODO(mortdeus): Implement texture2d cache lookup.
	img, _, err := image.Decode(r)
	if err != nil {
		panic(err)
	}
	nrgba, ok := img.(*image.NRGBA)
	if !ok {
		panic(errors.New("texture must be an NRGBA image"))
	}

	t := new(texture)
	// flip image: first pixel is lower left corner
	t.pxW, t.pxH = nrgba.Bounds().Dx(), nrgba.Bounds().Dy()
	t.data = make([]byte, t.pxW*t.pxH*4)
	lineLen := t.pxW * 4
	dest := len(t.data) - lineLen
	for src := 0; src < len(nrgba.Pix); src += nrgba.Stride {
		copy(t.data[dest:dest+lineLen], nrgba.Pix[src:src+nrgba.Stride])
		dest -= lineLen
	}
	gl.GenTextures(1, &t.id)
	gl.ActiveTexture(gl.TEXTURE0)
	gl.BindTexture(gl.TEXTURE_2D, t.id)

	gl.TexImage2D(
		gl.TEXTURE_2D, 0, gl.RGBA,
		t.pxW, t.pxH, 0, gl.RGBA, gl.UNSIGNED_BYTE, t.data)

	//TODO(mortdeus): Precompile this shader.
	t.shaderProg = Program(
		NewShader(POSITION_TEXTURE_VERT, gl.VERTEX_SHADER),
		NewShader(POSITION_TEXTURE_FRAG, gl.FRAGMENT_SHADER))

	t.pos = gl.GetUniformLocation(t.shaderProg, "a_pos")
	t.texCoord = gl.GetUniformLocation(t.shaderProg, "a_texCoord")
	return t
}

func (t *texture) Render(n Node) {
	/*x, y, h, w := n.

	vertices := [8]float32{
		x, y,
		x + w, y,
		x, y + h,
		x + w, y + h}

	coordinates := [8]float32{
		0, t.T,
		t.S, t.T,
		0, 0,
		t.S, 0}

	gl.EnableVertexAttribArray(t.pos)
	gl.EnableVertexAttribArray(t.texCoord)

	gl.VertexAttribPointer(t.pos, 2, gl.FLOAT, false, 0, &vertices)
	gl.VertexAttribPointer(t.texCoord, 2, gl.FLOAT, false, 0, &coordinates)

	gl.Clear(gl.COLOR_BUFFER_BIT)
	gl.DrawArrays(gl.TRIANGLE_STRIP, 0, 4)
	gl.Flush()
	gl.Finish()
	*/
}
