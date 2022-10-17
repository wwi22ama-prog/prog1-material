package main

import (
	"fmt"

	// Import der selbst definierten Packages, damit wir sie unten bneutzen können.
	// Das Import-Statement besteht hier aus dem Pfad aus der jeweiligen go.mod
	// und dem Namen des Packages.
	"codebeispiele/grundlagen/funktionen"
	"codebeispiele/grundlagen/strings"
)

// Für eigene Experimente.
func main() {

	// Ruft aus dem Package strings die Funktion Greet auf.
	// Das Package strings haben wir im Unterordner mit dem gleichen Namen definiert.
	// Damit der Aufruf funktioniert, muss das Package oben importiert werden.
	strings.Greet()

	// Ruft aus dem Package funktionen die Funktion
	fmt.Println(funktionen.ComputeSum(2, 5))
}
