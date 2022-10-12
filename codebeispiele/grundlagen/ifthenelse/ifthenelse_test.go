package main

import "fmt"

// Tests der Funktion Judge()
func ExampleJudge() {
	Judge(42)
	Judge(13)

	// Output:
	// 42 ist eine gute Zahl!
	// 13 ist keine gute Zahl!
}

// Tests der Funktion IsLarge()
// Diese Funktion liefert entweder true oder false.
// Diesen Datentyp nennt man bool.
func ExampleIsLarge() {
	fmt.Println(IsLarge(42))
	fmt.Println(IsLarge(13))

	// Output:
	// Juhuuuu!
	// true
	// false

}

// Tests der Funktion PrintDivisors()
func ExamplePrintDivisors() {
	PrintDivisors(3, 25)
	PrintDivisors(17, 100)
	PrintDivisors(2, 10)

	// Output:
	// 0
	// 3
	// 6
	// 9
	// 12
	// 15
	// 18
	// 21
	// 24
	// 0
	// 17
	// 34
	// 51
	// 68
	// 85
	// 0
	// 2
	// 4
	// 6
	// 8
	// 10
}

// Tests der Funktion SumNRec()
func ExampleSumNRec() {
	fmt.Println(SumNRec(5))
	fmt.Println(SumNRec(13))
	fmt.Println(SumNRec(200))

	// Output:
	// 15
	// 91
	// 20100
}
