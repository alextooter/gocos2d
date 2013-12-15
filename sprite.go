package gocos2d

import (
	_ "code.google.com/p/vp8-go/webp"
	//	gl "github.com/mortdeus/egles/es2"
	"image"
	_ "image/jpeg"
	_ "image/png"
	"io"
)

type Sprite interface {
	Node
}
type sprite struct {
	*node
	image.NRGBA
}

func NewSprite(id string, r io.Reader) (*sprite, error) {
	//TODO(mortdeus): Implement texture2d cache lookup.
	img, _, err := image.Decode(r)
	if err != nil {
		return nil, err
	}
	s := &sprite{NewNode(id), *image.NewNRGBA(img.Bounds())}
	bounds := s.Bounds()
	for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
		for x := bounds.Min.X; x < bounds.Max.X; x++ {
			s.Set(x, y, img.At(x, y))
		}
	}

	return s, nil
}

func (s *sprite) Draw() error {
	return nil
}
