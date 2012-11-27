package main

import (
	"github.com/mortdeus/gocos2d"
)

type (
	Scene struct {
		*gocos2d.Scene
	}
	Groundhog struct {
		*gocos2d.Sprite
	}
)

func (this *Scene) Init(t gocos2d.Tag) {
	this.Scene = new(gocos2d.Scene)
	this.Scene.Init(t)
}
func (this *Groundhog) Init(t gocos2d.Tag) {
	this.Sprite = new(gocos2d.Sprite)
	this.Sprite.Init(t)
}
