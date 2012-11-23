package gocos2d

/*The Director is a singleton that is designed to simplify how you manage your game.
It is used to initialize a openGL context and It maintains a stack of scenes.
You are able to update and draw your scenes directly by calling Director.Update
and Director.Draw in your game loop.*/
type Director struct {
}

func (d *Director) Init() {
}
func (d *Director) Cleanup() {
}
func (d *Director) Update() {
}
func (d *Director) Draw() {
}
