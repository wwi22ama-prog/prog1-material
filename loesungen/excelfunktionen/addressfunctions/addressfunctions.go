package addressfunctions

import (
	"loesungen/excelfunktionen/misc"
	"strconv"
	"strings"
)

// Erwartet eine Zell-Adresse und liefert deren Zeilennummer.
func RowNumber(cellAddress string) int {
	rowAddress := misc.FilterDigits(cellAddress)
	result, _ := strconv.Atoi(rowAddress)
	return result
}

// Erwartet eine Zell-Adresse und liefert deren Spaltennummer.
func ColumnNumber(cellAddress string) int {
	columnAddress := misc.FilterUpperCase(cellAddress)
	result := 0
	for len(columnAddress) > 0 {
		result *= 26
		result += int(columnAddress[0] - 'A' + 1)
		columnAddress = columnAddress[1:]
	}
	return result
}

// Erwartet eine Bereichs-Adresse und liefert die Nummer der ersten Zeile.
func FirstRow(rangeAddress string) int {
	return RowNumber(strings.Split(rangeAddress, ":")[0])
}

// Erwartet eine Bereichs-Adresse und liefert die Nummer der ersten Spalte.
func FirstColumn(rangeAddress string) int {
	return ColumnNumber(strings.Split(rangeAddress, ":")[0])
}

// Erwartet eine Bereichs-Adresse und liefert die Nummer der letzten Zeile.
func LastRow(rangeAddress string) int {
	return RowNumber(strings.Split(rangeAddress, ":")[1])
}

// Erwartet eine Bereichs-Adresse und liefert die Nummer der letzten Spalte.
func LastColumn(rangeAddress string) int {
	return ColumnNumber(strings.Split(rangeAddress, ":")[1])
}

// Erwartet eine Bereichs-Adresse und liefert die Anzahl der Zeilen.
func RowCount(rangeAddress string) int {
	return LastRow(rangeAddress) - FirstRow(rangeAddress) + 1
}

// Erwartet eine Bereichs-Adresse und liefert die Anzahl der Spalten.
func ColumnCount(rangeAddress string) int {
	return LastColumn(rangeAddress) - FirstColumn(rangeAddress) + 1
}

// Erwartet einen Adress-Ausdruck und liefert true,
// falls es sich dabei um einen Bereich handelt.
func IsRange(address string) bool {
	parts := strings.SplitN(address, ":", 2)
	return len(parts) == 2 && IsCell(parts[0]) && IsCell(parts[1])
}

// Erwartet einen Adress-Ausdruck und liefert true,
// falls es sich dabei um eine Zelle handelt.
func IsCell(address string) bool {
	lettersSeen := false
	digitsSeen := false
	for _, v := range address {
		if !misc.IsUpperCase(v) && !misc.IsDigit(v) {
			return false
		}
		if misc.IsUpperCase(v) && digitsSeen {
			return false
		}
		lettersSeen = lettersSeen || misc.IsUpperCase(v)
		digitsSeen = digitsSeen || misc.IsDigit(v)
	}
	return lettersSeen && digitsSeen
}
