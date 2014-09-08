package main

import (
	"fmt"
	"math"
)

func Sqrtx(x float64) (z float64) {
	z = x
	za := z + 1e4
	for math.Abs(za-z) > 1e-18 {
		println("ok", z, za)
		za = z
		z = z - (z*z-x)/(2*z)
	}
	return
}

func main() {
	fmt.Println(Sqrtx(0.1))
	fmt.Println(math.Sqrt(0.1))
}
