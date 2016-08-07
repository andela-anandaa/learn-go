package main

import "fmt"

type Vertex struct {
	Lat, Long float64
}

var m map[string]Vertex

func main() {
	m = make(map[string]Vertex)
	m["Beyonde Labs"] = Vertex{
		40.684, 23.399,
	}
	fmt.Println(m["Beyonde Labs"])

	// map literal
	var n = map[string]Vertex{
		"Bell Labs": Vertex{
			40.68, -74.39,
		},
		"Google": Vertex{
			37.42, -122.08,
		},
	}

	n["Beyonde Labs"] = Vertex{40.68, 23.39}

	fmt.Println(n)

	delete(n, "Google")
	fmt.Println(n)

	elem, ok := n["Microsoft"]
	fmt.Println("The value:", elem, "Present?", ok)
}