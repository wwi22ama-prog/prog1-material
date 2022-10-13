package main

import "fmt"

func main() {
	var input int

	fmt.Print("Bitte eine Zahl eingeben: ")
	fmt.Scanln(&input)
	fmt.Println("Summe:", Sum(input))
}

// Funktion, die einen Parameter n vom Typ int erwartet und die auch wieder ein int liefert.
func Sum(n int) int {
	result := 0
	for i := 1; i <= n; i++ {
		result = result + i
	}
	return result
}
