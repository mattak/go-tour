package main

import "fmt"

func adder(basefactor int) func(int) int {
	sum := 0
	return func(x int) int {
		sum += basefactor * x
		return sum
	}
}

func main() {
	pos, neg := adder(+1), adder(-1)
	for i := 0; i < 10; i++ {
		fmt.Println(pos(i), neg(i))
	}
}
