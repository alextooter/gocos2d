package main

import(
	
	"gocos2d"
	"github.com/jteeuwen/glfw"
)

func main() {
	//Create the director.
	d, _ := gocos2d.CreateDirector(); 
	//Initialize the Director
	d.Init()
	//This will be drawn out into it's own director.Update() and Director.Draw()
	//methods later.
	for glfw.WindowParam(glfw.Opened) == 1 {
                glfw.SwapBuffers()
        }
        	//Tell the director to clean up()
        	d.Cleanup()

	println("Gocos2d executed correctly")
}