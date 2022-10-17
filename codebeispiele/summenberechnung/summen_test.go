package main

import "fmt"

func ExampleSumme1Bis5() {
	Summe1Bis5()

	// Output:
	// Summe: 15
}

func ExampleSumme1BisN() {
	fmt.Println(Summe1BisN(5))
	fmt.Println(Summe1BisN(10))
	fmt.Println(Summe1BisN(15))

	// Output:
	// 15
	// 55
	// 120
}
