package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {

	// Eine Datei einlesen.
	lines := ReadLinesFromFile("wetter.csv")

	// Die Daten auswerten.
	fmt.Println("Temperaturen am Morgen:")

	// Für jede gelesene Zeile...
	for _, line := range lines {

		// ... die Zeile auftrennen in eine Liste von einzelnen Strings.
		data := strings.Split(line, ";")

		// ... ausgewählte Daten aus der Liste ausgeben.
		fmt.Printf("    %v: %v Grad\n", data[0], data[1])
	}
	fmt.Println("Temperaturen am Mittag:")
	for _, line := range lines {
		data := strings.Split(line, ";")
		//fmt.Println(data)
		fmt.Printf("    %v: %v Grad\n", data[0], data[2])
	}
}

// Liest Zeilen aus einer CSV-Datei und liefert sie als ...
func ReadLinesFromFile(fileName string) []string {

	// Eine Datei öffnen.
	// os.Open() erwartet einen Dateinamen und liefert ein File- und ein Error-Objekt.
	file, err := os.Open(fileName)

	// Err ist nicht nil (d.h. es existiert), falls die Datei nicht geöffnet werden konnte.
	if err != nil {
		log.Fatal(err)
	}

	// Die Datei schließen.
	defer file.Close()

	// Scanner erzeugen. Ein Scanner kann Datenströme lesen und interpretieren.
	scanner := bufio.NewScanner(file)

	// Mit dem Scanner Zeile für Zeile lesen und als Strings liefern.
	result := []string{}
	for scanner.Scan() {
		result = append(result, scanner.Text())
	}
	return result
}
