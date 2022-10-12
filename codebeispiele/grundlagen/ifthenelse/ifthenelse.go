package main

import "fmt"

func main() {
	// Tests der Funktion Judge()
	Judge(42)
	Judge(13)

	// Tests der Funktion IsLarge()
	// Diese Funktion liefert entweder true oder false.
	// Diesen Datentyp nennt man bool.
	fmt.Println(IsLarge(42))
	fmt.Println(IsLarge(13))

	// Tests der Funktion PrintDivisors()
	PrintDivisors(3, 25)
	PrintDivisors(17, 100)
	PrintDivisors(2, 10)

	// Tests der Funktion SumNRec()
	fmt.Println(SumNRec(5))
	fmt.Println(SumNRec(13))
	fmt.Println(SumNRec(200))
}

// Erwartet eine Zahl n bewerte sie, ob es eine gute Zahl ist.
// Die Bewertung wird auf die Konsole geschrieben.
func Judge(n int) {
	if n == 42 {
		fmt.Printf("%v ist eine gute Zahl!\n", n)
	} else {
		fmt.Printf("%v ist keine gute Zahl!\n", n)
	}
}

// Erwartet eine Zahl n und liefert true, falls n eine große Zahl ist.
// Jubelt bei großen Zahlen außerdem auf der Konsole!
func IsLarge(n int) bool {
	if n > 25 {
		fmt.Println("Juhuuuu!")
		return true
	}
	return false

}

// Erwartet zwei Zahlen m und n.
// Gibt alle durch m teilbaren Zahlen zwischen 0 und n aus.
func PrintDivisors(m, n int) {
	for i := 0; i <= n; i++ {
		if i%m == 0 {
			fmt.Println(i)
		}
	}
}

// Alternative Variante der Summenfunktion aus den vorherigen Beispielen.
// Erwartet eine Zahl n und liefert die Summe der Zahlen von 1 bis n.
// Verwendet dafür aber keine Schleife, sondern eine if-Anweisung und
// einen sog. rekursiven Aufruf.
func SumNRec(n int) int {
	if n <= 1 {
		return n
	}
	return n + SumNRec(n-1)
}
