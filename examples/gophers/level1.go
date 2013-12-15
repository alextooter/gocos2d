package main

import "gocos2d.org"

type Level struct {
	gocos2d.Scene
}

func NewLevel() *Level {
	return &Level{gocos2d.NewScene("level")}
}

func (this *Level) Update() {

}
func (this *Level) Draw() {

}
func (this *Level) OnEnter() {

}
func (this *Level) OnExit() {

}
