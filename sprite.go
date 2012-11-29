package gocos2d

type Sprite struct {
	*Node
	isBatchNode bool
}

func (this *Sprite) Init(t Tag) {
	this.Node = new(Node)
	this.Node.Init(t)
}

func (this *Sprite) IsBatchNode() bool {
	return this.isBatchNode
}
