package main

import "fmt"

// main()-Funktion. Einstiegspunkt ins Programm.
// Diese Funktion ruft verschiedene andere Funktionen auf, die dann Dinge tun.
// Die aufgerufenen Funktionen sind (fast alle) weiter unten definiert.
// Wenn von main() aus eine Funktion aufgerufen wird, springt der "Kontrollfluss"
// (d.h. der Programmablauf) in diese Funktion und führt den dortigen Code aus.
// Am Ende der Funktion kehrt das Programm an die aufrufende Stelle zurück und es
// geht von dort aus weiter.
func main() {
	// Aufruf der Funktion Greet().
	Greet()

	// Zwei verschiedene Aufrufe der Funktion PrintSum().
	// Die Zahlen, die hier in den Klammern stehen, heißen "Argumente".
	// Sie werden innerhalb der Funktion jeweils für x und y eingesetzt.
	PrintSum(3, 4)
	PrintSum(12, -2)

	// Verschiedene Aufrufe der Funktion ComputeSum().
	// Die Funktion gibt selbst nichts aus, aber mit ihrem Ergebnis kann
	// weitergerechnet werden.
	// D.h. es kann wieder als Argument an andere Funktionen gegeben werden,
	// oder es kann in eine Variable geschrieben werden.
	// Damit wir den Effekt der Funktion sehen können, muss das Ergebnis letzlich
	// in irgendeiner Print()-Anweisung landen.
	fmt.Println(ComputeSum(3, 4))
	sum1 := ComputeSum(15, 8)
	fmt.Println(sum1)
	sum2 := ComputeSum(25, 14)
	sum3 := ComputeSum(sum1, sum2)
	PrintSum(sum3, ComputeSum(13, sum2))

	// Aufrufe der Funktion PrintTwice().
	// Hier passiert im Grunde nichts neues, aber dieses ist eine Funktion,
	// die anstelle einer Zahl einen String erwartet.
	PrintTwice("Hallo")
	PrintTwice("Dies ist ein längerer String.")

	// Aufrufe der Funktion PrintSep().
	// Gibt jeweils die beiden Strings aus und sorgt durch die Wahl des
	// dritten Arguments dafür, dass diese Strings auf verschiedene Weise
	// voneinander getrennt dargestellt werden.
	PrintSep("Hallo", "Welt", " ")
	PrintSep("Käse", "kuchen", "")
	PrintSep("Zeile1", "Zeile2", "\n")

	// Aufrufe der Funktion ConcatSep().
	// Hier gibt die Funktion wieder nichts auf der Konsole aus.
	// Daher müssen wir die Ausgabe selbst machen.
	// Dafür ist die Funktion flexibler nutzbar, weil mit ihrem Ergebnis
	// weitergearbeitet werden kann.
	s1 := ConcatSep("Hallo", "Welt", " ")
	fmt.Println(s1)
	fmt.Println(ConcatSep("Käse", "kuchen", ""))

	// Sehr komplexer Aufruf von ConcatSep().
	// Dieser Aufruf könnte auch auf einer Zeile stehen, wäre dann aber
	// sehr viel schwerer lesbar. Hier steht jeweils ein Aufruf von ConcatSep()
	// und darunter auf einer Ebene eingerückt die Argumente für diesen Aufruf.
	// Diese Argumente können selbst wieder Ergebnisse von Funktionen sein.
	sep := "\n"
	fmt.Println(
		ConcatSep(
			ConcatSep(
				"Zeile3",
				"Zeile4",
				sep),
			ConcatSep(
				"Zeile5",
				ConcatSep(
					"Zeile6",
					"Zeile7",
					sep),
				sep),
			sep),
	)
}

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
