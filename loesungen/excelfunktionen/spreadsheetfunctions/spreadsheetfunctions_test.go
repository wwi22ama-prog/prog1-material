package spreadsheetfunctions

import "fmt"

// Hilfsfunktion: Liefert eine Test-Tabelle.
func testTable1() [][]string {
	return [][]string{
		{"1", "1", "DHBW"},
		{"2", "3", "Mannheim"},
		{"A", "2", "oder"},
		{"B", "7", "irgendwo"},
		{"C", "5", "anders"},
	}
}

func ExampleComputeSum() {
	fmt.Println(ComputeSum(testTable1(), "B1:B5"))

	// Output:
	// 18
}

func ExampleComputeMax() {
	fmt.Println(ComputeMax(testTable1(), "B1:B5"))

	// Output:
	// 7
}

func ExampleComputeLookup() {
	fmt.Println(ComputeLookup(testTable1(), "A1", "B1:B5", "C1:C5"))
	fmt.Println(ComputeLookup(testTable1(), "A2", "B1:B5", "C1:C5"))

	// Output:
	// DHBW
	// oder
}

func ExampleComputeVLookup() {
	fmt.Println(ComputeVLookup(testTable1(), "A1", "B1:C5", 2))
	fmt.Println(ComputeVLookup(testTable1(), "A2", "B1:C5", 2))

	// Output:
	// DHBW
	// oder
}

func ExampleComputeFunction() {
	fmt.Println(ComputeFunction(testTable1(), "SUMME(B1:B5)"))
	fmt.Println(ComputeFunction(testTable1(), "MAX(B1:B5)"))
	fmt.Println(ComputeFunction(testTable1(), "LOOKUP(A1;B1:B5;C1:C5)"))
	fmt.Println(ComputeFunction(testTable1(), "LOOKUP(A2;B1:B5;C1:C5)"))
	fmt.Println(ComputeFunction(testTable1(), "VLOOKUP(A1;B1:C5;2)"))
	fmt.Println(ComputeFunction(testTable1(), "VLOOKUP(A2;B1:C5;2)"))

	// Output:
	// 18
	// 7
	// DHBW
	// oder
	// DHBW
	// oder
}
