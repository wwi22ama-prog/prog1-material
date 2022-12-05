package misc

// Erwartet eine Zell-Adresse und entfernt alles, was kein Großbuchstabe ist.
// Liefert einen neuen String mit dem Ergebnis.
func FilterUpperCase(cellAddress string) string {
	result := ""
	for _, el := range cellAddress {
		if IsUpperCase(el) {
			result += string(el)
		}
	}
	return result
}

// Erwartet eine Zell-Adresse und entfernt alles, was keine Ziffer ist.
// Liefert einen neuen String mit dem Ergebnis.
func FilterDigits(cellAddress string) string {
	result := ""
	for _, el := range cellAddress {
		if IsDigit(el) {
			result += string(el)
		}
	}
	return result
}

// Erwartet ein Zeichen und liefert true, falls dies ein Großbuchstabe ist.
func IsUpperCase(b rune) bool {
	return b >= 'A' && b <= 'Z'
}

// Erwartet ein Zeichen und liefert true, falls dies eine Ziffer ist.
func IsDigit(b rune) bool {
	return b >= '0' && b <= '9'
}
