package misc

import "fmt"

// Hilfsfunktion, liefert eine Testtabelle.
func testTable1() [][]string {
	return [][]string{
		{"A1", "B1", "C1", "D1"},
		{"A2", "B2", "C2", "D2"},
		{"A3", "B3", "C3", "D3"},
	}
}

func ExampleTranspose() {
	fmt.Println(Transpose(testTable1()))

	// Output:
	// [[A1 A2 A3] [B1 B2 B3] [C1 C2 C3] [D1 D2 D3]]
}

func ExampleGetColumn() {
	fmt.Println(GetColumn(testTable1(), 0))
	fmt.Println(GetColumn(testTable1(), 1))
	fmt.Println(GetColumn(testTable1(), 2))

	// Output:
	// [A1 A2 A3]
	// [B1 B2 B3]
	// [C1 C2 C3]
}

func ExampleFlattenTableString() {
	table1 := [][]string{{"AB", "CD"}, {"EF", "GH"}}

	fmt.Println(FlattenTableString(table1))

	// Output:
	// [AB CD EF GH]
}

func ExampleFlattenTableInt() {
	table1 := [][]int{{1, 2}, {3, 4}}

	fmt.Println(FlattenTableInt(table1))

	// Output:
	// [1 2 3 4]
}
