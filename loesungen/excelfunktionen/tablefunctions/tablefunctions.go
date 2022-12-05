package tablefunctions

import "strings"

// Erwartet einen Funktions-Ausdruck und zerlegt ihn in eine Liste seiner Bestandteile:
// Die Liste enthält folgende Teile:
// Position 0: Name der Funktion
// An weiteren Positionen: Die Argumente des Funktionsaufrufs
func ParseFunctionCall(callExpression string) []string {
	splitCall := strings.Split(callExpression, "(")
	funcName := splitCall[0]
	params := strings.Split(splitCall[1], ";")
	params[len(params)-1] = strings.TrimRight(params[len(params)-1], ")")
	return append([]string{funcName}, params...)
}

// Erwartet einen Funktions-Ausdruck und liefert den Funktionsnamen.
func GetFunctionName(callExpression string) string {
	return ParseFunctionCall(callExpression)[0]
}

// Erwartet einen Funktions-Ausdruck und liefert das n-te Argument.
func GetArgument(callExpression string, n int) string {
	return ParseFunctionCall(callExpression)[n+1]
}
