package main

import (
	"gocos2d.org/sdk"
	"os"
)

type groundhog struct {
	gocos2d.Sprite
}

func NewGroundhog() *groundhog {
	img, err := os.Open("img/gopher.png")
	if err != nil {
		panic(err)
	}
	return &groundhog{gocos2d.NewSprite("groundhog", img)}
}
func (g *groundhog) Update() error {
	return nil
}
func (g *groundhog) Draw() error {
	return nil
}
func (g *groundhog) OnEnter() error {
	return nil
}
func (g *groundhog) OnExit() error {
	return nil
}
