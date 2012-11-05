package main

import (
	"code.google.com/p/gocos2d"
	"os"
)

var director = gocos2d.Director()
var bank = gocos2d.ImageBank()

func main() {
	director.Start()
	load()
	director.ClearColor(0, .25, 1, 0)
	for gocos2d.Running {
		director.Update()
		director.Draw()
	}
	director.Cleanup()
}

func load() {
	bank.Cache("groundhog.png", "groundhog")
	file, err := os.Create("groundhog.tga")
	if err != nil {
		panic(err)
	}
	n, err := file.Write(bank.Get("groundhog").Bytes())
	if err != nil {
		panic(err)
	}
	println("tga written to disk: wrote: ", n, " bytes")
}
