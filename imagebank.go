package gocos2d

import (
	"bytes"
	"errors"
	"image"
	_ "image/gif"
	_ "image/jpeg"
	_ "image/png"
	"os"
)

type imageBank struct {
	cache     map[string]*image.Image
	itemCount int
}

func (bank *imageBank) Cache(fn, key string) error {
	exists := bank.Check(key)
	if exists {
		return errors.New("File already exists")
	}

	file, err := os.Open(fn)
	if err != nil {
		return err
	}
	m, _, err := image.Decode(file)

	img, ok := m.(*image.NRGBA)
	if !ok {
		b := m.Bounds()
		i := image.NewNRGBA(b)
		draw.Draw(i, b, m, b.Min, draw.Src)
		img = i
	}

	bank.cache[key] = &img
	return nil
}
func (ib *imageBank) Check(key string) bool {
	_, ok := ib.cache[key]
	return ok
}
func (ib *imageBank) Remove(key string) error {
	if ib.Check(key) {
		delete(ib.cache, key)
	}

	return errors.New("There isnt any image data stored under that key.")
}
func (ib *imageBank) Get(fn string) (img *image.Image) {
	img = ib.cache[fn]
	return
}
func (ib *imageBank) Flush() {
	for key := range ib.cache {
		delete(ib.cache, key)
	}
}
