package main

import (
	"fmt"
	"math"
)

const EPSILON = 1e-9

type ErrNegativeSqrt float64

func (f ErrNegativeSqrt) Error() string {
	return fmt.Sprintf("cannot Sqrt nagative number: %f", f)
}

func Sqrt(f float64) (float64, error) {
	if f < 0 {
		return 0, ErrNegativeSqrt(f)
	}

	z := f
	za := z + EPSILON

	for math.Abs(z-za) > EPSILON {
		za = z
		z = z - (z*z-f)/(2*z)
	}

	return f, nil
}

func main() {
	fmt.Println(Sqrt(2))
	fmt.Println(Sqrt(-2))
}
