package main

import "code.google.com/p/go-tour/pic"
import "fmt"

func Pic(dx, dy int) [][]uint8 {
	image := make([][]uint8, dy, dy)
	for iy := 0; iy < dy; iy++ {
		image[iy] = make([]uint8, dx, dx)
	}

	for iy := 0; iy < dy; iy++ {
		for ix := 0; ix < dx; ix++ {
			// image[iy][ix] = uint8(ix^iy)

			vx := float64(ix)
			vy := float64(iy)
			image[iy][ix] = uint8(vx*vx + vy*vy)
			// image[iy][ix] = uint8(vx * vy)
			// image[iy][ix] = uint8(vx * vy - vx * vx * 0.3 - vy * vy * 0.3)
		}
	}

	return image
}

func main() {
	pic.Show(Pic)
}
