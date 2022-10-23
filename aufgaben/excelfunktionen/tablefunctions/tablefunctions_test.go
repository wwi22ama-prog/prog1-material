package tablefunctions

import "fmt"

func ExampleParseFunctionCall() {
	fmt.Println(ParseFunctionCall("SUMME(A1:B3"))
	fmt.Println(ParseFunctionCall("SVERWEIS(D5;A1:C5;3)"))

	// Output:
	// [SUMME A1:B3]
	// [SVERWEIS D5 A1:C5 3]
}

func ExampleGetFunctionName() {
	fmt.Println(GetFunctionName("SUMME(A1:B3"))
	fmt.Println(GetFunctionName("SVERWEIS(D5;A1:C5;3)"))

	// Output:
	// SUMME
	// SVERWEIS
}

func ExampleGetArgument() {
	fmt.Println(GetArgument("SUMME(A1:B3", 0))
	fmt.Println(GetArgument("SVERWEIS(D5;A1:C5;3)", 0))
	fmt.Println(GetArgument("SVERWEIS(D5;A1:C5;3)", 1))
	fmt.Println(GetArgument("SVERWEIS(D5;A1:C5;3)", 2))

	// Output:
	// A1:B3
	// D5
	// A1:C5
	// 3
}
