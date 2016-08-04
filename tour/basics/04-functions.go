package main

import "fmt"

func add(x int, y int) int {
	return x + y
}

// When two or more consecutive named function 
// parameters share a type, you can omit the type 
// from all but the last. 
func prod(x, y int) int {
	return x * y
}

// multiple results
func swap(x, y string) (string, string) {
	return y, x
}

// named return values
// => naked return
func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	return
}

func main() {
	fmt.Println(add(42, 13))
	fmt.Println(prod(3, 5))
	a, b := swap("hello", "world")
	fmt.Println(a, b)
	fmt.Println(split(13))
}
