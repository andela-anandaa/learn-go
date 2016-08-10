package main

import (
	"fmt"
	"math"
)

// An interface type is defined as a set of method signatures. 

type Abser interface {
	Abs() float64
}

type I interface {
	M()
}

func main() {
	var a Abser
	f := MyFloat(-math.Sqrt2)
	v := Vertex{3, 4}

	a = f 		// a MyFloat implements Abser
	a = &v 		// a *Vertex impletements Abser

	// In the following line, v is a Vertex (not *Vertex)
	// and does NOT implement Abser.
	// a = v

	fmt.Println(a.Abs())

	// Interfaces are implemented implicitly

	// Nil interface values
	var i I
	i.M() // run-time error
}


type MyFloat float64

func (f MyFloat) Abs() float64 {
	if f < 0 {
		return float64(-f)
	}
	return float64(f)
}

type Vertex struct {
	X, Y float64
}

func (v *Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}