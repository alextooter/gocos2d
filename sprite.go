package gocos2d

type Sprite struct {
	Node
	isBatchNode bool
}

func (this *Sprite) Init(id string) {
	this.Node.Init(id)
}

func (this *Sprite) IsBatchNode() bool {
	return this.isBatchNode
}
