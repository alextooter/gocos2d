package gocos2d

type Sprite struct {
	Node
	Tex2d
}

func (s *Sprite) Init(tag string, t *Tex2d) {
	s.Node.Init(tag)
	s.Tex2d = *t
}
