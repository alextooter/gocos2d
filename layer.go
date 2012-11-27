package gocos2d

type Layer struct {
	*Node
	isParallax bool
}

func (this *Layer) IsParallax() *bool {
	return &this.isParallax
}
