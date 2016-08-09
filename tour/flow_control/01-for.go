package main

import "fmt"

func main() {
	sum := 0
	for i := 0; i < 10; i++ {
		sum += i
	}
	fmt.Println(sum)

	// init and post statements are optional
	sum_ := 1
	for ; sum_ < 5; {
		sum_ += sum_
	}
	fmt.Println(sum_)

	// Go's while
	i := 1
	for i < 5 {
		i += i
	}
	fmt.Println(i)

	// Forever
	count := 1
	for {
		fmt.Println(count)
		count ++

		if count > 10 {
			break
		}
	}
}