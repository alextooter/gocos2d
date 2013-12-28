package gocos2d

type Layer interface {
	Node
}

type layer struct {
	*node
}

func NewLayer(tag string, z float32) *layer {
	l := &layer{NewNode(tag)}
	return l
}
