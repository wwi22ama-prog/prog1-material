package funktionen

import "fmt"

// Gibt einen String auf der Konsole aus.
func Greet() {
	fmt.Println("Funktion 'greet()' aufgerufen.")
}

// Gibt die Summe von x und y auf der Konsole aus.
// Variablen, die im Funktionskopf stehen, heißen "Parameter".
// Für sie müssen beim Aufruf passende Argumente angegeben werden
// (siehe Aufrufe in der main()-Funktion).
func PrintSum(x, y int) {
	fmt.Printf("%v + %v == %v\n", x, y, x+y)
}

// Gibt nichts auf der Konsole aus, aber berechnet die Summe von x und y
// und liefert das Ergebnis an die aufrufende Funktion zurück.
// Hier steht hinter den Parametern noch ein Rückgabetyp.
// Hier wird festgelegt, dass das Ergebnis dieser Funktion eine Zahl ist.
// Dieser Typ muss zum Ausdruck hinter der return-Anweisung passen.
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
