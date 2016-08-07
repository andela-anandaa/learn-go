package main

import (
	"fmt"
	"strings"
)

func main() {
	var a [2]string
	a[0] = "Hello"
	a[1] = "World"
	fmt.Println(a[0], a[1])
	fmt.Println(a)

	primes := [6]int{2, 3, 5, 7, 11, 13}
	fmt.Println(primes)

	// slices
	var s []int = primes[0:3]
	fmt.Println(s)

	// A slice is a dynamically-sized, flexible view 
	// into the elements of an array

	// A slice does not store any data, it just describes a section of 
	// an underlying array.
	// Changing the elements of a slice modifies the corresponding 
	// elements of its underlying array. 

	names := [4]string{
		"John",
		"James",
		"Mark",
		"Paul",
	}

	fmt.Println(names)

	x := names[0:2]
	y := names[1:3]
	fmt.Println(x, y)

	y[0] = "Matthew"
	fmt.Println(x, y)
	fmt.Println(names)

	// Slice literals
	q := []int{2, 5, 10, 12, 20}
	fmt.Println(q)

	r := []bool{true, false, true, true, false, true}
	fmt.Println(r)

	sk := []struct{
		i int
		b bool
	}{
		{2, true},
		{3, false},
		{5, true},
		{7, true},
		{11, false},
		{13, true},
	}

	fmt.Println(sk)

	// slice length and capacity
	fmt.Println(len(sk), cap(sk))

	// nil slice
	var ns []int
	if ns == nil {
		fmt.Println("nil array:", ns)
	}

	// creating a slice with make
	// Slices can be created with the built-in make function; 
	// this is how you create dynamically-sized arrays. 

	ma := make([]int, 5)
	printSlice("ma", ma)

	mb := make([]int, 0, 5)
	printSlice("mb", mb)

	mc := mb[:2]
	printSlice("mc", mc)

	md := mc[2:5]
	printSlice("md", md)

	// test Slices of Slices
	tickTack()
}

func printSlice(s string, x []int) {
	fmt.Printf("%s len=%d cap=%d %v\n", s, len(x), cap(x), x)
}

func tickTack() {
	// Slices of Slices
	// create a tick-tac-toe board.
	board := [][]string{
		[]string{"_", "_", "_"},
		[]string{"_", "_", "_"},
		[]string{"_", "_", "_"},
	}

	// The players take turns
	board[0][0] = "X"
	board[2][2] = "O"
	board[1][2] = "X"
	board[1][0] = "O"
	board[0][2] = "X"

	for i := 0; i < len(board); i++ {
		fmt.Printf("%s\n", strings.Join(board[i], " "))
	}

	appendingToSlice()

	// range

	// When ranging over a slice, two values are returned for each iteration. 
	// The first is the index, and the second is a copy of the element at that index. 
	var pow = []int{1, 2, 4, 8, 16, 32, 64, 128}

	for i, v := range pow {
		fmt.Printf("2**%d = %d\n", i, v)
	}

	// You can skip the index or value by assigning to _. 
	for _, value := range pow {
		fmt.Printf("%d\n", value)
	}
}

func appendingToSlice() {
	var s []int
	printSlice("s:", s)

	// append works on nil slices.
	s = append(s, 0)
	printSlice("s:", s)

	// The slice grows as needed
	s = append(s, 1)
	printSlice("s:", s)

	// we can add more than one element at a time
	s = append(s, 5, 6, 8, 10, 12)
	printSlice("s:", s)
}