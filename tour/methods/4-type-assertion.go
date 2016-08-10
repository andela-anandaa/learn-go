package main

import "fmt"

func main() {
	var i interface{} = "hello"

	// A type assertion provides access to an 
	// interface value's underlying concrete value.

	s := i.(string)
	fmt.Println(s)

	s, ok := i.(string)
	fmt.Println(s, ok)

	f, ok := i.(float64)
	fmt.Println(f, ok)

	// If i does not hold a T, the statement will trigger a panic.
	// f = i.(float64)		// panic
	// fmt.Println(f)

	do(21)
	do("hello")
	do(true)

}

func do(i interface{}) {
	// type switch
	switch v := i.(type) {
	case int:
		fmt.Printf("Twice %v is %v\n", v, v * 2)
	case string:
		fmt.Printf("%q is %v bytes long\n", v, len(v))
	default:
		fmt.Printf("I don't know about type %T\n", v)
	}
}