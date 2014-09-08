package main

import "fmt"

type Vertex struct {
	Lat, Long float64
}

func main() {
	var m = map[string]Vertex{
		"Akihabara": {
			35.697860, 139.775553,
		},
		"UDX": {
			35.700545, 139.772831,
		},
	}
	fmt.Println(m["Akihabara"])
	fmt.Println(m["UDX"])
}
