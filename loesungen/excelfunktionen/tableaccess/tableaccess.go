package tableaccess

import (
	"loesungen/excelfunktionen/addressfunctions"
	"loesungen/excelfunktionen/misc"
	"strconv"
)

// Erwartet eine Tabelle und eine Zell-Adresse und
// liefert den Wert an dieser Adresse als String.
func GetCellValueString(table [][]string, cellAddress string) string {
	row := addressfunctions.RowNumber(cellAddress)
	column := addressfunctions.ColumnNumber(cellAddress)

	return table[row-1][column-1]
}

// Erwartet eine Tabelle und eine Zell-Adresse und
// liefert den Wert an dieser Adresse als Zahl.
func GetCellValueInt(table [][]string, cellAddress string) int {
	result, _ := strconv.Atoi(GetCellValueString(table, cellAddress))
	return result
}

// Erwartet eine Tabelle und eine Bereichs-Adresse.
// Liefert einen Ausschnitt der Tabelle, der nur auf diesen Bereich eingeschränkt ist.
func GetRange(table [][]string, rangeAddress string) [][]string {
	firstRow := addressfunctions.FirstRow(rangeAddress)
	lastRow := addressfunctions.LastRow(rangeAddress)

	firstColumn := addressfunctions.FirstColumn(rangeAddress)
	lastColumn := addressfunctions.LastColumn(rangeAddress)

	table = table[firstRow-1 : lastRow]
	transposed := misc.Transpose(table)
	transposed = transposed[firstColumn-1 : lastColumn]
	return misc.Transpose(transposed)
}

// Erwartet eine Tabelle und eine Bereichs-Adresse.
// Liefert einen Ausschnitt der Tabelle, der nur auf diesen Bereich eingeschränkt ist.
// Liefert eine Int-Slice. Alle Elemente, die nicht nach Int konvertiert werden können,
// werden als 0 interpretiert.
func GetRangeInt(table [][]string, rangeAddress string) [][]int {
	stringRange := GetRange(table, rangeAddress)

	result := [][]int{}
	for rowIndex, row := range stringRange {
		result = append(result, []int{})
		for _, entry := range row {
			entryAsInt, _ := strconv.Atoi(entry)
			result[rowIndex] = append(result[rowIndex], entryAsInt)
		}
	}
	return result
}
