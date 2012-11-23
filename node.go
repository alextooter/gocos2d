package gocos2d

type (
	Node_ interface {
		ID() uint
		Init()
		OnEnter()
		OnExit()
		Cleanup()
		Draw()
		Visit()
		Transform()
		Schedule(string, func(), int)
		Unschedule(string)
		RunAction(Action_)
		StopAction(Action_)
		AddChild(Node_)
		RemoveChild(uint)
		ConvertTo()
	}
	Node struct {
		Id       uint
		Position Position
		Z        int
		Camera   Camera
		Anchor   Anchor
		Parent   Node_
		Children map[uint]Node_
	}
)

func (this *Node) ID() uint {
	return this.Id
}

func (this *Node) Init() {
	this = new(Node)
	this.Id = GID
	GID++
	this.Children = make(map[uint]Node_)
}
func (this *Node) Cleanup() {
	for _, n := range this.Children {
		n.Cleanup()
	}
	this = nil
}
func (this *Node) AddChild(child Node_) {
	this.Children[child.ID()] = child
}
func (this *Node) RemoveChild(id uint) {
	delete(this.Children, id)
}
func (this *Node) Draw() {

}
func (this *Node) Visit() {

}
func (this *Node) Transform() {

}
func (this *Node) Schedule(name string, cb func(), interval int) {

}
func (this *Node) Unschedule(name string) {

}
func (this *Node) RunAction(act Action_) {

}
func (this *Node) StopAction(act Action_) {

}
func (this *Node) ConvertTo() {

}
func (this *Node) OnEnter() {

}
func (this *Node) OnExit() {

}
