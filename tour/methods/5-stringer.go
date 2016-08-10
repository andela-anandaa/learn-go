package main

import "fmt"

type Person struct {
	Name string
	Age int
}

// A Stringer is a type that can describe itself as a string. 
// The fmt package (and many others) look for this interface to print values.
/*
	type Stringer interface {
	    String() string
	}
*/

func (p Person) String() string {
	return fmt.Sprintf("<%v (%v years)>", p.Name, p.Age)
}

func main() {
	a := Person{"James Brown", 42}
	z := Person{"Joe Black", 20}
	fmt.Println(a, z)
}