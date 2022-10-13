package main

import (
	"fmt"
)

func main() {

	/*
		var result int
		result = 0
	*/
	/*
	  result = result + 1
	  result = result + 2
	  result = result + 3
	  result = result + 4
	  result = result + 5
	*/

	/*
	  var n int
	  n = 5
	*/

	//for i:= 1; i<=n; i = i+1 {
	//for i:= 1; i<=n; i += 1 {
	// for i:= 1; i<=n; i+1 { wäre falsch

	fmt.Printf("1 + 2 + 3 + 4 + 5 == %v\n", Sum(5))
	fmt.Printf("Summe von 1 bis 150 == %v\n", Sum(150))

	var input int

	fmt.Print("Bitte eine Zahl eingeben: ")
	fmt.Scanln(&input)
	fmt.Println("Summe:", Sum(input))

	s := "Hallo Welt"
	// for i := 0; i< len(s); i++ {
	// for i := range s {
	for _, c := range s {
		fmt.Printf("%c\n", c)
	}

}

// Funktion, die einen Parameter n vom Typ int erwartet und die auch wieder ein int liefert.
func Sum(n int) int {
	result := 0
	for i := 1; i <= n; i++ {
		result = result + i
	}
	return result
}
