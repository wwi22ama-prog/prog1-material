package variablen

import "fmt"

// Deklariert eine Variable x mit Typ int.
// Weist dieser Variable den Wert 42 zu und gibt x aus.
// Weist anschließend 23 zu und gibt x erneut aus.
func VarInt1() {
	var x int
	x = 42
	fmt.Println(x)

	x = 23
	fmt.Println(x)
}

// Definiert eine int-Variable y und weist ihr dabei direkt den Wert 55 zu.
// Gibt anschließend den Wert von y aus.
// Der Typ muss nicht angegeben werden, er wird vom Go-Compiler bestimmt.
func VarInt2() {
	y := 55
	fmt.Println(y)
}

// Definiert eine Variable s vom Typ string und weist ihr einen Wert zu.
// Gibt anschließend zuerst den Wert und dann den Typ von s aus.
func VarString() {
	s := "Hallo String"
	fmt.Println(s)
	fmt.Printf("%T\n", s)
}
