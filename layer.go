package gocos2d

type Layer interface {
	Node
}

type layer struct {
	*node
}

func NewLayer(tag string, z float64) *layer {
	l := &layer{NewNode(tag)}
	l.SetZ(z)
	return l
}
