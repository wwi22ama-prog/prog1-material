package operatoren

import "fmt"

// Die "klassischen" arithmetischen Operatoren.
func ArithmetischeOperatoren() {

	fmt.Println(12 + 4)
	fmt.Println(12 - 4)
	fmt.Println(12 * 4)
	fmt.Println(12 / 4)
}

// Besonderheiten beim Umgang mit Kommazahlen.
func ArithmetischeOperatorenBesonderheiten() {
	// Vorsicht bei der Division mit ganzen Zahlen:
	fmt.Println(12 / 5) // Der Rest wird weggeworfen.
	fmt.Println(12 % 5) // Der "modulo"-Operator liefert den Rest.

	// Es gibt auch Kommazahlen, man muss sie aber explizit benutzen:
	fmt.Println(12.0 / 5)
}

// Operatoren funktionieren auch mit Variablen und können verschachtelt werden.
func ArithmetischeOperatorenVariablen() {
	x1 := 42
	fmt.Println(x1 / 6)
	fmt.Println((x1/6 + 3) * 5)

}

// Boolesche Operatoren und Vergleiche:
func BoolescheOperatoren() {

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
