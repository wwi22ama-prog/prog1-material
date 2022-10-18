package main

import "fmt"

func ExampleGreet() {
	Greet()

	// Output:
	// Funktion 'greet()' aufgerufen.
}

func ExamplePrintSum() {
	PrintSum(3, 4)
	PrintSum(12, -2)

	// Output:
	// 3 + 4 == 7
	// 12 + -2 == 10
}

func ExampleComputeSum() {
	fmt.Println(ComputeSum(3, 4))
	sum1 := ComputeSum(15, 8)
	fmt.Println(sum1)
	sum2 := ComputeSum(25, 14)
	sum3 := ComputeSum(sum1, sum2)
	PrintSum(sum3, ComputeSum(13, sum2))

	// Output:
	// 7
	// 23
	// 62 + 52 == 114
}

func ExamplePrintTwice() {

	PrintTwice("Hallo")
	PrintTwice("Dies ist ein längerer String.")

	// Output:
	// HalloHallo
	// Dies ist ein längerer String.Dies ist ein längerer String.
}

func ExamplePrintSep() {
	PrintSep("Hallo", "Welt", " ")
	PrintSep("Käse", "kuchen", "")
	PrintSep("Zeile1", "Zeile2", "\n")

	// Output:
	// Hallo Welt
	// Käsekuchen
	// Zeile1
	// Zeile2
}

func ExampleConcatSep_simple() {
	s1 := ConcatSep("Hallo", "Welt", " ")
	fmt.Println(s1)
	fmt.Println(ConcatSep("Käse", "kuchen", ""))

	// Output:
	// Hallo Welt
	// Käsekuchen
}

func ExampleConcatSep_complex() {
	sep := "\n"
	fmt.Println(
		ConcatSep(
			ConcatSep(
				"Zeile3",
				"Zeile4",
				sep),
			ConcatSep(
				"Zeile5",
				ConcatSep(
					"Zeile6",
					"Zeile7",
					sep),
				sep),
			sep),
	)

	// Output:
	// Zeile3
	// Zeile4
	// Zeile5
	// Zeile6
	// Zeile7
}
