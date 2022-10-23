package tableaccess

// Erwartet eine Tabelle und eine Zell-Adresse und
// liefert den Wert an dieser Adresse als String.
func GetCellValueString(table [][]string, cellAddress string) string {
	// TODO
	return ""
}

// Erwartet eine Tabelle und eine Zell-Adresse und
// liefert den Wert an dieser Adresse als Zahl.
func GetCellValueInt(table [][]string, cellAddress string) int {
	// TODO
	return 0
}

// Erwartet eine Tabelle und eine Bereichs-Adresse.
// Liefert einen Ausschnitt der Tabelle, der nur auf diesen Bereich eingeschränkt ist.
func GetRange(table [][]string, rangeAddress string) [][]string {
	// TODO
	return table
}

// Erwartet eine Tabelle und eine Bereichs-Adresse.
// Liefert einen Ausschnitt der Tabelle, der nur auf diesen Bereich eingeschränkt ist.
// Liefert eine Int-Slice. Alle Elemente, die nicht nach Int konvertiert werden können,
// werden als 0 interpretiert.
func GetRangeInt(table [][]string, rangeAddress string) [][]int {
	result := [][]int{}
	// TODO
	return result
}
