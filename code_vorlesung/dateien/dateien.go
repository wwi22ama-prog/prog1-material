package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {

	lines := ReadLinesFromFile("wetter.csv")

	fmt.Println("Temperaturen am Morgen:")
	for _, line := range lines {
		data := strings.Split(line, ";")
		//fmt.Println(data)
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
	file, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(file)

	result := []string{}
	for scanner.Scan() {
		result = append(result, scanner.Text())
	}
	return result
}
