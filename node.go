package gocos2d

type (
	Node struct {
		anchor   Anchor
		position Position
		rotation Rotation
		camera   Camera
		skew     Skew
		scale    Scale
		tag      Tag
		z        ZOrder
		bbox     BoundingBox
		grid     Grid
		parent   INode
		children *Children
	}
	Children []INode
)

func (this *Node) Init(tag Tag) *Node {
	return new(Node)

}
func (this *Node) Cleanup() {

}
func (this *Node) Update() {

}
func (this *Node) Draw() {

}
func (this *Node) OnEnter() {

}
func (this *Node) OnExit() {

}
func (this *Node) Visit() {

}

func (this *Node) Transform(uint) {

}
func (this *Node) ConvertTo(uint) {

}

func (this *Node) Anchor() *Anchor {
	return &this.anchor
}
func (this *Node) Position() *Position {
	return &this.position
}
func (this *Node) Rotation(bool) *Rotation {
	return &this.rotation
}
func (this *Node) Scale() *Scale {
	return &this.scale
}
func (this *Node) Skew() *Skew {
	return &this.skew
}
func (this *Node) Tag() *Tag {
	return &this.tag
}
func (this *Node) ZOrder() *ZOrder {
	return &this.z
}
func (this *Node) Children() *Children {

	return this.Children()
}
func (this *Node) Parent() INode {
	return this.parent
}
func (this *Node) Grid() *Grid {
	return &this.grid
}
func (this *Node) Camera() *Camera {
	return &this.camera
}
func (this *Node) BoundingBox() *BoundingBox {
	return &this.bbox
}
