package gocos2d

import (
	"io"
)

type Sprite interface {
	Node
}
type sprite struct {
	*node
	*texture
}

func NewSprite(tag string, r io.Reader) *sprite {
	return &sprite{NewNode(tag), newTexture(r)}
}

func (s *sprite) Draw() error {
	s.Render(s)
	return nil
}
