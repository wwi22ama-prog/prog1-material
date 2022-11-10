package main

import "fmt"

func main() {
	// 1-D-Int-Slice
	intslice1d := []int{2, 3, 4}

	fmt.Println(intslice1d)

	// 2D-Int-slice ist eine Liste von 1D-Slices:
	intslice2d := [][]int{{2, 3, 4}, {5, 6, 7}}
	fmt.Println(intslice2d)

	// Leere 2D-Slice:
	intslice2d_leer := [][]int{}
	//	fmt.Println(intslice2d_leer)

	// Drei Zeilen anhängen:
	intslice2d_leer = append(intslice2d_leer, []int{42})
	intslice2d_leer = append(intslice2d_leer, []int{25})
	intslice2d_leer = append(intslice2d_leer, []int{38})

	// An Zeile 1 noch ein Element anhängen:
	intslice2d_leer[1] = append(intslice2d_leer[1], 77)

	// Ein Element an die 38 anhängen:
	intslice2d_leer[2] = append(intslice2d_leer[2], 77)

	fmt.Println(intslice2d_leer)

	// Alle Zeilen ausgeben:
	for _, line := range intslice2d_leer {
		fmt.Println(line)
	}

	// Alle Elemente ohne Klammern mit Sternen ausgeben:
	for _, line := range intslice2d_leer {
		for _, element := range line {
			fmt.Print(element, "*")
		}
		fmt.Println()
	}

	// Das Element an einer bestimmten Stelle ausgeben
	fmt.Println(intslice2d_leer[1][0])
}
