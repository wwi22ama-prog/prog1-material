package misc

import "fmt"

func ExampleFilterUpperCase() {
	fmt.Println(FilterUpperCase("AB%&ab123"))
	fmt.Println(FilterUpperCase("ABabCD"))
	fmt.Println(FilterUpperCase("abc"))
	fmt.Println(FilterUpperCase(""))
	// Output:
	// AB
	// ABCD
	//
	//
}

func ExampleFilterDigits() {
	fmt.Println(FilterDigits("AB%&ab123"))
	fmt.Println(FilterDigits("AB123cd456%ac"))
	fmt.Println(FilterDigits("ABab%"))
	fmt.Println(FilterDigits(""))
	// Output:
	// 123
	// 123456
	//
	//
}

func ExampleIsUpperCase() {
	fmt.Println(IsUpperCase('A'))
	fmt.Println(IsUpperCase('b'))
	fmt.Println(IsUpperCase('1'))
	fmt.Println(IsUpperCase('%'))

	// Output:
	// true
	// false
	// false
	// false
}

func ExampleIsDigit() {
	fmt.Println(IsDigit('A'))
	fmt.Println(IsDigit('b'))
	fmt.Println(IsDigit('1'))
	fmt.Println(IsDigit('%'))

	// Output:
	// false
	// false
	// true
	// false
}
