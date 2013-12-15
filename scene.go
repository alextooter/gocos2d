package gocos2d

type Scene interface {
	Node
}

type scene struct {
	*node
}

func NewScene(tag string) *scene {
	s := &scene{NewNode(tag)}
	s.SetAnchor(0.5, 0.5)
	return s
}

func (s *scene) Update() error {
	for i := range s.children {
		if err := s.children[i].Update(); err != nil {
			return err
		}
	}
	return nil
}

func (s *scene) Draw() error {
	for i := range s.children {
		if err := s.children[i].Draw(); err != nil {
			return err
		}
	}
	return nil
}

func (s *scene) AddLayer(l Layer) {
	l.SetParent(s)
	s.node.AddChild(l.Tag(), l)

}
