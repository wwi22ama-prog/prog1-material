package strings

import "fmt"

// Fragt den Benutzer nach seinem Namen und begrüßt ihn.
func Greet() {
	var name string
	fmt.Print("Gib deinen Namen ein: ")
	fmt.Scanln(&name)
	fmt.Println()

	fmt.Printf("Hallo %v!\n", name)
}
