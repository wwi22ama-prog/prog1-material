package main

import "fmt"

// Erwartet eine Zahl n sowie zwei Strings border und fill.
// Gibt eine Zeile der Länge n aus, deren linker und rechter Rand
// aus dem String border bestehen und deren Inneres aus fill besteht.
// Nach der Zeile wird ein Zeilenumbruch ausgegeben.
func PrintCustomRow(n int, border, fill string) {
	fmt.Print(border)
	for i := 0; i < n-2; i++ {
		fmt.Print(fill)
	}
	if n >= 2 {
		fmt.Print(border)
	}
	fmt.Println()
}

// Gibt ein benutzerdefiniertes Rechteck der Breite m und der Höne n
// aus, dessen Begrenzungen bzw. Inneres durch die Strings border und fill
// gegeben sind.
// Funktioniert nur korrekt für m,n >= 2
func PrintCustomRectangle(m, n int, border, fill string) {
	PrintCustomRow(m, border, border)
	for i := 0; i < n-2; i++ {
		PrintCustomRow(m, border, fill)
	}
	if n >= 2 {
		PrintCustomRow(m, border, border)
	} else {
		fmt.Println()
	}
}
