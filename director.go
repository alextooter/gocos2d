package gocos2d

/*The Director is a singleton that is designed to simplify how you manage your game.
It is used to initialize a openGL context and It maintains a stack of scenes.
You are able to update and draw your scenes directly by calling Director.Update
and Director.Draw in your game loop.*/
type Director struct {
	Running bool
	*ActionManager
	*Scheduler
	currentScene IScene
	stack        []IScene
}

func (this *Director) Init() {
	this.Running = true
	this.ActionManager = new(ActionManager)
	this.Scheduler = new(Scheduler)
	this.stack = make([]IScene, 0)
}
func (this *Director) Push(s IScene) {
	this.stack = append(this.stack, s)
	this.currentScene = s

}
func (this *Director) Pop() IScene {
	this.stack = this.stack[:len(this.stack)-1]
	defer func() {
		if (len(this.stack) - 1) > 0 {
			this.currentScene = this.stack[len(this.stack)-1]
		}
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
