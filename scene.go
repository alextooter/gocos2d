package gocos2d

type Scene struct {
	*Node
	isCurrent bool
	
}

func (this *Scene) IsCurrent() bool {
	return this.isCurrent
}
