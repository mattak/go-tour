package main

import "fmt"
import cmplx "math/cmplx"

const EPSILON = 1e-15

func Cbrt(x complex128) complex128 {
	z := x
	za := z + EPSILON*2

	for cmplx.Abs(za-z) > EPSILON {
		za = z
		z = z - (z*z*z-x)/(3*z*z)
	}

	return z
}

func main() {
	src := complex128(2)
	dst := Cbrt(src)
	fmt.Println("answer:", dst)
	fmt.Println("correct?:", cmplx.Abs(src-cmplx.Pow(dst, 3)) < EPSILON)
}
