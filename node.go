package gocos2d

import (
	"log"
)

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

func (this *Node) Node_() *Node {
	return this
}

func (this *Node) Init(id Tag) {
	this.tag = id
	tmp := make(Children, 0)
	this.children = &tmp

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

func (this *Node) AddChild(n INode) {
	*this.children = append(*this.children, n)
}
func (this *Node) GetChild(id Tag) INode {
	for _, n := range *this.children {
		if n.Node_().tag == id {
			return n
		}
	}
	log.Panicf(
		"gocos2d: %s wasnt a found child tag in %s, shutting down",
		id, *this.Tag())
	return nil
}
func (this *Node) RemoveChild(id Tag) {
	for i, n := range *this.children {
		if n.Node_().tag == id {
			tmp := make(Children, len(*this.children)-1)
			_ = copy(tmp, (*this.children)[:i])
			if i != cap(*this.children) {
				_ = copy(tmp, (*this.children)[i+1:])
			}
			this.children = &tmp
			return
		}
	}
	log.Panicf(
		"gocos2d: %s wasnt a found child tag in %s, shutting down",
		id, *this.Tag())
}
