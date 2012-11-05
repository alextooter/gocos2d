package gocos2d

import (
	"github.com/jteeuwen/glfw"
	"image"
)

var (
	Running              = true
	directorInitialized  = false
	imageBankInitialized = false
)

/*The Director is a singleton that is designed to simplify how you manage your game.
It is used to initialize a openGL context and It maintains a stack of scenes.
You are able to update and draw your scenes directly by calling Director.Update
and Director.Draw in your game loop.*/
func Director() *director {
	if !directorInitialized {
		d := new(director)
		d.stack.init()
		directorInitialized = true
		return d
	}
	println("You can only have one Director.")
	Running = false
	return nil
}
func ImageBank() *imageBank {
	if !imageBankInitialized {
		bank := new(imageBank)
		bank.cache = make(map[string]*bytes.Buffer)
		imageBankInitialized = true
		return bank
	}
	println("You can only have one ImageBank.")
	return nil
}
func Texture2d(img *image.Image) (*texture2d, error) {
	tex := new(texture2d)
	var err error

	//TODO: copy image data into a texture2D
	//that takes into account power of 2 dimensions expected by opengles2

	if err != nil {
		return nil, err
	}
	return tex, nil
}
func Sprite(img *bytes.Buffer) *sprite {
	s := new(sprite)
	var err error
	s.Texture2d, err = Texture2d(img)
	if err != nil {
		println(err)
	}
	return s
}
