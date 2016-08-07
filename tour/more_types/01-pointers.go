package main

import "fmt"

func main() {
	i, j := 42, 280

	p := &i 			// point to i
	// var p *int = &i
	fmt.Println(*p)		// read i through the pointer
	*p = 21				// set i through the pointer
	fmt.Println(i)		// see the new value of i

	p = &j				// point to j
	*p = *p / 7			// divide j through the pointer
	fmt.Println(j)		// see the new value of j


	// Unlike C, Go has no pointer arithmetic.
}