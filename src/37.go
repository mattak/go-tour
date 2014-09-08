package main

import "fmt"

type Vertex struct {
	Lat, Long float64
}

var m map[string]Vertex

func main() {
	m = make(map[string]Vertex)
	m["Aomono Yokocho"] = Vertex{
		35.609416, 139.742986,
	}
	fmt.Println(m["Aomono Yokocho"])
}
