package main

import (
	"golang.org/x/tour/pic"
	"image"
	"image/color"
	"math"
)

type Image struct {
	W, H int
}

func (i Image) ColorModel() color.Model {
	return color.RGBAModel
}

func (i Image) Bounds() image.Rectangle {
	return image.Rect(0, 0, i.W, i.H)
}

func (i Image) At(x, y int) color.Color {
	v := uint8((x + y) % math.MaxUint8)
	return color.RGBA{R: v, G: v, B: 255, A: 255}
}

func main() {
	m := Image{10, 10}
	pic.ShowImage(m)
}
