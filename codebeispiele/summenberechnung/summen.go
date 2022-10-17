package main

import "fmt"

// Leere main() für Experimente.
func main() {}

// Berechnet die Summe der Zahlen von 1 bis 5 und gibt das Ergebnis auf die Konsole aus.
func Summe1Bis5() {
	result := 1
	result = result + 2
	result = result + 3
	result = result + 4
	result = result + 5
	fmt.Println("Summe:", result)
}

// Berechnet die Summe der Zahlen von 1 bis n und liefert das Ergebnis zurück.
func Summe1BisN(n int) int {
	result := 0
	for i := 1; i <= n; i++ {
		result += i
	}
	return result
}

// Alternative Variante der Summenfunktion.
// Verwendet aber keine Schleife, sondern eine if-Anweisung und
// einen sog. rekursiven Aufruf.
func Sum1BisNRekursiv(n int) int {
	if n <= 1 {
		return n
	}
	return n + Sum1BisNRekursiv(n-1)
}
