package main

import (
	"fmt"
	"math"
)

type ErrNegativeSqrt struct {
	err string
	arg float64
}

func (e *ErrNegativeSqrt) Error() string {
	return fmt.Sprintf("%v: %v\n", e.err, e.arg)
}

func Sqrt(x float64) (float64, error) {
	// using Newton's method
	// z[n+1] = z[n] - {(z[n]^2 - x) / (2 * z[n])}
	// n + 1 = n - { (n^2 - x) / (2 * n)}

	if x < 0 {
		return 0, &ErrNegativeSqrt{"cannot Sqrt negative number", x}
	}

	diff := float64(0.5)
	z := float64(1)
	var rhs, lhs float64

	for {
		lhs = z + 1
		rhs = z - (((z * z) - x) / (2 * z))
		diff = math.Abs(lhs - rhs)

		if diff < 1 {
			return z + 1, nil
		} else {
			z += 1
		}
	}

	return 0, nil
}

func main() {
	x, err := Sqrt(-4225)
	fmt.Println(x, err)
}