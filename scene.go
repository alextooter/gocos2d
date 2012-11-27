package gocos2d

type Scene struct {
	*Node
}

func (this *Scene) Init(t Tag) {
	this.Node = new(Node)
	this.Node.Init(t)
}
