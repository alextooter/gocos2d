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
