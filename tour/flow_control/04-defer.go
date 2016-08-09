package main

import "fmt"

func stacked_defers() {
	fmt.Println("counting")

	for i := 0; i < 10; i++ {
		defer fmt.Println(i)
	}

	fmt.Println("done")
}

func main() {
	max_num := 0

	defer fmt.Println("Number: ", max_num)
	defer fmt.Println("world")

	count := 0

	for ; count < 10; {
		max_num = count * count
		count += 1
	}

	stacked_defers()

	fmt.Println("Number: ", max_num)

	fmt.Println("hello")
}