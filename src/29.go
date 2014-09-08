package main

import "fmt"

type Vertex struct {
	X, Y int
}

func main() {
	var a Vertex = Vertex{0, 1}
	v := new(Vertex)
	fmt.Println(v)
	v.X, v.Y = 11, 9
	fmt.Println(v)
	a.X = 1
	fmt.Println(a)
}
