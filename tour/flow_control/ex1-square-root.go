package main

import (
	"fmt"
	"math"
)

func Sqrt(x float64) float64 {
	// using Newton's method
	// z[n+1] = z[n] - {(z[n]^2 - x) / (2 * z[n])}
	// n + 1 = n - { (n^2 - x) / (2 * n)}
	diff := float64(0.5)
	z := float64(1)
	var rhs, lhs float64

	for {
		lhs = z + 1
		rhs = z - (((z * z) - x) / (2 * z))
		diff = math.Abs(lhs - rhs)

		if diff < 1 {
			return z + 1
		} else {
			z += 1
		}
	}
	return 0
}

func main() {
	fmt.Println(Sqrt(4225))
}