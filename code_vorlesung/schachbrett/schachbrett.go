package main

import "fmt"

// Gibt ein Schachbrettmuster aus nxn Quadraten aus,
// bei dem jedes Quadrat aus mxm Zeichen besteht.
func PrintChessBoard(m, n int) {
	for row := 0; row < n; row++ {
		if row%2 == 0 {
			PrintChessBoardRow(m, n, "*", "+")
		} else {
			PrintChessBoardRow(m, n, "+", "*")
		}
	}
}

// Gibt eine Reihe eines Schachbrettmuster aus n Quadraten aus,
// bei dem jedes Quadrat aus mxm Zeichen besteht.
func PrintChessBoardRow(m, n int, firstFill, secondFill string) {
	for i := 0; i < m; i++ {
		for j := 0; j < n*m; j++ {
			squareNumber := j / m
			if squareNumber%2 == 0 {
				fmt.Print(firstFill)
			} else {
				fmt.Print(secondFill)
			}
		}
		fmt.Println()
	}
}
