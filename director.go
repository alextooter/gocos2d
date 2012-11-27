package gocos2d

/*The Director is a singleton that is designed to simplify how you manage your game.
It is used to initialize a openGL context and It maintains a stack of scenes.
You are able to update and draw your scenes directly by calling Director.Update
and Director.Draw in your game loop.*/
type Director struct {
	*ActionManager
	*Scheduler
	currentScene *Scene
	stack        []*Scene
}

func (this *Director) Init() {
	this.ActionManager = new(ActionManager)
	this.Scheduler = new(Scheduler)
	this.stack = make([]*Scene, 0)
}
func (this *Director) Push(s *Scene) {
	if this.currentScene != nil {
		this.currentScene.isCurrent = false
	}
	this.stack = append(this.stack, s)
	this.currentScene = s
}
func (this *Director) Pop() *Scene {
	this.stack = this.stack[:len(this.stack)-1]
	defer func() {
		this.currentScene = this.stack[len(this.stack)]
	}()
	return this.currentScene
}
func (this *Director) Destroy(n INode) {
	n.Cleanup()
}
func (this *Director) Pause() {

}
func (this *Director) Unpause() {

}
func (this *Director) Cleanup() {
}
func (this *Director) Update() {
}
func (this *Director) Draw() {
}
