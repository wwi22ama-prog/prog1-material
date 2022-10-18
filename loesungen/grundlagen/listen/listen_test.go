package main

import "fmt"

func ExampleLargest() {
	l1 := []int{2, 1, 17, 25, 3, 12, 42, 5}
	l2 := []int{2}
	l3 := []int{1, 2, 3, 4}
	l4 := []int{5, 4, 3, 2}

	fmt.Println(Largest(l1))
	fmt.Println(Largest(l2))
	fmt.Println(Largest(l3))
	fmt.Println(Largest(l4))

	// Output:
	// 6 42
	// 0 2
	// 3 4
	// 0 5
}

func ExampleReverse() {
	l1 := []int{2, 1, 17, 25, 3, 12, 42, 5}
	l2 := []int{2}
	l3 := []int{1, 2, 3, 4}
	l4 := []int{5, 4, 3, 2}

	fmt.Println(Reverse(l1))
	fmt.Println(Reverse(l2))
	fmt.Println(Reverse(l3))
	fmt.Println(Reverse(l4))

	// Output:
	// [5 42 12 3 25 17 1 2]
	// [2]
	// [4 3 2 1]
	// [2 3 4 5]
}

func ExampleSum() {
	l1 := []int{2, 1, 17, 25, 3, 12, 42, 5}
	l2 := []int{2}
	l3 := []int{1, 2, 3, 4}
	l4 := []int{5, 4, 3, 2, 1}

	fmt.Println(Sum(l1))
	fmt.Println(Sum(l2))
	fmt.Println(Sum(l3))
	fmt.Println(Sum(l4))

	// Output:
	// 107
	// 2
	// 10
	// 15
}

func ExampleAverage() {
	l1 := []int{2, 1, 17, 25, 3, 12, 42, 5}
	l2 := []int{2}
	l3 := []int{1, 2, 3, 4}
	l4 := []int{5, 4, 3, 2, 1}

	fmt.Println(Average(l1))
	fmt.Println(Average(l2))
	fmt.Println(Average(l3))
	fmt.Println(Average(l4))

	// Output:
	// 13.375
	// 2
	// 2.5
	// 3
}

func ExampleMedian() {
	l1 := []int{2, 1, 17, 25, 3, 12, 42, 5}
	l2 := []int{2}
	l3 := []int{1, 2, 3, 4}
	l4 := []int{5, 4, 3, 2, 1}

	fmt.Println(Median(l1))
	fmt.Println(Median(l2))
	fmt.Println(Median(l3))
	fmt.Println(Median(l4))

	// Output:
	// 5
	// 2
	// 2
	// 3
}

func ExampleCountLarger() {
	l1 := []int{2, 1, 17, 25, 3, 12, 42, 5}
	l2 := []int{2}
	l3 := []int{1, 2, 3, 4}

	fmt.Println(CountLarger(l1, 5))
	fmt.Println(CountLarger(l2, 0))
	fmt.Println(CountLarger(l2, 3))
	fmt.Println(CountLarger(l3, 3))

	// Output:
	// 4
	// 1
	// 0
	// 1
}

func ExampleRemove() {
	l1 := []int{2, 1, 17, 25, 3, 12, 42, 5}
	l2 := []int{2}

	fmt.Println(Remove(l1, 25))
	fmt.Println(Remove(l1, 3))
	fmt.Println(Remove(l1, 6))
	fmt.Println(Remove(l2, 2))

	// Output:
	// [2 1 17 3 12 42 5]
	// [2 1 17 25 12 42 5]
	// [2 1 17 25 3 12 42 5]
	// []
}

func ExampleDigits() {
	fmt.Println(Digits(42))
	fmt.Println(Digits(3))
	fmt.Println(Digits(0))
	fmt.Println(Digits(150))

	// Output:
	// [4 2]
	// [3]
	// [0]
	// [1 5 0]
}

func ExampleElementProduct() {
	l1 := []int{2, 1, 17, 25, 3, 12, 42, 5}
	l2 := []int{21, 3, 2, 15, 4, 2, 25, 12}

	fmt.Println(ElementProduct(l1, l2))
	fmt.Println(ElementProduct(l1, l1))
	fmt.Println(ElementProduct(l2, l2))

	// Output:
	// [42 3 34 375 12 24 1050 60]
	// [4 1 289 625 9 144 1764 25]
	// [441 9 4 225 16 4 625 144]
}

func ExampleProduct() {
	l1 := []int{1, 2, 3, 4, 5}
	l2 := []int{3, 6, 9}
	l3 := []int{}

	fmt.Println(Product(l1))
	fmt.Println(Product(l2))
	fmt.Println(Product(l3))

	// Output:
	// 120
	// 162
	// 1
}

func ExampleIsbn13Checksum() {
	l1 := []int{9, 7, 8, 3, 4, 5, 5, 5, 0, 2, 3, 6}
	fmt.Println(Isbn13Checksum(l1))

	// Output:
	// 7
}

func ExampleDigitSum() {
	fmt.Println(DigitSum(42))
	fmt.Println(DigitSum(8))
	fmt.Println(DigitSum(80))
	fmt.Println(DigitSum(105))
	fmt.Println(DigitSum(0))

	// Output:
	// 6
	// 8
	// 8
	// 6
	// 0
}
