package funktionen

import "fmt"

// Aufruf der Funktion Greet().
func ExampleGreet() {
	Greet()

	// Output:
	// Funktion 'greet()' aufgerufen.
}

// Zwei verschiedene Aufrufe der Funktion PrintSum().
// Die Zahlen, die hier in den Klammern stehen, heißen "Argumente".
// Sie werden innerhalb der Funktion jeweils für x und y eingesetzt.
func ExamplePrintSum() {
	PrintSum(3, 4)
	PrintSum(12, -2)

	// Output:
	// 3 + 4 == 7
	// 12 + -2 == 10
}

// Verschiedene Aufrufe der Funktion ComputeSum().
// Die Funktion gibt selbst nichts aus, aber mit ihrem Ergebnis kann
// weitergerechnet werden.
// D.h. es kann wieder als Argument an andere Funktionen gegeben werden,
// oder es kann in eine Variable geschrieben werden.
// Damit wir den Effekt der Funktion sehen können, muss das Ergebnis letzlich
// in irgendeiner Print()-Anweisung landen.
func ExampleComputeSum() {
	fmt.Println(ComputeSum(3, 4))
	sum1 := ComputeSum(15, 8)
	fmt.Println(sum1)
	sum2 := ComputeSum(25, 14)
	sum3 := ComputeSum(sum1, sum2)
	PrintSum(sum3, ComputeSum(13, sum2))

	// Output:
	// 7
	// 23
	// 62 + 52 == 114
}

// Aufrufe der Funktion PrintTwice().
// Hier passiert im Grunde nichts neues, aber dieses ist eine Funktion,
// die anstelle einer Zahl einen String erwartet.
func ExamplePrintTwice() {

	PrintTwice("Hallo")
	PrintTwice("Dies ist ein längerer String.")

	// Output:
	// HalloHallo
	// Dies ist ein längerer String.Dies ist ein längerer String.
}

// Aufrufe der Funktion PrintSep().
// Gibt jeweils die beiden Strings aus und sorgt durch die Wahl des
// dritten Arguments dafür, dass diese Strings auf verschiedene Weise
// voneinander getrennt dargestellt werden.
func ExamplePrintSep() {
	PrintSep("Hallo", "Welt", " ")
	PrintSep("Käse", "kuchen", "")
	PrintSep("Zeile1", "Zeile2", "\n")

	// Output:
	// Hallo Welt
	// Käsekuchen
	// Zeile1
	// Zeile2
}

// Aufrufe der Funktion ConcatSep().
// Hier gibt die Funktion wieder nichts auf der Konsole aus.
// Daher müssen wir die Ausgabe selbst machen.
// Dafür ist die Funktion flexibler nutzbar, weil mit ihrem Ergebnis
// weitergearbeitet werden kann.
func ExampleConcatSep_simple() {
	s1 := ConcatSep("Hallo", "Welt", " ")
	fmt.Println(s1)
	fmt.Println(ConcatSep("Käse", "kuchen", ""))

	// Output:
	// Hallo Welt
	// Käsekuchen
}

// Sehr komplexer Aufruf von ConcatSep().
// Dieser Aufruf könnte auch auf einer Zeile stehen, wäre dann aber
// sehr viel schwerer lesbar. Hier steht jeweils ein Aufruf von ConcatSep()
// und darunter auf einer Ebene eingerückt die Argumente für diesen Aufruf.
// Diese Argumente können selbst wieder Ergebnisse von Funktionen sein.
func ExampleConcatSep_complex() {
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

	// Output:
	// Zeile3
	// Zeile4
	// Zeile5
	// Zeile6
	// Zeile7
}
