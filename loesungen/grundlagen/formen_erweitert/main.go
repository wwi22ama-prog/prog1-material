package main

import "fmt"

func main() {
	form := ReadForm()
	breite := ReadDim("Breite")
	if form == 1 {
		PrintSquare(breite)
	} else if form == 3 {
		PrintTriangle(breite)
	} else {
		hoehe := ReadDim("Höhe")
		PrintRectangle(breite, hoehe)
	}
}

// Fragt den Benutzer so lange nach einer Form, bis er etwas gültiges eingibt.
func ReadForm() int {
	form := 0
	fmt.Println("Bitte eine Form wählen:")
	fmt.Println("1 : Quadrat")
	fmt.Println("2 : Rechteck")
	fmt.Println("3 : Dreieck")
	fmt.Scanln(&form)
	fmt.Println()

	if form > 0 && form < 4 {
		return form
	}

	fmt.Println("Die Eingabe war ungültig, bitte noch einmal.")
	return ReadForm()
}

// Fragt den Benutzer so lange nach einer Dimension, bis er etwas gültiges eingibt.
func ReadDim(name string) int {
	dim := 0
	fmt.Printf("Bitte eine %v angeben (mindestens 1): ", name)
	fmt.Scanln(&dim)
	fmt.Println()

	if dim > 0 {
		return dim
	}

	fmt.Println("Die Eingabe war ungültig, bitte noch einmal.")
	return ReadDim(name)
}
