package gocos2d

import (
	"errors"
)

/*The Director is a singleton that is designed to simplify how you manage your game.
It is used to initialize a openGL context and It maintains a stack of scenes.
You are able to update and draw your scenes directly by calling Director.Update
and Director.Draw in your game loop.*/
type Director struct {
	currentScene	*Scene
	sceneStack	[]Scene
	glWindow	window
}

//You can only have one instance of the director. 
var directorInitialized = false

//Director should be declared as a global variable to insure it does not fall out of scope.
func CreateDirector() (director, error) {
	var d *director
	if directorInitialized != true {
		d = new(director); if d == nil {
			panic("Director creation has failed.")
		}
		directorInitialized = true
		if err := d.initSceneStack(); err != nil {
			panic("Allocating the scene stack failed")
		}
		return d, nil
	}
	return nil, errors.New("You can only have one Director.")
}

/*You add scenes to the scene stack, however to reduce your game's memory overhead; you are not allowed
to add more than 5 scenes. You should never need more than 5 scenes loaded into memory at any given time. If you 
need to implement a UI and you were considering the best option was to create scenes for menus,
HUDS ect; consider using a scene layer instead. 
*/
func (d *director) AddScene(s Scene) error {
	if &d.sceneStack == nil {
		panic("The scene stack isnt initialized.")
	}
	if len(d.sceneStack) < cap(d.sceneStack) {
		d.sceneStack = append(d.sceneStack, s)
		return nil
	}
		panic("The Scene stack is full.")
}
func (d *director) initSceneStack() error {
	if d.sceneStack != nil {
		panic("You can only have one scene stack.")
	}
	if d.sceneStack = make([]Scene, 5); d.sceneStack == nil {
		panic("The scene stack failed to initialize")
	}
	return nil
}

/*Director.Init() creates an openGL window context and sets all the default values*/
func (d *director) Init() error {
	rgba := [4]int{0, 0, 0, 0}
	err := d.glWindow.init("", 0, 0, rgba, true)
	if err != nil {
		return err
	}
	return nil
}
