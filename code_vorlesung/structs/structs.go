package main

import "fmt"

// Datentyp für Einträge in einem Deutsch-Englisch-Wörterbuch.
type Eintrag struct {
	De string
	En string
}

type Woerterbuch []Eintrag

func main() {

	// Einen Eintrag erstellen.
	e1 := Eintrag{"Auto", "car"}

	fmt.Println(e1)
	fmt.Println(e1.String())
	fmt.Println(e1.De)
	fmt.Println(e1.En)

	//e1.De = "LKW"
	//fmt.Println(e1)

	// Noch einen Eintrag erstellen.
	e2 := Eintrag{"Fahrrad", "bicycle"}

	// Beide Einträge zum Wörterbuch hinzufügen.
	w1 := Woerterbuch{e1, e2}

	fmt.Println(w1)

	query := "Auto"
	fmt.Println(w1.Uebersetzung(query))

}

// Abfrage des Wörterbuchs nach query.
func (w Woerterbuch) Uebersetzung(query string) string {
	for _, e := range w {
		if e.De == query {
			return e.En
		}
	}
	return ""
}

// Gibt einen Eintrag als String zurück.
func (e Eintrag) String() string {
	return fmt.Sprintf("%s : %s", e.De, e.En)
}

// Gibt ein ganzes Wörterbuch als String zurück.
func (w Woerterbuch) String() string {
	result := ""
	for _, e := range w {
		result += e.String() + "\n"
	}
	return result
}
