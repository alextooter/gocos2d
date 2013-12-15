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

	hog, err := NewGroundhog()
	if err != nil {
		panic(err)
	}
	lvl1 := NewLevel()
	lvl1.AddChild(hog)
	dir.Push(lvl1)
}
