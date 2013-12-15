//http://mortdeus.mit-license.org/
package main

import (
	"fmt"
	"gocos2d.org"
	"os"
)

var dir = gocos2d.Dirctor
var lvl1 = new(Level)
var groundhog = new(Groundhog)

func main() {
	Init()
	for dir.Running {
		dir.Update()
		dir.Draw()
	}
	dir.Cleanup()

}
func Init() {
	gocos2d.AppID = "Gophers"
	dir.Init()
	lvl1.Init("lvl1")
	
	hog, err := Groundhog{gocos2d.NewSprite("groundhog0")}

	dir.Push(lvl1)
}
