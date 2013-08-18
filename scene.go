package gocos2d

type Scene struct {
	Node
}

func (this *Scene) Init(id string) {
	this.Node.Init(id)
}
