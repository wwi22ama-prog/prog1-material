package main

import (
	"fmt"
	"sort"
)

func main() {

}

func Anfangsbeispiele() {
	liste1 := make([]int, 0)
	fmt.Println(liste1)

	liste1 = append(liste1, 42)
	fmt.Println(liste1)

	liste1 = append(liste1, 25)
	liste1 = append(liste1, 57)
	liste1 = append(liste1, 23)
	fmt.Println(liste1)

	fmt.Println(liste1[3])

	liste1[2] = 999
	fmt.Println(liste1)

	// liste1[4] = 9999

	liste2 := []int{15, 25, 35, 45}
	fmt.Println(liste2)

	liste1 = append(liste1, 77, 55, 203)
	liste1 = append(liste1, liste2...)
	fmt.Println(liste1)

	sort.Ints(liste1)
	fmt.Println(liste1)

	liste3 := liste1[3:6]
	fmt.Println(liste3)

	liste3[2] = 999999999
	fmt.Println(liste3)
	fmt.Println(liste1)

	for i := 0; i < len(liste1); i++ {
		fmt.Println(liste1[i])
	}
	for i := 0; i < len(liste1); i++ {
		v := liste1[i]
		fmt.Println(v)
	}
	for i, v := range liste1 {
		fmt.Println(i, v)
	}
	for i := range liste1 {
		fmt.Println(i)
	}
	for _, v := range liste1 {
		fmt.Println(v)
	}

}
