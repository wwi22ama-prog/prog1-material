package listfunctions

import "loesungen/excelfunktionen/misc"

// Erwartet einen Suchstring (query) sowie eine Liste von Strings (list).
// Liefert die Position, an der query in list vorkommt, falls vorhanden.
// Liefert -1, falls query nicht in list vorkommt.
func FindString(query string, list []string) int {
	for i, v := range list {
		if v == query {
			return i
		}
	}
	return -1
}

// Erwartet einen Suchstring (Query) sowie zwei Listen von Strings (keys, values).
// Liefert den String aus values, der der Position von query in keys entspricht.
// Falls query nicht in keys vorkommt, soll der leere String geliefert werden.
func LookupString(query string, keys, values []string) string {
	pos := FindString(query, keys)
	if pos == -1 {
		return ""
	}
	return values[pos]
}

// Erwartet eine Liste von Zahlen und liefert die Summe aller dieser Zahlen.
func SumInt(list []int) int {
	result := 0
	for _, v := range list {
		result += v
	}
	return result
}

// Erwartet eine Liste von Zahlen und liefert das Maximum aller dieser Zahlen.
func MaxInt(list []int) int {
	if len(list) == 0 {
		return 0
	}
	result := list[0]
	for _, v := range list[1:] {
		if v > result {
			result = v
		}
	}
	return result
}

// Erwartet eine 2D-Tabelle von Zahlen und liefert die Summe aller dieser Zahlen.
func SumIntTable(list [][]int) int {
	return SumInt(misc.FlattenTableInt(list))
}

// Erwartet eine 2D-Tabelle von Zahlen und liefert das Maximum aller dieser Zahlen.
func MaxIntTable(list [][]int) int {
	return MaxInt(misc.FlattenTableInt(list))
}
