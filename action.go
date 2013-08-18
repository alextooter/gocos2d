package gocos2d

type Action struct {
	id     string
	isDone bool
	target INode
}

type ActionManager struct {
}

func (am *ActionManager) Run(a *Action) {

}

func (a *Action) Action() *Action {
	return a
}
func (a *Action) Init(id string) {
	a.id = id
}
func (a *Action) Step() {

}
func (a *Action) Stop() {

}
func (a *Action) Update() {

}
