package schleifen

import "fmt"

// Erwartet eine Zahl n und zählt von 0 bis zu dieser Zahl.
func CountTo(n int) {
	for i := 0; i <= n; i++ {
		fmt.Println(i)
	}
}

// Erwartet eine Zahl n und eine Schrittweite step.
// Zählt von 0 bis zu dieser Zahl mit der angegebenen Schrittweite.
func CountToStep(n, step int) {
	for i := 0; i <= n; i += step {
		fmt.Println(i)
	}
}

// Erwartet eine Zahl n und zählt von n bis 0 herunter.
func CountDown(n int) {
	for ; n >= 0; n-- {
		fmt.Println(n)
	}
}

// Erwartet einen String s und eine Zahl count.
// Liefert einen neuen String, indem s count Mal hintereinandergehängt wird.
func ConcatNTimes(s string, count int) string {
	result := ""
	for i := 0; i < count; i++ {
		result = result + s
	}
	return result
}

// Erwartet eine Zahl n und liefert die Summe der Zahlen von 1 bis n.
func SumN(n int) int {
	result := 0
	for i := 1; i <= n; i++ {
		result += i
	}
	return result
}

// Erwartet so lange immer wieder Eingaben vom Benutzer, bis dieser eine Null eingibt.
// Schreibt nach jeder Eingabe das Doppelte dessen auf die Konsole, was der Benutzer
// eingegeben hat.
func DisplayDoubles() {
	input := 1

	for input != 0 {
		fmt.Print("Bitte eine Zahl eingeben: ")
		fmt.Scanln(&input)
		fmt.Printf("%v\n\n", 2*input)
	}
}

// Erwartet einen String s.
// Verwendet eine range-for-Schleife, um jeden Buchstaben von s
// zusammen mit dessen Position einzeln auszugeben.
func PrintStringLetters(s string) {
	for i, v := range s {
		fmt.Printf("%v: %c, ", i, v)
	}
}
