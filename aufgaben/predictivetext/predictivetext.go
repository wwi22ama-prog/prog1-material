package main

// Dies ist die bearbeitete Version der Aufgaben aus der Vorlesung.
// Eine komplette Lösung folgt auch noch unter "loesungen".

import "strings"

// Erwartet eine Liste von Strings dataBase und einen einzelnen String prefix.
// Liefert eine gefilterte mit genau den Strings aus dataBase, von denen prefix
// tatsächlich ein Präfix ist. D.h. die Strings, die mit prefix beginnen.
func FilterByPrefix(dataBase []string, prefix string) []string {
	result := []string{}
	for _, el := range dataBase {
		if strings.HasPrefix(el, prefix) {
			result = append(result, el)
		}
	}
	return result
}

// Erwartet zwei Strings s1 und s2. Prüft, ob s2 aus s1 entsteht,
// indem man ein Zeichen aus s1 weglässt.
func HasExtraLetter(s1, s2 string) bool {
	// Uninteressante Fälle ausschließen:
	// Wenn die Längen nicht passen, kann nicht mehr true herauskommen.
	if len(s2) != len(s1)-1 {
		return false
	}

	// Idee: Wir zählen die Unterschiede zwischen den Strings.
	// Dazu lassen eine Schleife über beide Strings laufen.
	// pos1 wandert durch s1, pos2 wandert durch s2.
	// Während pos1 immer wandert (also einfach der Schleifenzähler ist),
	// wandert pos2 nur, wenn die beiden Buchstaben an diesen Stellen gleich sind.
	differences := 0
	pos2 := 0
	for pos1 := range s1 { // for pos1 := 0; pos1 < len(s1); pos1++
		if pos2 == len(s2) || s1[pos1] != s2[pos2] {
			differences++
		} else {
			pos2++
		}
	}
	// Wenn wir am Ende genau einen Unterschied zwischen den Strings haben,
	// ist s1 um genau einen Buchstaben größer als s2.
	return differences == 1
}

// Erwartet eine Liste von Strings dataBase und einen einzelnen String substring.
// Liefert eine gefilterte mit genau den Strings aus dataBase, aus denen substring
// entsteht indem man ein Zeichen entfernt.
func FilterByExtraLetter(dataBase []string, substring string) []string {
	result := []string{}
	for _, element := range dataBase {
		if HasExtraLetter(element, substring) {
			result = append(result, element)
		}
	}
	return result
}

// Erwartet zwei Strings s1 und s2. Prüft, ob s2 aus s1 entsteht,
// indem man beliebig viele Zeichen aus s1 weglässt.
func ContainsScatteredSubstring(s1, s2 string) bool {
	// TODO
	return false
}

// Erwartet eine Liste von Strings dataBase und einen einzelnen String substring.
// Liefert eine gefilterte mit genau den Strings aus dataBase, aus denen substring
// entsteht indem man beliebig viele Zeichen entfernt.
func FilterByScatteredSubstring(dataBase []string, substring string) []string {
	result := []string{}
	for _, element := range dataBase {
		if ContainsScatteredSubstring(element, substring) {
			result = append(result, element)
		}
	}
	return result
}

// Hilfsfunktion: Liefert den String, der aus s entsteht, wenn man die Zeichen
// an den Stellen i und j vertauscht.
func SwitchLetters(s string, i, j int) string {
	result := ""
	// TODO
	return result
}

// Erwartet zwei Strings s1 und s2. Prüft, ob s2 aus s1 entsteht,
// indem man zwei benachbarte unterschiedliche Zeichen vertauscht.
func ContainsSwitchedLetters(s1, s2 string) bool {
	// TODO
	return false
}

// Erwartet eine Liste von Strings dataBase und einen einzelnen String substring.
// Liefert eine gefilterte mit genau den Strings aus dataBase, aus denen substring
// durch Vertauschen benachbarter Buchstaben hervorgeht
func FilterBySwitchedLetters(dataBase []string, substring string) []string {
	result := []string{}
	// TODO
	return result
}

// Hilfsfunktion: Prüft, ob der gegebene String in der Liste enthalten ist.
func ListContains(list []string, s string) bool {
	// TODO
	return false
}

// Erwartet eine Liste von Strings und liefert die gleiche Liste ohne Duplikate.
func RemoveDuplicates(list []string) []string {
	result := []string{}
	// TODO
	return result
}

// Erwartet eine Liste von Strings dataBase und einen einzelnen String incomplete.
// Liefert eine Liste von Vorschlägen zu Vervollständigung von incomplete.
// Die Vorschläge sollen in folgender Reihenfolge in der Liste vorkommen:
//
// 1. Alle gültigen Fortsetzungen von incomplete
// 2. Alle Strings, aus denen incomplete durch Buchstabedreher entstanden sein könnte
// 3. Alle Strings, bei deren incomplete eines oder mehrere Zeichen fehlen.
//
// Die Liste der Vorschläge enthält keine Duplikate!
func SuggestCompletions(dataBase []string, incomplete string) []string {
	// TODO
	return []string{}
}
