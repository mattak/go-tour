package main

import "fmt"

func add(x int, y int) int {
	return x + y
}

func addS(x string, y string) string {
	return x + y
}

func main() {
	fmt.Println("add int:", add(1, 2))
	fmt.Println("add String:", addS("hello", " world"))
}
