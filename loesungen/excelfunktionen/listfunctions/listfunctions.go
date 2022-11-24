package listfunctions

// Erwartet einen Suchstring (query) sowie eine Liste von Strings (list).
// Liefert die Position, an der query in list vorkommt, falls vorhanden.
// Liefert -1, falls query nicht in list vorkommt.
func FindString(query string, list []string) int {
	// TODO
	return -1
}

// Erwartet einen Suchstring (Query) sowie zwei Listen von Strings (keys, values).
// Liefert den String aus values, der der Position von query in keys entspricht.
// Falls query nicht in keys vorkommt, soll der leere String geliefert werden.
func LookupString(query string, keys, values []string) string {
	// TODO
	return ""
}

// Erwartet eine Liste von Zahlen und liefert die Summe aller dieser Zahlen.
func SumInt(list []int) int {
	result := 0
	// TODO
	return result
}

// Erwartet eine Liste von Zahlen und liefert das Maximum aller dieser Zahlen.
func MaxInt(list []int) int {
	result := list[0]
	// TODO
	return result
}

// Erwartet eine 2D-Tabelle von Zahlen und liefert die Summe aller dieser Zahlen.
func SumIntTable(list [][]int) int {
	// TODO
	return 0
}

// Erwartet eine 2D-Tabelle von Zahlen und liefert das Maximum aller dieser Zahlen.
func MaxIntTable(list [][]int) int {
	// TODO
	return 0
}
