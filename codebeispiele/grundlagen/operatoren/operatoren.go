package main

import "fmt"

func main() {
	// Ein Überblick über die wichtigsten Operatoren, die es in Go gibt.

	// Die "klassischen" arithmetischen Operatoren:
	fmt.Println(12 + 4)
	fmt.Println(12 - 4)
	fmt.Println(12 * 4)
	fmt.Println(12 / 4)

	// Vorsicht bei der Division mit ganzen Zahlen:
	fmt.Println(12 / 5) // Der Rest wird weggeworfen.
	fmt.Println(12 % 5) // Der "modulo"-Operator liefert den Rest.

	// Es gibt auch Kommazahlen, man muss sie aber explizit benutzen:
	fmt.Println(12.0 / 5)

	// Das geht natürlich auch mit Variablen und verschachtelt:
	x1 := 42
	fmt.Println(x1 / 6)
	fmt.Println((x1/6 + 3) * 5)

	// Boolesche Operatoren und Vergleiche:

	// Logisches "Und"
	fmt.Println(true && false)
	fmt.Println(42 < 103 && 15*3 > 42)

	// Logisches "Oder"
	fmt.Println(true || false)
	fmt.Println(42 < 103 || 15*3 > 42)

	// Negation
	fmt.Println(!true)
	fmt.Println(!false)
	fmt.Println(!(15 < 12 || 24 == 3*8))

	// Ungleich
	fmt.Println(23-15 != 13)
}
