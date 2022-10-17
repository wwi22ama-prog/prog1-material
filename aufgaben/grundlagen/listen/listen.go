package main

func main() {}

// Erwartet eine Liste und liefert Position und Wert des größten Elements.
func Largest(l []int) (int, int) {
	// TODO
	return 0, 0
}

// Erwartet eine Liste l.
// Liefert eine neue Liste, die die Elemente von l in umgekehrter Reihenfolge enthält.
func Reverse(l []int) []int {
	result := make([]int, 0)
	// TODO
	return result
}

// Erwartet eine Liste und liefert die Summe aller Elemente der Liste.
func Sum(l []int) int {
	// TODO
	return 0
}

// Erwartet eine Liste und liefert den Durchschnitt aller Elemente der Liste.
func Average(l []int) float64 {
	// TODO
	return 0
}

// Erwartet eine Liste und liefert den Median der Liste.
func Median(l []int) int {
	// Hinweise:
	// - Der Median ist der mittlere Wert aller Elemente.
	//   Also der Wert, der bei der sortierten Liste genau in der Mitte stehen würde.
	// - Bei gerader Listenlänge soll das linke der beiden mittleren Elemente
	// verwendet werden.

	// TODO
	return 0
}

// Erwartet eine Liste und eine Zahl n.
// Liefert die Anzahl der Elemente, die größer sind als n.
func CountLarger(l []int, n int) int {
	// TODO
	return 0
}

// Erwartet eine Liste und eine Zahl n.
// Liefert eine neue Liste, aus der alle Vorkommen von n entfernt worden sind.
func Remove(l []int, n int) []int {
	result := make([]int, 0)
	// TODO
	return result
}

// Erwartet eine Zahl n.
// Liefert eine Liste mit den Ziffern von n.
func Digits(n int) []int {
	result := make([]int, 0)
	// TODO
	return result
}

// Erwartet zwei Listen l1 und l2.
// Liefert eine neue Liste, die an Stelle i die das Produkt l1[i] * l2[i] enthält.
// Die Tests nehmen an, dass die Listen gleich lang sind.
// Zusatzaufgabe: Erweitern Sie die Funktion und die Tests, so dass sie für ungleich
// lange Listen ein sinnvolles Verhalten haben.
func ElementProduct(l1, l2 []int) []int {
	result := make([]int, 0)
	// TODO
	return result
}

// Erwartet eine Liste.
// Liefert das Produkt der Elemente in dieser Liste.
func Product(l []int) int {
	result := 1

	for _, v := range l {
		result *= v
	}

	return result
}

// Erwartet eine Liste von Ziffern.
// Liefert deren ISBN-13-Prüfsumme.
func Isbn13Checksum(digits []int) int {
	// TODO
	return 0
}

// Erwarte eine Zahl n.
// Liefert die Quersumme von n.
func DigitSum(n int) int {
	// TODO
	return 0
}
