package main

import "fmt"

func main() {}

// Erwartet zwei Zahlen m und n.
// Liefert true, falls m ein Teiler von n ist.
// Eine Lösung ist vorgegeben, die man auch in der Praxis verwenden würde.
// Aufgabe: Ersetzen Sie diese Lösung durch eine, die den Modulo-Operator nicht verwendet.
func Divides(m, n int) bool {
	if m == 0 {
		return false
	}
	if n < 0 {
		n = -n
	}
	if m < 0 {
		m = -m
	}
	for n >= m {
		n -= m
	}
	return n == 0
}

// Erwartet eine Zahl n.
// Liefert true, falls n eine Primzahl ist.
func IsPrime(n int) bool {
	if n <= 1 {
		return false
	}
	for i := 2; i < n; i++ {
		if Divides(i, n) {
			return false
		}
	}
	return true
}

// Erwartet eine Zahl n.
// Gibt alle Primzahlen auf der Konsole aus, die kleiner als n sind.
func PrintPrimes(n int) {
	for i := 2; i < n; i++ {
		if IsPrime(i) {
			fmt.Println(i)
		}
	}
}

// Erwartet eine Zahl n.
// Liefert die nächstgrößere Primzahl.
// Liefert n, falls n selbst eine Primzahl ist.
func NextPrime(n int) int {
	result := n
	for ; !IsPrime(result); result++ {
	}
	return result
}

// Erwartet eine Zahl n.
// Liefert den nächsten Primzahlzwilling, der größer ist als n.
// Genauer: Liefert die kleinste Zahl k, so dass:
// * k >= n
// * k ist eine Primzahl
// * k + 2 ist eine Primzahl
func NextPrimeTwin(n int) int {
	for n = NextPrime(n); !IsPrime(n + 2); n = NextPrime(n + 1) {
	}
	return n
}

// Erwartet eine Zahl n.
// Liefert die größte Primzahl, die echt kleiner als n ist.
// Falls es keine solche Zahle gibt, wird 0 geliefert.
func GreatestPrimeBelow(n int) int {
	if n <= 2 {
		return 0
	}
	for ; !IsPrime(n); n-- {
	}
	return n
}
