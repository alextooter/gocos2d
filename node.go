package gocos2d

type (
	Node struct {
		anchor
		pos
		rot
		camera
		skew
		scale
		tag
		zOrder
		grid
		rect
		parent INode
		children
	}
)

func (n *Node) Init(id string) {
	n.tag = tag(id)
	n.children.init()
}
func (n *Node) Cleanup() {

}
func (n *Node) Update() {

}
func (n *Node) Draw() {

}
func (n *Node) OnEnter() {

}
func (n *Node) OnExit() {

}
func (n *Node) Visit() {

}

func (n *Node) Transform(uint) {

}
func (n *Node) ConvertTo(uint) {

}
