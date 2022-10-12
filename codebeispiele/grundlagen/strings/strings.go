package main

import "fmt"

func main() {
	s1 := "Hallo"
	s2 := "String"

	// Ausgabe von Strings mittels fmt.Println().
	// Die folgenden Kommandos geben einzelne Strings aus oder solche, die aus anderen
	// zusammengesetzt sind.
	fmt.Println("Huhu")
	fmt.Println(s1)
	fmt.Println(s2)
	fmt.Println(s1 + s2)
	fmt.Println(s1 + " " + s2)

	// Wir können die Länge eines Strings bestimmen:
	fmt.Println(len(s1))
	fmt.Println(len(s2))
	fmt.Println(len(s1 + s2))

	// Man kann auch einzelne Buchstaben von Strings verwenden.
	// Allerdings ist das Ergebnis etwas überraschend...
	fmt.Println(s1[2])
	for i, c := range s1 {
		fmt.Printf("%v: %v\n", i, c)
	}

	// Besser ist es auf diese Weise:
	for i, c := range s1 {
		fmt.Printf("%v: %v\n", i, string(c))
	}

	// Es gibt auch Ausschnitte aus Strings, die mehr als einen Buchstaben enthalten.
	s3 := s1[0:3]
	fmt.Println(s3)
	s4 := s1[:3]
	fmt.Println(s4)
	s5 := s1[2:5]
	fmt.Println(s5)
	s6 := s1[2:]
	fmt.Println(s6)

	// Eine Schleife, die einen String mit dem Alphabet vollschreibt:
	s7 := ""
	for c := 'a'; c <= 'z'; c++ {
		s7 += string(c)
	}
	fmt.Println(s7)

}
