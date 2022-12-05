package misc

// Erwartet eine 2D-Tabelle (Rechteckige slice) und liefert die transponierte Tabelle.
func Transpose(table [][]string) [][]string {
	rowCount := len(table[0])
	columnCount := len(table)

	result := make([][]string, rowCount)
	for row := range result {
		result[row] = make([]string, columnCount)
	}
	for row := 0; row < rowCount; row++ {
		for column := 0; column < columnCount; column++ {
			result[row][column] = table[column][row]
		}
	}
	return result
}

// Erwartet eine 2D-Tabelle und eine Spaltennummer. Liefert die n-te Spalte als Liste.
func GetColumn(table [][]string, column int) []string {
	return Transpose(table)[column]
}

// Erwartet eine 2D-Tabelle aus Strings und macht eine "flache" 1D-Liste daraus.
func FlattenTableString(table [][]string) []string {
	result := []string{}
	for _, row := range table {
		result = append(result, row...)
	}
	return result
}

// Erwartet eine 2D-Tabelle aus Zahlen und macht eine "flache" 1D-Liste daraus.
func FlattenTableInt(table [][]int) []int {
	result := []int{}
	for _, row := range table {
		result = append(result, row...)
	}
	return result
}
