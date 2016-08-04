package main

import (
	"fmt"
	"math"
)

var i, j int = 1, 2

// svd not available outside func body
// k := 30.50

func main() {
	var c, python, java = true, false, "no!"
	// short variable declaration
	k := 30.50
	fmt.Println(i, j, k, c, python, java)

	// zero
	var i_ int
	var f float64
	var b bool
	var s string
	fmt.Printf("%v %v %v %q\n", i_, f, b, s)

	// type conversions
	var x, y int = 3, 65
	var f_ float64 = math.Sqrt(float64(x*x + y*y))
	var z uint = uint(f_)
	fmt.Println(x, y, z)

	// type inference
	v := 42.6
	cx := 0.343 + 5i
	fmt.Printf("v is of type %T\n", v)
	fmt.Printf("cx is of type %T\n", cx)

	// constants
	const pi = 3.14
	
}