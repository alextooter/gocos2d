package gocos2d

import (
	"log"
)

type (
	Node struct {
		Anchor   Anchor
		Position Position
		Rotation Rotation
		Camera   Camera
		Skew     Skew
		Scale    Scale
		tag      Tag
		Z        ZOrder
		Bbox     BoundingBox
		Grid     Grid
		parent   INode
		children *children
	}
	children []INode
)

func (this *Node) Node_() *Node {
	return this
}

func (this *Node) Init(id Tag) {
	this.tag = id
	tmp := make(children, 0)
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
		id, this.Node_().tag)
	return nil
}
func (this *Node) RemoveChild(id Tag) {
	for i, n := range *this.children {
		if n.Node_().tag == id {
			tmp := make(children, len(*this.children)-1)
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
		id, this.Node_().tag)
}
