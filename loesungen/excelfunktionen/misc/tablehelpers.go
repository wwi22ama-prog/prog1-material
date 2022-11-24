package misc

// Erwartet eine 2D-Tabelle (Rechteckige slice) und liefert die transponierte Tabelle.
func Transpose(table [][]string) [][]string {
	rowCount := len(table[0])
	result := make([][]string, rowCount)
	// TODO
	return result
}

// Erwartet eine 2D-Tabelle und eine Spaltennummer. Liefert die n-te Spalte als Liste.
func GetColumn(table [][]string, column int) []string {
	return []string{}
}

// Erwartet eine 2D-Tabelle aus Strings und macht eine "flache" 1D-Liste daraus.
func FlattenTableString(table [][]string) []string {
	result := []string{}
	// TODO
	return result
}

// Erwartet eine 2D-Tabelle aus Zahlen und macht eine "flache" 1D-Liste daraus.
func FlattenTableInt(table [][]int) []int {
	result := []int{}
	// TODO
	return result
}
