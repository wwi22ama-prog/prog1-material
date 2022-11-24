package misc

// Erwartet eine Zell-Adresse und entfernt alles, was kein Großbuchstabe ist.
// Liefert einen neuen String mit dem Ergebnis.
func FilterUpperCase(cellAddress string) string {
	result := ""
	// TODO
	return result
}

// Erwartet eine Zell-Adresse und entfernt alles, was keine Ziffer ist.
// Liefert einen neuen String mit dem Ergebnis.
func FilterDigits(cellAddress string) string {
	result := ""
	for _, c := range cellAddress {
		if c >= '0' && c <= '9' {
			result += string(c)
		}
	}
	return result
}

// Erwartet ein Zeichen und liefert true, falls dies ein Großbuchstabe ist.
func IsUpperCase(b rune) bool {
	// TODO
	return false
}

// Erwartet ein Zeichen und liefert true, falls dies eine Ziffer ist.
func IsDigit(b rune) bool {
	// TODO
	return false
}
