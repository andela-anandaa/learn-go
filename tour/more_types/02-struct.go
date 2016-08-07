package main

import "fmt"

type Vertex struct {
	X int
	Y int
}

type Person struct {
	name string
	age int
	location string
}

// struct literal
type Coord struct {
	lon, lat int
}

var (
	c1 = Coord{1, 2}			// has type Ord
	c2 = Coord{lon: 1}			// lat: 0 is implicit
	c3 = Coord{}				// lon: 0 and lat: 0
	pc = &Coord{1, 2}			// has type *Coord
)

func main() {
	v := Vertex{1, 2}
	fmt.Println(v)

	var p Person
	p.name = "James Nandaa"
	p.age = 22
	p.location = "CA"

	fmt.Println(p)

	// pointers to struct
	p = Person{"Jane", 24, "KE"}

	pp := &p
	pp.age = 40 			// without explicit dereferencing
	(*pp).location = "UG"
	fmt.Println(p)

	// for struct literals
	fmt.Println(c1, c2, c3, pc)
}