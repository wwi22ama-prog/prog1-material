package tableaccess

import "fmt"

// Hilfsfunktion für die Tests: Liefert eine Tabelle mit Beispieldaten.
func testTable1() [][]string {
	return [][]string{
		{"AB", "123", "DHBW"},
		{"CD", "456", "Mannheim"},
	}
}

// Hilfsfunktion für die Tests: Liefert eine Tabelle mit Beispieldaten.
// Größere Tabelle für Bereichstests.
func testTable2() [][]string {
	return [][]string{
		{"A1", "B1", "C1", "D1", "E1", "F1"},
		{"A2", "B2", "C2", "D2", "E2", "F2"},
		{"A3", "B3", "C3", "D3", "E3", "F3"},
		{"A4", "B4", "C4", "D4", "E4", "F4"},
	}
}

func ExampleGetCellValueString() {
	fmt.Println(GetCellValueString(testTable1(), "A1"))
	fmt.Println(GetCellValueString(testTable1(), "B1"))
	fmt.Println(GetCellValueString(testTable1(), "C1"))
	fmt.Println(GetCellValueString(testTable1(), "A2"))
	fmt.Println(GetCellValueString(testTable1(), "B2"))
	fmt.Println(GetCellValueString(testTable1(), "C2"))

	// Output:
	// AB
	// 123
	// DHBW
	// CD
	// 456
	// Mannheim
}

func ExampleGetCellValueInt() {
	fmt.Println(GetCellValueInt(testTable1(), "B1"))
	fmt.Println(GetCellValueInt(testTable1(), "B2"))

	// Output:
	// 123
	// 456
}

func ExampleGetRange() {
	fmt.Println(GetRange(testTable2(), "A1:A4"))
	fmt.Println(GetRange(testTable2(), "A2:A4"))
	fmt.Println(GetRange(testTable2(), "B1:C3"))
	fmt.Println(GetRange(testTable2(), "B1:B3"))
	fmt.Println(GetRange(testTable2(), "C1:C3"))

	// Output:
	// [[A1] [A2] [A3] [A4]]
	// [[A2] [A3] [A4]]
	// [[B1 C1] [B2 C2] [B3 C3]]
	// [[B1] [B2] [B3]]
	// [[C1] [C2] [C3]]
}

func ExampleGetRangeInt() {
	fmt.Println(GetRangeInt(testTable1(), "B1:B2"))
	fmt.Println(GetRangeInt(testTable1(), "A1:B2"))
	fmt.Println(GetRangeInt(testTable2(), "A1:A4"))
	fmt.Println(GetRangeInt(testTable2(), "A2:A4"))
	fmt.Println(GetRangeInt(testTable2(), "B1:C3"))

	// Output:
	// [[123] [456]]
	// [[0 123] [0 456]]
	// [[0] [0] [0] [0]]
	// [[0] [0] [0]]
	// [[0 0] [0 0] [0 0]]
}
