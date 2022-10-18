package main

import (
	"sort"
)

func main() {}

// Erwartet eine Liste und liefert Position und Wert des größten Elements.
func Largest(l []int) (int, int) {
	largestPos := 0
	largestValue := l[0]

	for i, v := range l {
		if v > largestValue {
			largestPos = i
			largestValue = v
		}
	}

	return largestPos, largestValue
}

// Erwartet eine Liste l.
// Liefert eine neue Liste, die die Elemente von l in umgekehrter Reihenfolge enthält.
func Reverse(l []int) []int {
	length := len(l)
	result := make([]int, length)
	for i, v := range l {
		result[length-i-1] = v
	}
	return result
}

// Erwartet eine Liste und liefert die Summe aller Elemente der Liste.
func Sum(l []int) int {
	result := 0
	for _, v := range l {
		result += v
	}
	return result
}

// Erwartet eine Liste und liefert den Durchschnitt aller Elemente der Liste.
func Average(l []int) float64 {
	return float64(Sum(l)) / float64(len(l))
}

// Erwartet eine Liste und liefert den Median der Liste.
func Median(l []int) int {
	// Hinweise:
	// - Der Median ist der mittlere Wert aller Elemente.
	//   Also der Wert, der bei der sortierten Liste genau in der Mitte stehen würde.
	// - Bei gerader Listenlänge soll das linke der beiden mittleren Elemente
	// verwendet werden.

	sortedList := make([]int, 0)
	sortedList = append(sortedList, l...)
	sort.Ints(sortedList)

	length := len(sortedList)
	medianPos := length / 2

	if length%2 == 0 {
		return sortedList[medianPos-1]
	}

	return sortedList[medianPos]
}

// Erwartet eine Liste und eine Zahl n.
// Liefert die Anzahl der Elemente, die größer sind als n.
func CountLarger(l []int, n int) int {
	result := 0

	for _, v := range l {
		if v > n {
			result++
		}
	}

	return result
}

// Erwartet eine Liste und eine Zahl n.
// Liefert eine neue Liste, aus der alle Vorkommen von n entfernt worden sind.
func Remove(l []int, n int) []int {
	result := make([]int, 0)

	for _, v := range l {
		if v != n {
			result = append(result, v)
		}
	}

	return result
}

// Erwartet eine Zahl n.
// Liefert eine Liste mit den Ziffern von n.
// Die Funktion muss nur für positive n funktionieren.
func Digits(n int) []int {
	if n == 0 {
		return []int{0}
	}
	result := make([]int, 0)

	for ; n != 0; n /= 10 {
		result = append(result, n%10)
	}
	return Reverse(result)
}

// Erwartet zwei Listen l1 und l2.
// Liefert eine neue Liste, die an Stelle i die das Produkt l1[i] * l2[i] enthält.
// Die Tests nehmen an, dass die Listen gleich lang sind.
// Zusatzaufgabe: Erweitern Sie die Funktion und die Tests, so dass sie für ungleich
// lange Listen ein sinnvolles Verhalten haben.
func ElementProduct(l1, l2 []int) []int {
	result := make([]int, 0)

	for i := range l1 {
		result = append(result, l1[i]*l2[i])
	}

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
	weights := []int{1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3}
	weightedSum := Sum(ElementProduct(digits, weights))
	return 10 - weightedSum%10
}

// Erwarte eine Zahl n.
// Liefert die Quersumme von n.
func DigitSum(n int) int {
	return Sum(Digits(n))
}
