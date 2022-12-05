package spreadsheetfunctions

import (
	"fmt"
	"loesungen/excelfunktionen/listfunctions"
	"loesungen/excelfunktionen/misc"
	"loesungen/excelfunktionen/tableaccess"
	"loesungen/excelfunktionen/tablefunctions"
	"strconv"
)

// Erwartet eine Tabelle und einen Funktionsaufruf.
// Führt die geforderte Berechnung durch und liefert das Ergebnis als String.
func ComputeFunction(table [][]string, call string) string {
	functionName := tablefunctions.GetFunctionName(call)

	if functionName == "SUMME" {
		return ComputeSum(table, tablefunctions.GetArgument(call, 0))
	}
	if functionName == "MAX" {
		return ComputeMax(table, tablefunctions.GetArgument(call, 0))
	}
	if functionName == "LOOKUP" {
		return ComputeLookup(
			table,
			tablefunctions.GetArgument(call, 0),
			tablefunctions.GetArgument(call, 1),
			tablefunctions.GetArgument(call, 2),
		)
	}
	if functionName == "VLOOKUP" {
		column, _ := strconv.Atoi(tablefunctions.GetArgument(call, 2))
		return ComputeVLookup(
			table,
			tablefunctions.GetArgument(call, 0),
			tablefunctions.GetArgument(call, 1),
			column,
		)
	}

	return ""
}

// Erwartet eine Tabelle und einen Bereichs-Ausdruck.
// Berechnet die Summe der Elemente im Bereich und liefert das Ergebnis als String.
func ComputeSum(table [][]string, rangeAddress string) string {
	result := listfunctions.SumIntTable(tableaccess.GetRangeInt(table, rangeAddress))
	return fmt.Sprintf("%v", result)
}

// Erwartet eine Tabelle und einen Bereichs-Ausdruck.
// Bestimmt das größte der Elemente im Bereich und liefert das Ergebnis als String.
func ComputeMax(table [][]string, rangeAddress string) string {
	result := listfunctions.MaxIntTable(tableaccess.GetRangeInt(table, rangeAddress))
	return fmt.Sprintf("%v", result)
}

// Erwartet eine Tabelle, einen Zell- und zwei Bereichs-Ausdrücke.
// Löst den Zell-Ausdrucks auf und sich das Ergebnis im ersten Bereichs-Ausdruck.
// Liefert das entsprechende Element aus dem zweiten Bereichs-Ausdruck.
// Kommt das Element nicht vor, wird ein leerer String geliefert.
// Die beiden Bereiche müssen eindimensional sein!
func ComputeLookup(table [][]string, queryCell, keyRange, valueRange string) string {
	query := tableaccess.GetCellValueString(table, queryCell)
	keys := misc.FlattenTableString(tableaccess.GetRange(table, keyRange))
	values := misc.FlattenTableString(tableaccess.GetRange(table, valueRange))
	return listfunctions.LookupString(query, keys, values)
}

// Erwartet eine Tabelle, einen Zell- und einen Bereichs-Ausdruck sowie eine
// Spaltennummer. Löst den Zell-Ausdruck auf und sucht das Ergebnis in der ersten
// Spalte des Bereichs-Ausdrucks.
// Liefert den Eintrag aus der angegebenen Spalte des Bereichs-Ausdrucks.
// Kommt das Element nicht vor, wird ein leerer String geliefert.
func ComputeVLookup(table [][]string, queryCell, rangeAddress string, column int) string {
	query := tableaccess.GetCellValueString(table, queryCell)
	tableRange := tableaccess.GetRange(table, rangeAddress)
	keys := misc.GetColumn(tableRange, 0)
	values := misc.GetColumn(tableRange, column-1)
	return listfunctions.LookupString(query, keys, values)
}
