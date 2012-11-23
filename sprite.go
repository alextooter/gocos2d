package gocos2d

type (
	Sprite_ interface {
	}
	Sprite struct {
		*Node
		Image *Texture2d
	}
)
