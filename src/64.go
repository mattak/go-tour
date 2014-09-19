package main

import "fmt"

func sum(a []int, c chan int) {
	sum := 0
	for _, v := range a {
		sum += v
	}
	c <- sum
}

func main() {
	a := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	c := make(chan int)

	go sum(a[5:], c)
	go sum(a[:5], c)

	x := <-c
	y := <-c

	fmt.Println(x)
	fmt.Println(y)
}
