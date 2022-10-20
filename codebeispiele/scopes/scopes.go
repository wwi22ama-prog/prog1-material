package main

import "fmt"

func Bar() {
	x := 42
	fmt.Println(x)
}

func Foo(x int) {
	fmt.Println(x)
	Bar()
	fmt.Println(x)
}

func main() {
	// Foo(23)
	// Confusion()
	CountDown(5)
}

func Confusion() {
	x := 11
	{
		fmt.Println(x)
		y := 22
		{
			x := 33
			fmt.Println(x)
		}
		fmt.Println(y)
	}
	//	fmt.Println(y) // y ist hier nicht mehr gültig.
	fmt.Println(x)
}

func CountDown(x int) {
	if x == 0 {
		return
	}
	fmt.Println(x)
	CountDown(x - 1)
	fmt.Println(x)
}
