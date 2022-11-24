package addressfunctions

import "fmt"

func ExampleRowNumber() {
	fmt.Println(RowNumber("A4"))
	fmt.Println(RowNumber("B2"))
	fmt.Println(RowNumber("C1"))
	fmt.Println(RowNumber("AA16"))
	fmt.Println(RowNumber("AB14"))
	fmt.Println(RowNumber("BB25"))
	fmt.Println(RowNumber("ZZ77"))
	fmt.Println(RowNumber("ACJ123"))

	// Output:
	// 4
	// 2
	// 1
	// 16
	// 14
	// 25
	// 77
	// 123
}

func ExampleColumnNumber() {
	fmt.Println(ColumnNumber("A4"))
	fmt.Println(ColumnNumber("B2"))
	fmt.Println(ColumnNumber("C1"))
	fmt.Println(ColumnNumber("AA16"))
	fmt.Println(ColumnNumber("AB14"))
	fmt.Println(ColumnNumber("BB25"))
	fmt.Println(ColumnNumber("ZZ77"))
	fmt.Println(ColumnNumber("ACJ123"))

	// Output:
	// 1
	// 2
	// 3
	// 27
	// 28
	// 54
	// 702
	// 764
}

func ExampleFirstRow() {
	fmt.Println(FirstRow("A4:C15"))
	fmt.Println(FirstRow("A4:A15"))
	fmt.Println(FirstRow("B5:C17"))

	// Output:
	// 4
	// 4
	// 5
}

func ExampleFirstColumn() {
	fmt.Println(FirstColumn("A4:C15"))
	fmt.Println(FirstColumn("A4:A15"))
	fmt.Println(FirstColumn("B5:C17"))

	// Output:
	// 1
	// 1
	// 2
}

func ExampleLastRow() {
	fmt.Println(LastRow("A4:C15"))
	fmt.Println(LastRow("A4:A15"))
	fmt.Println(LastRow("B5:C17"))

	// Output:
	// 15
	// 15
	// 17
}

func ExampleLastColumn() {
	fmt.Println(LastColumn("A4:C15"))
	fmt.Println(LastColumn("A4:A15"))
	fmt.Println(LastColumn("B5:C17"))

	// Output:
	// 3
	// 1
	// 3
}

func ExampleRowCount() {
	fmt.Println(RowCount("A4:C15"))
	fmt.Println(RowCount("A4:A15"))
	fmt.Println(RowCount("B5:C17"))

	// Output:
	// 12
	// 12
	// 13
}

func ExampleColumnCount() {
	fmt.Println(ColumnCount("A4:C15"))
	fmt.Println(ColumnCount("A4:A15"))
	fmt.Println(ColumnCount("B5:C17"))

	// Output:
	// 3
	// 1
	// 2
}

func ExampleIsRange() {
	fmt.Println(IsRange("A4:C15"))
	fmt.Println(IsRange("AZ4:CC15"))
	fmt.Println(IsRange("A4"))
	fmt.Println(IsRange("B15"))
	fmt.Println(IsRange("AZ4::CC15"))

	// Output:
	// true
	// true
	// false
	// false
	// false
}

func ExampleIsCell() {
	fmt.Println(IsCell("A4:C15"))
	fmt.Println(IsCell("AZ4:CC15"))
	fmt.Println(IsCell("A4"))
	fmt.Println(IsCell("B15"))
	fmt.Println(IsCell("ae"))
	fmt.Println(IsCell("34A"))

	// Output:
	// false
	// false
	// true
	// true
	// false
	// false
}
