package gocos2d

import (
	_ "code.google.com/p/vp8-go/webp"
	//gl "github.com/mortdeus/egles/es2"
	"image"
	_ "image/jpeg"
	_ "image/png"
	"io"
)

type Tex2d struct {
	*image.NRGBA
}

func NewTex2d(r io.Reader) (*Tex2d, error) {
	//TODO: Implement texture2d cache lookup.
	img, _, err := image.Decode(r)
	if err != nil {
		return nil, err
	}
	tex := &Tex2d{image.NewNRGBA(img.Bounds())}
	bounds := tex.Bounds()
	for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
		for x := bounds.Min.X; x < bounds.Max.X; x++ {
			tex.Set(x, y, img.At(x, y))
		}
	}
	return tex, nil
}
