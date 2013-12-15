package gocos2d

import "sync"

type Node interface {
	sync.Locker
	Update() error
	Draw() error
	OnEnter() error
	OnExit() error
	Cleanup() error
	Visit() error
	Transform(uint) error
	ConvertTo(uint) error

	AddChild(string, Node)
	GetChild(string) Node
	RemoveChild(string)
}
type node struct {
	id string
	rect
	rot
	camera
	skew
	scale
	zOrder
	//grid

	parent   Node
	children []Node
	lookup   map[string]int
	sync.Mutex
}

func NewNode(id string) *node {
	return &node{
		id:       id,
		children: make([]Node, 0),
		lookup:   make(map[string]int, 0)}
}

func (n *node) Cleanup() error {
	n.Lock()
	defer n.Unlock()

	return nil
}
func (n *node) Update() error {
	n.Lock()
	defer n.Unlock()
	return nil
}
func (n *node) Draw() error {
	n.Lock()
	defer n.Unlock()
	return nil
}
func (n *node) OnEnter() error {
	n.Lock()
	defer n.Unlock()
	return nil
}
func (n *node) OnExit() error {
	n.Lock()
	defer n.Unlock()
	return nil
}
func (n *node) Visit() error {
	n.Lock()
	defer n.Unlock()
	return nil
}

func (n *node) Transform(uint) error {
	n.Lock()
	defer n.Unlock()
	return nil
}
func (n *node) ConvertTo(uint) error {
	n.Lock()
	defer n.Unlock()
	return nil
}

func (n *node) AddChild(tag string, child Node) {
	n.Lock()
	defer n.Unlock()

	if _, exists := n.lookup[tag]; !exists {
		n.lookup[tag] = len(n.children)
		n.children = append(n.children, child)
	}
}
func (n *node) GetChild(tag string) Node {
	n.Lock()
	defer n.Unlock()
	if i, exists := n.lookup[tag]; exists {
		return n.children[i]
	}
	return nil
}

func (n *node) RemoveChild(tag string) {
	n.Lock()
	defer n.Unlock()

	/*
		Rather than physically remove the child from the list;
		this function just replaces the child with a tombstone.
		The reason we have to tombstone is because we rely on a
		map to quickly lookup the index of our children stored in
		our slice.

		One major setback of tombstoning like this is the fact that
		the node.children only grows longer and longer in algorithms
		that frequently add and remove children. Therefore we at some
		point we should consider implementing more sophisticated memory
		handling logic here.
	*/
	if i, exists := n.lookup[tag]; exists {
		n.children[i] = nil
		delete(n.lookup, tag)
	}
}
