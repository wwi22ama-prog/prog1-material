# Aufgaben zur Nachbildung von Excel-Funktionalität

Eine Reihe von Aufgaben, bei denen es darum geht, Funktionen zu implementieren,
die denen von Tabellenkalkulationen wie z.B. Excel ähnlich sind.

Beispiele für solche Funktionen sind:

* `MAX()`: Bestimmen des größten Elements eines Bereichs.
* `SUMME()`: Bestimmen der Summe aller Zahlen in einem Bereich.
* `VERWEIS()`: Suche nach einem Element in einem Bereich und liefern des entsprechenden
  Elements aus einem anderen Bereich.
* `SVERWEIS()`: Suche nach einem Element in der ersten Spalte eines Bereichs und liefern
  des entsprechenden Elements aus einer anderen Spalte.

## Modellierung eines Tabellenblatts

Ein Tabellenblatt besteht aus Zeilen und Spalten, die z.B. in der Form "A4", "C23"
oder "A2:D15" adressiert werden können.
Technisch gesehen ist das Blatt eine zweidimensionale Liste.
Um die Umsetzung von Excel-Funktionen zu erleichtern, werden eine Reihe von
teilweise aufeinander aufbauenden Hilfsfunktionen gebraucht:

## Überblick zur Vorgehensweise

Um das Projekt umzusetzen, wird eine Reihe von aufeinander aufbauenden Hilfsfunktionen
gebraucht. Diese werden hier kurz beschrieben, um einen Überblick zu geben.

Die Funktionen sind in Packages gruppiert. Diese werden hier kurz beschrieben und
befinden sich in den Unterverzeichnissen dieses Moduls.
In den jeweiligen Readme-Dateien werden auch die Voraussetzungen der Packages
beschrieben, anhand derer Sie entscheiden können, in welcher Reihenfolge sie die
Aufgaben bearbeiten.

### Auswertung von Adressen (Package `addressfunctions`)

Bei diesen Funktionen geht es darum, Adress-Ausdrücke wie "A4" oder "B2:D15" zu
analysieren (auch "parsen" genannt).

* Zeile und Spalte eines Zell-Ausdrucks (z.B. "A4" bestimmen)
* Anzahl der Zeilen und Spalten eines Bereichs-Ausdrucks (z.B. "B2:D15") bestimmen.
* Erste Zeile und erste Spalte eines Bereichs-Ausdrucks bestimmen.

### Liefern von Daten zu Adressen (Package `tableaccess`)

Hier geht es darum, zu gegebenen Adress-Ausdrücken die passenden Daten aus einer
Gesamt-Tabelle zu liefern.

* Den tatsächlichen Wert bei einem Zell-Ausdruck.
* Eine (ein- oder zweidimensionale) Liste bei einem Bereichs-Ausdruck.

### Auswertung von Funktions-Ausdrücken (Package `tablefunctions)`

Der nächste Schritt ist es, nicht nur Adress-Ausdrücke zu verstehen, sondern auch
Funktionsaufrufe parsen zu können. Es sollen Ausdrücke wie `SUMME(A4:A15)`
verstanden und in Teilinformationen zerlegt werden.

* Erkennen der verwendeten Funktion
* Erkennen der einzelnen Argumente
* Erkennen der betroffenen Zellen bzw. Bereiche

### Technische Umsetzung der Datenbankfunktionen (Package `listfunctions`)

Implementierung von Funktionen, die für die umzusetzenden Excel-Funktionalitäten
hilfreich sind.

* Bestimmen des Maximums aus einer Liste.
* Bestimmen der Summe aller Elemente einer Liste.
* Suchfunktion für Elemente in Listen.

### Allgemeine Hilfsfunktionen (Package `misc`)

Verschiedene Hilfsfunktionen, z.B. um Strings zu analysieren.

### Umsetzung der Tabellen-Funktionen (Package `spreadsheetfunctions`)

Mit den bisherigen Hilfsfunktionen können die Tabellenfunktionen nun umgesetzt werden.
