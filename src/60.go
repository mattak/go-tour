package main

import (
	"code.google.com/p/go-tour/pic"
	"image"
	color "image/color"
	"math"
)

type Image struct{}

func (i *Image) Bounds() image.Rectangle {
	return image.Rect(0, 0, 300, 300)
}

func (i *Image) ColorModel() color.Model {
	return color.RGBAModel
}

func (i *Image) At(x, y int) color.Color {
	r := uint8(x ^ y)
	g := uint8(x*x + y*y)
	b := uint8(x * y)
	a := uint8(math.Sin(float64(x*y)) * 255)

	return color.RGBA{r, g, b, a}
}

func main() {
	m := Image{} // image.NewRGBA(image.Rect(0, 0, 100, 100))
	pic.ShowImage(&m)
}
