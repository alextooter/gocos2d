package gocos2d

type Scene interface {
	Node
}

type scene struct {
	node
}

func NewScene(tag string) *scene {
	return &scene{*NewNode(tag)}
}
