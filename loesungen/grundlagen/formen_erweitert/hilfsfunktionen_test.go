package main

import "fmt"

func ExamplePrintCustomRow() {
	PrintCustomRow(3, "A", "B")
	PrintCustomRow(5, "A", "B")
	PrintCustomRow(5, "*", "+")

	// Output:
	// ABA
	// ABBBA
	// *+++*
}

func ExamplePrintCustomRectangle() {
	PrintCustomRectangle(3, 4, "A", "X")
	fmt.Println()
	PrintCustomRectangle(5, 2, "A", "X")
	fmt.Println()
	PrintCustomRectangle(2, 2, "A", "X")

	// Output:
	// AAA
	// AXA
	// AXA
	// AAA
	//
	// AAAAA
	// AAAAA
	//
	// AA
	// AA
}
