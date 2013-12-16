package gocos2d

type Scene interface {
	Node
	AddLayer(Layer)
}

type scene struct {
	*node
}

func NewScene(tag string) *scene {
	s := &scene{NewNode(tag)}
	s.SetAnchor(0.5, 0.5)
	return s
}

func (s *scene) AddLayer(l Layer) {
	l.SetParent(s)
	s.node.AddChild(l.Tag(), l)

}
