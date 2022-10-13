package main

import "fmt"

func ExampleSum() {
	fmt.Println(Sum(15))
	fmt.Println(Sum(5))
	fmt.Println(Sum(1))
	fmt.Println(Sum(0))
	fmt.Println(Sum(-1))

	// Output:
	// 120
	// 15
	// 1
	// 0
	// 0
}
