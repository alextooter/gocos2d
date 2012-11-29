package gocos2d

type Action struct {
	id     Tag
	isDone bool
	target *INode
}

type ActionManager struct {
}

func (this *ActionManager) Run(a *Action) {

}

func (this *Action) Action() *Action {
	return this
}
func (this *Action) Init(id Tag) {
	this.id = id
}
func (this *Action) Step() {

}
func (this *Action) Stop() {

}
func (this *Action) Update() {

}
