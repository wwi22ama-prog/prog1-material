package listen

import "fmt"

func IntSlice1() {
	var intSlice1 []int = []int{2, 4, 6, 8}

	for i := 0; i < len(intSlice1); i++ {
		fmt.Println(intSlice1[i])
	}
	for _, v := range intSlice1 {
		fmt.Println(v)
	}
}

func IntSlice2() {
	intSlice2 := []int{1, 3, 5, 7}

	for _, v := range intSlice2 {
		fmt.Println(v)
	}
}

func IntSlice3() {
	intSlice2 := []int{1, 3, 5, 7}
	intSlice3 := make([]int, 0)

	intSlice3 = append(intSlice3, 42)
	intSlice3 = append(intSlice3, intSlice2...)
	intSlice3 = append(intSlice3, 38)

	fmt.Println(intSlice3)
}
