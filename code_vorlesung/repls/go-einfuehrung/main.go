package main

import "fmt"

func main() {
	// Gibt aus, was in den Klammern steht.
	fmt.Print("Hello, World! ")

	// Gibt den Inhalt der Klammern aus und macht einen Zeilenumbruch.
	fmt.Println(42)

	fmt.Println(3 + 4)
	fmt.Println(3 < 4)
	fmt.Println(3 == 4)
	fmt.Println(3 != 4)

	// Deklaration einer Variablen vom Typ int.
	var x int

	// Zuweisung eines Wertes an x
	x = 23

	// x ausgeben
	fmt.Printf("x == %v\n", x)
	fmt.Printf("x == %v\n", "Hallo")
	fmt.Printf("%v, x == %v\n", "Hallo", x)

	// Den Typ von x ausgeben
	fmt.Printf("x ist vom Typ %T\n", x)
	fmt.Printf("32.5 ist vom Typ %T\n", 32.5)
}
