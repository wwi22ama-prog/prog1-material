package main

import "fmt"

// Gibt einen String auf der Konsole aus.
func Greet() {
	fmt.Println("Funktion 'greet()' aufgerufen.")
}

// Gibt die Summe von x und y auf der Konsole aus.
func PrintSum(x, y int) {
	fmt.Printf("%v + %v == %v\n", x, y, x+y)
}

// Liefert die Summe von x und y
func ComputeSum(x, y int) int {
	return x + y
}

// Erwartet einen String und gibt ihn zwei Mal
// direkt hintereinander auf der Konsole aus.
func PrintTwice(s string) {
	fmt.Printf("%v%v\n", s, s)
}

// Erwartet drei Strings s1, s2 und sep.
// Gibt s1 und s2 auf der Konsole aus und verwendet dabei sep als Trenner dazwischen.
func PrintSep(s1, s2, sep string) {
	fmt.Printf("%v%v%v\n", s1, sep, s2)
}

// Erwartet drei Strings s1, s2 und sep.
// Liefert den zusammengesetzten String, der aus s1, sep und s2 besteht.
func ConcatSep(s1, s2, sep string) string {
	return s1 + sep + s2
}
