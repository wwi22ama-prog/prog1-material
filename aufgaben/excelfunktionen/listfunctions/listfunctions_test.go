package listfunctions

import "fmt"

func ExampleFindString() {

	l1 := []string{"hallo", "Welt", "und", "mehr", "Wörter"}
	l2 := []string{}

	fmt.Println(FindString("hallo", l1))
	fmt.Println(FindString("Welt", l1))
	fmt.Println(FindString("und", l1))
	fmt.Println(FindString("mehr", l1))
	fmt.Println(FindString("Wörter", l1))
	fmt.Println(FindString("Nichts", l1))
	fmt.Println(FindString("", l1))
	fmt.Println(FindString("Irgendwas", l2))
	fmt.Println(FindString("", l2))

	// Output:
	// 0
	// 1
	// 2
	// 3
	// 4
	// -1
	// -1
	// -1
	// -1
}

func ExampleLookupString() {
	keys := []string{"hallo", "Welt", "und", "mehr", "Wörter"}
	values := []string{"hello", "world", "and", "more", "words"}

	fmt.Println(LookupString("hallo", keys, values))
	fmt.Println(LookupString("Welt", keys, values))
	fmt.Println(LookupString("und", keys, values))
	fmt.Println(LookupString("mehr", keys, values))
	fmt.Println(LookupString("Wörter", keys, values))
	fmt.Println(LookupString("", keys, values))
	fmt.Println(LookupString("Nichts", keys, values))

	// Output:
	// hello
	// world
	// and
	// more
	// words
	//
	//
}

func ExampleSumInt() {
	l1 := []int{1, 2, 3, 4, 5}
	l2 := []int{1, 1, 2, 3, 5, 8}
	l3 := []int{}

	fmt.Println(SumInt(l1))
	fmt.Println(SumInt(l2))
	fmt.Println(SumInt(l3))

	// Output:
	// 15
	// 20
	// 0
}

func ExampleMaxInt() {
	l1 := []int{1, 5, 2, 3, 8, 4}
	l2 := []int{5, 1, -1}
	l3 := []int{-1, -2, -3}
	l4 := []int{}

	fmt.Println(MaxInt(l1))
	fmt.Println(MaxInt(l2))
	fmt.Println(MaxInt(l3))
	fmt.Println(MaxInt(l4))

	// Output:
	// 8
	// 5
	// -1
	// 0
}

func ExampleSumIntTable() {
	l1 := [][]int{{1, 3, 2}, {4, 6, 5}}

	fmt.Println(SumIntTable(l1))

	// Output:
	// 21
}

func ExampleMaxIntTable() {
	l1 := [][]int{{1, 3, 2}, {4, 6, 5}}

	fmt.Println(MaxIntTable(l1))

	// Output:
	// 6
}
