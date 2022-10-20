package main

func ExamplePrintChessBoard() {
	PrintChessBoard(4, 3)

	// Output:
	// ****++++****
	// ****++++****
	// ****++++****
	// ****++++****
	// ++++****++++
	// ++++****++++
	// ++++****++++
	// ++++****++++
	// ****++++****
	// ****++++****
	// ****++++****
	// ****++++****
}

func ExamplePrintChessBoardRow() {
	PrintChessBoardRow(4, 3, "*", " ")

	// Output:
	// ****    ****
	// ****    ****
	// ****    ****
	// ****    ****
}
