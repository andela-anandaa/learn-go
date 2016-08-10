package main

import (
	"fmt"
	"math"
)

type Vertex struct {
	X, Y float64
}

func (v Vertex) Abs() float64 {
	return math.Sqrt(v.X * v.X + v.Y * v.Y)
}

// a method is just a function with a receiver argument.
func Abs(v Vertex) float64 {
	return math.Sqrt(v.X * v.X + v.Y * v.Y)
}

// You can declare a method on non-struct types, too.
type MyFloat float64

func (f MyFloat) Abs() float64 {
	if f < 0 {
		return float64(-f)
	}
	return float64(f)
}

// pointer receivers

func (v *Vertex) Scale(f float64) {
	v.X *= f
	v.Y *= f
}

func Scale(v *Vertex, f float64) {
	v.X *= f
	v.Y *= f
}


func main() {
	v := Vertex{3, 4}
	fmt.Println(v.Abs())

	fmt.Println(Abs(v))

	f := MyFloat(-math.Sqrt2)
	fmt.Println(f.Abs())

	v.Scale(10)
	fmt.Println(v.Abs())
	Scale(&v, 5)
	fmt.Println(v.Abs())
}