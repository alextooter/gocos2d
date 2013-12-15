//http://mortdeus.mit-license.org/
package gocos2d

type Action interface {
	Step()
	Stop()
	Update()
	Func(Node)
}
type action struct {
	running bool
	fn      func(Node)
}

type ActionManager interface {
	Register(string, Action)
	Call(string, Node)
	Run(func(Node))

	Pause()
	Resume()
}
type actionManager struct {
	curAction Action
	actions   []Action
	lookup    map[string]Action
}

func NewActionManager() ActionManager {
	return &actionManager{
		actions: make([]Action, 0),
		lookup:  make(map[string]Action, 0)}
}
func (am *actionManager) Run(fn func(Node)) {
	am.actions = append(am.actions, NewAction(fn))
}
func (am *actionManager) Register(tag string, a Action) {
	am.lookup[tag] = a
}

func (am *actionManager) Call(tag string, n Node) {
	am.lookup[tag].Func(n)
}
func (am *actionManager) Pause() {

}
func (am *actionManager) Resume() {

}
func NewAction(fn func(Node)) Action {
	return &action{fn: fn}
}
func (a *action) Step() {

}
func (a *action) Stop() {

}
func (a *action) Update() {

}
func (a *action) Func(n Node) {
	a.fn(n)
}
