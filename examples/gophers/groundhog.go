package main

import (
	"fmt"
	"gocos2d.org"
	"os"
)

type groundhog struct {
	gocos2d.Sprite
}

func NewGroundhog() (*groundhog, error) {
	fd := os.Open("img/groundhog.png")
	return groundhog{gocos2d.NewSprite("groundhog", "img/groundhog.png")}
}

func (gh *groundhog) Update() {
	fmt.Println("groundhog update")
}
func (gh *groundhog) Draw() {
	fmt.Println("groundhog draw")
}
func (gh *groundhog) OnEnter() {
	fmt.Println("groundhog onenter")
}
func (gh *groundhog) OnExit() {
	fmt.Println("groundhog onexit")
}
