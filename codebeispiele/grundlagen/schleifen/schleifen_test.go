package main

import "fmt"

// Aufrufe der Funktion CountTo()
func ExampleCountTo() {
	CountTo(5)
	CountTo(17)

	// Output:
	// 0
	// 1
	// 2
	// 3
	// 4
	// 5
	// 0
	// 1
	// 2
	// 3
	// 4
	// 5
	// 6
	// 7
	// 8
	// 9
	// 10
	// 11
	// 12
	// 13
	// 14
	// 15
	// 16
	// 17
}

// Aufrufe der Funktion CountToStep()
func ExampleCountToStep() {
	CountToStep(5, 2)
	CountToStep(17, 2)
	CountToStep(17, 4)
	CountToStep(17, 6)

	// Output:
	// 0
	// 2
	// 4
	// 0
	// 2
	// 4
	// 6
	// 8
	// 10
	// 12
	// 14
	// 16
	// 0
	// 4
	// 8
	// 12
	// 16
	// 0
	// 6
	// 12
}

func ExampleCountDown() {
	CountDown(5)
	CountDown(17)

	// Output:
	// 5
	// 4
	// 3
	// 2
	// 1
	// 0
	// 17
	// 16
	// 15
	// 14
	// 13
	// 12
	// 11
	// 10
	// 9
	// 8
	// 7
	// 6
	// 5
	// 4
	// 3
	// 2
	// 1
	// 0
}

// Aufrufe der Funktion ConcatNTimes()
// Diese Funktion liefert wieder ein Ergebnis,
// das wir ausgeben oder verwenden können.
func ExampleConcatNTimes() {
	fmt.Println(ConcatNTimes("Hallo", 3))
	s1 := ConcatNTimes("Welt", 4)
	fmt.Println(s1)
	s2 := ConcatNTimes(ConcatNTimes("ABC ", 2), 4)
	fmt.Println(s2)

	// Output:
	// HalloHalloHallo
	// WeltWeltWeltWelt
	// ABC ABC ABC ABC ABC ABC ABC ABC
}

// Aufrufe der Funktion SumN().
func ExampleSumN() {
	fmt.Println(SumN(3))
	fmt.Println(SumN(5))
	fmt.Println(SumN(10))
	fmt.Println(SumN(500))

	// Output:
	// 6
	// 15
	// 55
	// 125250
}
