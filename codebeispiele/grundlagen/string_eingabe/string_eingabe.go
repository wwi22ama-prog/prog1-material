package main

import "fmt"

func main() {
	var name string
	fmt.Print("Gib deinen Namen ein: ")
	fmt.Scanln(&name)
	fmt.Println()

	fmt.Printf("Hallo %v!", name)
}
